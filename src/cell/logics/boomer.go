// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package boomer provides commands to run load tests and display results.
package logics

import (
	"crypto/tls"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"bytes"
	"html/template"
	"io"

	"cell/function"
	"cell/protocol"
	"master/settings"

	"log"
)

type result struct {
	err           error
	statusCode    int
	duration      time.Duration
	contentLength int64
}

type ConTime struct {
	Conc    int
	ReqTime int // second
}

type Boomer struct {
	// Request is the request to be made.
	//Url  string
	Body string
	Host string

	Request  *http.Request
	Requests []*http.Request

	RequestBody string

	// N is the total number of requests to make.
	N int

	// C is the concurrency level, the number of concurrent workers to run.
	C int

	// Timeout in seconds.
	Timeout int

	// Qps is the rate limit.
	Qps int

	// DisableCompression is an option to disable compression in response
	DisableCompression bool

	// DisableKeepAlives is an option to prevents re-use of TCP connections between different HTTP requests
	DisableKeepAlives bool

	// Output represents the output type. If "csv" is provided, the
	// output will be dumped as a csv stream.
	Output string

	// ProxyAddr is the address of HTTP proxy server in the format on "host:port".
	// Optional.
	ProxyAddr *url.URL

	results chan *result

	Cancel chan int
	Done   chan int

	IsCancel int32
}

func (b *Boomer) IsStop() bool {
	select {
	case _, ok := <-b.Cancel:
		return !ok
	default:
		return false
	}
	return false
}

func (b *Boomer) Stop() {
	atomic.AddInt32(&b.IsCancel, 10)

	if !b.IsStop() {
		close(b.Cancel)
	}
	<-b.Done

	CurrBoom = nil
	TaskId = 0
}

func (b *Boomer) StopAsync() {
	atomic.AddInt32(&b.IsCancel, 10)
	if !b.IsStop() {
		close(b.Cancel)
	}

	CurrBoom = nil
	TaskId = 0
}

// Run makes all the requests, prints the summary. It blocks until
// all work is done.
func (b *Boomer) Run(taskid int64) {
	rand.Seed(time.Now().Unix())

	b.results = make(chan *result, b.N)

	b.Done = make(chan int, 1)
	b.Cancel = make(chan int)
	b.IsCancel = 0

	go b.ProgressTick()

	start := time.Now()
	b.runWorkers()

	b.Done <- 0

	if !b.IsStop() {
		close(b.Cancel)
	}

	b.SendEndProgress()

	log.Println(CellId, " work done========================")

	if b.IsCancel == 0 {
		newReport(b.N, b.C, b.results, b.Output, time.Now().Sub(start), taskid).finalize()
	}

	close(b.results)

	if TaskId == taskid {
		CurrBoom = nil
		TaskId = 0
	}
}

func (b *Boomer) makeRequest(c *http.Client) {
	s := time.Now()
	var size int64
	var code int

	req := b.Request

	if len(b.results) > 1 {
		idx := rand.Intn(len(b.Requests))
		req = b.Requests[idx]
	}

	resp, err := c.Do(cloneRequest(req, b.RequestBody))
	duration := time.Now().Sub(s)

	if err == nil {
		code = resp.StatusCode
		size = resp.ContentLength
		io.Copy(ioutil.Discard, resp.Body)
	}

	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}

	b.results <- &result{
		statusCode:    code,
		duration:      duration,
		err:           err,
		contentLength: size,
	}
}

func (b *Boomer) runWorker(n int, tr *http.Transport) {
	var throttle <-chan time.Time
	if b.Qps > 0 {
		throttle = time.Tick(time.Duration(1e6/(b.Qps)) * time.Microsecond)
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(b.Timeout) * time.Second,
	}

	for i := 0; i < n; i++ {
		if b.IsStop() {
			return
		}

		if b.Qps > 0 {
			<-throttle
		}

		b.makeRequest(client)
	}
}

func (b *Boomer) runWorkers() {
	var wg sync.WaitGroup

	MaxIdleConns := b.C
	MaxIdleConnsPerHost := b.C
	IdleConnTimeout := 5 * time.Second
	ExpectContinueTimeout := 5 * time.Second

	if b.DisableKeepAlives {
		MaxIdleConns = 0
		MaxIdleConnsPerHost = 0
		IdleConnTimeout = 0 * time.Second
		ExpectContinueTimeout = 0 * time.Second
	}

	// have rwlock in transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},

		DisableCompression: b.DisableCompression,
		DisableKeepAlives:  b.DisableKeepAlives,

		//TLSHandshakeTimeout: time.Duration(b.Timeout) * time.Second,
		//Proxy:               http.ProxyURL(b.ProxyAddr),

		//DialContext: (&net.Dialer{
		//	Timeout:   5 * time.Second,
		//	KeepAlive: 5 * time.Second,
		//}).DialContext,

		MaxIdleConns:          MaxIdleConns,
		MaxIdleConnsPerHost:   MaxIdleConnsPerHost,
		IdleConnTimeout:       IdleConnTimeout,
		ExpectContinueTimeout: ExpectContinueTimeout,
	}

	// Ignore the case where b.N % b.C != 0.

	num := b.N/b.C
	for i := 0; i < b.C; i++ {
		if b.IsStop() {
			break
		}

		wg.Add(1)
		go func() {
			b.runWorker(num, tr)
			wg.Done()
		}()
	}
	wg.Wait()
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request, body string) *http.Request {
	r2 := new(http.Request)
	*r2 = *r

	url0_1 := r.URL.String()
	url0, _ := url.QueryUnescape(url0_1)

	if r.Method == "GET" {
		var buf bytes.Buffer

		tmpl, err := template.New("").Parse(url0)
		if err != nil {
		}

		f := NewFuncsType()
		tmpl.Execute(&buf, f)

		url1 := buf.String()

		if url1 != url0 {
			r2, _ = http.NewRequest(r.Method, url1, nil)
		}

	} else if r.Method == "POST" || r.Method == "PUT" {

		tmpl, err := template.New("").Parse(body)
		if err != nil {
		}
		var buf bytes.Buffer
		f := NewFuncsType()
		tmpl.Execute(&buf, f)
		body = buf.String()
		r2.ContentLength = int64(len([]byte(body)))
		r2.Body = ioutil.NopCloser(strings.NewReader(body))
	}

	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	r2.Host = r.Host
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}

	return r2
}

func (b *Boomer) Progress() int32 {
	percent := float32(len(b.results)) / float32(b.N) * 100
	return int32(percent)
}

func (b *Boomer) ProgressTick() {
	for {
		if b.IsStop() {
			return
		}

		b.SendProgress()

		<-time.After(time.Millisecond * 500)
	}
}

func (b *Boomer) SendProgress() {
	p := b.Progress()
	sendpb := &protocol.Task_Progress_C2S{
		TaskId:   &TaskId,
		Progress: &p,
	}

	function.NetworkManage.Send(settings.Task_Progress_C2S, sendpb)
}

func (b *Boomer) SendEndProgress() {
	var p int32 = 100
	sendpb := &protocol.Task_Progress_C2S{
		TaskId:   &TaskId,
		Progress: &p,
	}
	function.NetworkManage.Send(settings.Task_Progress_C2S, sendpb)
}

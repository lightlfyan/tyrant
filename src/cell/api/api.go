package api

import (
	"cell/function"
	"cell/logics"
	"cell/protocol"
	"encoding/base64"
	"master/settings"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"

	"io/ioutil"
	"log"
)

func Login(id int32) {
	pb := &protocol.Cell_Login_C2S{Id: &id}
	function.NetworkManage.Send(settings.Cell_Login_C2S, pb)
}

func Cell_Login_S2C(pb *protocol.Cell_Login_S2C) {
	log.Println("recive loginS2C ", *pb.Result)

	if *pb.Result {
		logics.CellId = *pb.Id
	} else {
		// 暴力退出
		panic("login error")
	}
}

func Task_Create_C2S(pb *protocol.Task_Create_S2C) {
	username := *pb.UserName
	password := *pb.Password
	body := *pb.Body
	accept := *pb.Accept
	contentType := *pb.ContentType
	num := int(*pb.Num)
	conc := int(*pb.Conc)
	qps := int(*pb.Qps)
	timeout := int(*pb.Timeout)
	method := *pb.Method
	urlstr := *pb.Url

	urls := Filter(strings.Split(urlstr, ";"), func(x string) bool { return x != "" })
	firsturl := urls[0]

	header := make(http.Header)
	if accept != "" {
		header.Set("Accept", accept)
	}

	if contentType != "" {
		header.Set("Content-Type", contentType)
	}

	for _, kv := range pb.Headers {
		if *kv.K == "" {
			continue
		}
		header.Set(*kv.K, *kv.V)
	}

	req, err := http.NewRequest(method, firsturl, nil)

	if err != nil {
		log.Println(err.Error())
	}

	req.Header = header

	host := *pb.Host
	if strings.HasPrefix(host, "http://") {
		host = host[7:]
	} else if strings.HasPrefix(host, "https://") {
		host = host[8:]
	}

	if *pb.Host != "" {
		req.Host = host
	}

	if username != "" || password != "" {
		req.SetBasicAuth(username, password)
	}

	/*
		if req.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			v, _ := url.ParseQuery(body)
			body = v.Encode()
		} else if req.Header.Get("Content-Type") == "application/json" {
			var json_bytes = []byte(body)
			buffer := new(bytes.Buffer)
			if err := json.Compact(buffer, json_bytes); err == nil {
				body = buffer.String()
			}
		}
	*/

	if method == "POST" || method == "PUT" {
		req.Body = ioutil.NopCloser(strings.NewReader(body))
		req.ContentLength = int64(len([]byte(body)))
	}

	req.Header.Set("User-Agent", "tyrant")

	reqs := make([]*http.Request, 0, len(urls))
	reqs = append(reqs, req)

	for i := 1; i < len(urls); i++ {
		reqs = append(reqs, cloneRequest(req, urls[i]))
	}

	boomer := &logics.Boomer{
		Body: body,
		//Host:        Host,
		//FastRequest: fastreq,

		Request:  req,
		Requests: reqs,

		RequestBody:        body,
		N:                  num,
		C:                  conc,
		Qps:                qps,
		Timeout:            timeout,
		DisableCompression: true,
		DisableKeepAlives:  !*pb.KeepAlive,
		//ProxyAddr:          url1,
		Output: "",
	}

	logics.NewTask(boomer, *pb.TaskId)

	sendpb := &protocol.Task_Create_C2S{
		TaskId: pb.TaskId,
	}
	function.NetworkManage.Send(settings.Task_Create_C2S, sendpb)

	//fakeFileName := "fake_" + *pb.Uid + ".txt"
	//Loadfakedata(fakeFileName)

	go boomer.Run(*pb.TaskId)
}

func cloneRequest(r *http.Request, url string) *http.Request {
	log.Println(url)
	// shallow copy of the struct
	r2, _ := http.NewRequest(r.Method, url, nil)
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}

	return r2
}

func hasPort(s string) bool { return strings.LastIndex(s, ":") > strings.LastIndex(s, "]") }

func removeEmptyPort(host string) string {
	if hasPort(host) {
		return strings.TrimSuffix(host, ":")
	}
	return host
}

func SetBasicAuth(r *fasthttp.Request, username, password string) {
	r.Header.Set("Authorization", "Basic "+basicAuth(username, password))
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"time"

	"flag"
	"fmt"
	"io/ioutil"
	"sync"

	"runtime"
	"strconv"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var concArg = flag.Int("conc", 10, "concurrency num")
var qpsArg = flag.Int("qps", 500, "qps limit")
var num = flag.Int("num", 10000, "num package send")
var host = flag.String("host", "10.104.118.36:9090", "ws host")

var pre = flag.String("pre", "pre", "uid pre")

var qps time.Duration
var conc int

var wgConn sync.WaitGroup
var wgReq sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	flag.Parse()

	conc = *concArg
	//qps = time.Second / (time.Duration(*qpsArg) / time.Duration(conc))
	//log.Println("qps: ", qps)

	wgConn.Add(conc)
	for i := 0; i < conc; i++ {
		makeRequest(i)
	}

	wgConn.Wait()

	startime := time.Now()
	wgReq.Wait()
	duration := time.Now().Sub(startime)

	qps1 := float64(*num) / duration.Seconds()

	log.Println("duration", duration, qps1)
}

func makeRequest(uid int) {
	defer wgConn.Done()

	uname := *pre + strconv.Itoa(uid)

	dest := fmt.Sprintf(`ws://%s/hermes/connect?zoneid=show_lock_seat&uid=%s%d&token=123&timestamp=%d&channels=seat`, *host, *pre, uid, time.Now().Unix())
	conn, resp, err := websocket.DefaultDialer.Dial(dest, nil)
	if err != nil {
		if resp != nil && resp.Body != nil {
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println("body: ", string(body))
		}
		log.Println(uid, err)
		return
	}

	log.Println("connedted: ", *pre, uid)
	c := &Client{
		uname: uname,
		conn:  conn,
		close: make(chan []int, 1),
	}

	wgReq.Add(1)
	go c.DoRequest()
}

// Client is an middleman between the websocket connection and the hub.
type Client struct {
	uname string
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	close chan []int
}

func (c *Client) DoRequest() {
	defer close(c.close)
	defer wgReq.Done()

	go c.writePump()
	c.readPump()
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
	}()

	counter := 0

	//c.conn.SetReadLimit(maxMessageSize)
	//c.conn.SetReadDeadline(time.Now().Add(pongWait))

	//c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		select {
		case _, ok := <-c.close:
			if ok {
				return
			}
		default:
			break
		}

		//_, message, err := c.conn.ReadMessage()
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		counter += 1

		if counter >= *num/conc {
			break
		}

		//io.Copy(ioutil.Discard, message)
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//log.Println(string(message))
	}
}

func (c *Client) writePump() {
	//ticker := time.NewTicker(qps)

	defer func() {
		//ticker.Stop()
		c.conn.Close()
	}()

	/*
		m := map[string]interface{}{
			"msg_id":     2004,
			"channel_id": "seat",
			"msg":        "test",
		}
		message, _ := json.Marshal(m)
	*/

	message := []byte("ping " + c.uname)

	for {
		//<-ticker.C
		c.conn.SetWriteDeadline(time.Now().Add(writeWait))
		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(message)
		if err := w.Close(); err != nil {
			log.Println(err)
			return
		}

		//c.conn.SetWriteDeadline(time.Now().Add(writeWait))
		//if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
		//	return
		//}
	}
}

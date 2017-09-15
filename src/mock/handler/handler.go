package handler

import (
	"mock/cfg"
	"time"

	"bytes"
	"compress/gzip"
	"math/rand"
	"strings"

	"fly"
	"strconv"
)

func MockHandler(ctx *fly.Context) {
	urlKey := ctx.Param("url")

	/*
	   query1 := ctx.Request.URL.Query().Get("query1")
	   fmt.Println(query1)

	   postkey := ctx.Request.PostFormValue("postkey")
	   fmt.Println(postkey)
	*/

	origin := ctx.Request.Header.Get("Origin")
	if origin != "" {
		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Allow-Credentials", "true")
	}

	urlCfg := cfg.GetCfg(urlKey, ctx.Request.Method)

	if urlCfg == nil {
		ctx.WriteString(500, "error")
		return
	}

	if urlCfg.ApiType == 1 {
		text := RandStringRunes(1024)
		for i := 0; i < 1; i++ {
			var b bytes.Buffer
			w := gzip.NewWriter(&b)
			w.Write([]byte(text))
			w.Close()
		}
		ctx.Header("Content-Type", urlCfg.ContentType)
		ctx.WriteString(200, "ok\n")
		return
	} else if urlCfg.ApiType == 2 {
		bigData := strings.Repeat("a", 1024*100)
		ctx.Header("Content-Type", urlCfg.ContentType)
		ctx.WriteString(200, bigData)
		return
	}

	channel, ok := cfg.UrlChanMap[urlKey]

	var timeWait <-chan time.Time

	if urlCfg.TimeWaitMs > 0 {
		timeWait = time.After(time.Duration(urlCfg.TimeWaitMs) * time.Millisecond)
	}

	respContent := urlCfg.Resp

	if respContent == "timestamp" {
		respContent = strconv.FormatInt(time.Now().Unix(), 10)
	}

	code := urlCfg.Code

	if ok {
		// 阻塞直到超时
		if urlCfg.TimeOutMs > 0 {
			select {
			case channel <- 0:
				defer func() {
					<-channel
				}()
			case <-time.After(time.Duration(urlCfg.TimeOutMs) * time.Millisecond):
				code = 408
				respContent = "timeout"
			}
		} else {
			//阻塞直到完成
			channel <- 0
			defer func() {
				<-channel
			}()
		}
	}

	// wait time
	if timeWait != nil {
		<-timeWait
	}

	lens := len([]byte(respContent))
	ctx.Header("Content-Length", strconv.FormatInt(int64(lens), 10))
	ctx.Header("Content-Type", urlCfg.ContentType)

	ctx.WriteString(code, respContent)
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

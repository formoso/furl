package furl

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	Url         string
	ElapsedTime int64
	NBytes      int64
	Body        []byte
	Err         error
}

func Get(urls []string) *chan Response {
	ch := initChan()
	for i := 1; i < len(urls); i++ {
		url := urls[i]
		r := initResponse(url)
		if r.Url != "body" {
			go readUrl(i, r, ch)
		}
	}
	return ch
}

func initResponse(url string) Response {
	r := Response{
		Url:         url,
		ElapsedTime: 0,
		NBytes:      0,
		Body:        []byte{},
		Err:         nil,
	}
	return r
}

func initChan() *chan Response {
	ch := make(chan (Response))
	return &ch
}

func readUrl(i int, r Response, ch *chan Response) {
	start := time.Now()
	resp, _ := http.Get(r.Url)
	_, r.Err = http.Get(r.Url)
	if r.Err != nil {
		*ch <- r
		return
	}
	defer resp.Body.Close()
	r = readBody(r, resp)
	if r.Err != nil {
		*ch <- r
		return
	}
	r = readNBytes(r, resp)
	if r.Err != nil {
		*ch <- r
		return
	}
	r.ElapsedTime = time.Since(start).Milliseconds()
	*ch <- r
}

func readBody(r Response, resp *http.Response) Response {
	r.Body, r.Err = io.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(r.Body))
	return r
}

func readNBytes(r Response, resp *http.Response) Response {
	r.NBytes, r.Err = io.Copy(ioutil.Discard, resp.Body)
	return r
}

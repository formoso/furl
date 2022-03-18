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

func Get(url []string, canal chan Response) {
	for i := 1; i < len(url); i++ {
		go func(i int) {
			start := time.Now()
			urls := url[i]
			r := initResponse(urls)
			resp, err := http.Get(urls)
			if err != nil {
				canal <- r
				return
			}

			defer resp.Body.Close()
			r.Body, err = io.ReadAll(resp.Body)
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(r.Body))
			r.NBytes, err = io.Copy(ioutil.Discard, resp.Body)
			if err != nil {
				canal <- r
				return
			}
			r.ElapsedTime = time.Since(start).Milliseconds()
			canal <- r
		}(i)
	}
}

func initResponse(url string) Response {
	r := Response{
		Url:         url,
		ElapsedTime: 0,
		NBytes:      0,
		Body:        []byte{},
	}
	return r
}

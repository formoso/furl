package furl

import (
	"io"
	"net/http"
	"time"
)

type Response struct {
	Url         string
	ElapsedTime int64
	NBytes      int
	Body        []byte
}

func Get(url string) (Response, error) {
	start := time.Now()
	r := initResponse(url)
	resp, err := http.Get(url)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()
	r.Body, err = io.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}
	r.NBytes = len(string(r.Body))
	r.ElapsedTime = time.Since(start).Milliseconds()

	return r, err
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

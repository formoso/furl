package furl

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/formoso/furl/pkg/arguments"
)

type Response struct {
	Url         string
	ElapsedTime int64
	NBytes      int64
	Body        []byte
	Err         error
}

func Get(urls []string) {
	ch := initChan()
	for i := 1; i < len(urls); i++ {
		url := urls[i]
		r := initResponse(url)
		if r.Url != "body" {
			go readUrl(i, r, ch)
		}
	}
	respostaGet(ch)
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

func respostaGet(ch *chan Response) {
	body, leng := arguments.NoBody()
	for i := 1; i < leng; i++ {
		resp := <-*ch
		if resp.Err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching URL:%v\n", resp.Err)
			os.Exit(1)
		}
		if body {
			printerBody(resp)
		} else {
			printerNoBody(resp)
		}
	}
}

func printerNoBody(resp Response) {
	fmt.Printf("%dms %7d %s\n", resp.ElapsedTime, resp.NBytes, resp.Url)
}

func printerBody(resp Response) {
	fmt.Printf("%dms %7d %s \n Body:%s\n", resp.ElapsedTime, resp.NBytes, resp.Url, resp.Body)
}

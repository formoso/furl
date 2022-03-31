package furl

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Url         string
	ElapsedTime int64
	NBytes      int64
	Body        []byte
	Err         error
}

func Get(urls []string, wBody bool) {
	ch := initChan()
	for _, url := range urls {
		r := initResponse(url)
		if r.Url != "body" {
			go readUrl(r, ch, wBody)
		}
	}
	respostaGet(ch, len(urls))
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

func readUrl(r Response, ch *chan Response, wBody bool) {
	start := time.Now()
	resp, err := http.Get(r.Url)
	if r.Err != nil {
		r.Err = err
		*ch <- r
		return
	}
	defer resp.Body.Close()
	if wBody {
		r = readBody(r, resp)
		if r.Err != nil {
			*ch <- r
			return
		}
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

func respostaGet(ch *chan Response, lenArray int) {

	for i := 0; i < lenArray; i++ {
		resp := <-*ch
		if resp.Err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching URL:%v\n", resp.Err)
			os.Exit(1)
		}
		printer(resp)
	}
}
func printer(resp Response) {
	var sBody string
	if len(resp.Body) > 0 {
		sBody = fmt.Sprintf("Body:%s\n", resp.Body)
	}
	fmt.Printf("%dms %7d %s \n%s", resp.ElapsedTime, resp.NBytes, resp.Url, sBody)
}

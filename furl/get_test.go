package furl

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet_OK(t *testing.T) {
	want := "Success!"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(want))
	}))
	defer srv.Close()
	canal := make(chan Response)
	url := []string{"", srv.URL}
	go Get(url, canal)
	resp := <-canal
	if resp.Err != nil {
		t.Errorf("Unexpected error on request: %s", resp.Err)
	}
}

func TestGet_Error_Get(t *testing.T) {
	canal := make(chan Response)
	url := []string{"", "321"}
	go Get(url, canal)
	resp := <-canal
	if resp.Err != nil {
		t.Errorf("Unexpected error on request: %s", resp.Err)
	}
	if resp.Err != nil && resp.ElapsedTime > 0 && resp.NBytes == 0 && resp.Url == "123" {
		t.Errorf("no value in request is expected: %v", resp)
	}
}

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
	resp, err := Get(srv.URL)
	if err != nil {
		t.Errorf("Unexpected error on request: %s", err)
	}
	if string(resp.Body) != want {
		t.Errorf("want %s, got %s", want, resp.Body)
	}
}

func TestGet_Error_Get(t *testing.T) {
	resp, err := Get("123")
	if err == nil {
		t.Errorf("Error expected on request: %v", resp)
	}
	if resp.Body != nil && resp.ElapsedTime > 0 && resp.NBytes == 0 && resp.Url == "123" {
		t.Errorf("no value in request is expected: %v", resp)
	}
}

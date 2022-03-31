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
	url := []string{srv.URL}
	wBody := true
	Get(url, wBody)
}

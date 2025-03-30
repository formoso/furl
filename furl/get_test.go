package furl

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Get_Success(t *testing.T) {
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

func Test_Get_NewRequest_Error(t *testing.T) {
	expected := "parse \"http://[fe80::1%en0]/\": invalid URL escape \"%en\""
	resp, err := Get("http://[fe80::1%en0]/")
	if err == nil {
		t.Errorf("Error expected on request: %v", resp)
	}
	if err.Error() != expected {
		t.Errorf("Expected error: %v, but received: %v", expected, err.Error())
	}

}

func Test_Get_DO_Error(t *testing.T) {
	resp, err := Get("123")
	if err == nil {
		t.Errorf("Error expected on request: %v", resp)
	}
	if resp.Body != nil && resp.ElapsedTime > 0 && resp.NBytes == 0 && resp.Url == "123" {
		t.Errorf("no value in request is expected: %v", resp)
	}
}

func TestBody_Nil_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// At this point the error is forced when reading the body by the client.
		// The client expects a body with size 1, but no body is being sent.
		// This causes io.ReadAll to return an error.
		w.Header().Set("Content-Length", "1")
	}))
	defer srv.Close()

	_, err := Get(srv.URL)
	if err == nil {
		t.Errorf("err should equal unexpected EOF; want %v", err)
	}
}

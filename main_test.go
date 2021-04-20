package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRoute(t *testing.T) {
	r := InitRouter()

	req, _ := http.NewRequest("GET", "/api/v1", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	want := "{\"message\": \"get called\"}"
	got := res.Body.String()

	if got != want {
		t.Errorf("want %v; got %v", want, got)
	}
}
func TestGetRouteAlternative(t *testing.T) {

	r := InitRouter()

	s := httptest.NewServer(r)
	defer s.Close()

	resp, err := http.Get(s.URL + "/api/v1")
	if err != nil {
		t.Errorf("test failed with error: %v", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	got := string(b)

	want := "{\"message\": \"get called\"}"

	if got != want {
		t.Errorf("unexpected response body %v; wanted %v", got, want)
	}
}

package github_test

import (
	"net/http"
	"net/http/httptest"
)

func mockServer(responseHeader map[string]string, responseBody []byte) *httptest.Server {
	handler := func(rw http.ResponseWriter, r *http.Request) {
		for key, value := range responseHeader {
			rw.Header().Set(key, value)
		}
		rw.Write(responseBody)
	}
	return httptest.NewServer(http.HandlerFunc(handler))
}

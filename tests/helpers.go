package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

func JSONRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	return req, w
}

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response 

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func Test_Auth(t *testing.T) {
	jsonToReturn := `
	{
		"error": false,
		"message": "some message"
	}`

	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header: make(http.Header),
		}
	})

	testApp.Client = client

	postBody := map[string]interface{}{
		"email": "me@here.net",
		"password": "verysecret",
	}

	body, _ := json.Marshal(postBody)

	req, _ := http.NewRequest("POST", "/auth", bytes.NewReader(body))
	resRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Auth)

	handler.ServeHTTP(resRecorder, req)

	if resRecorder.Code != http.StatusAccepted {
		t.Errorf("expected http.StatusAccepted but got %d", resRecorder.Code)
	}
}
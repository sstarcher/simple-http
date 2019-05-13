package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tests = []struct {
	method       string
	url          string
	expectedBody string
	headers      map[string]string
}{
	{http.MethodGet, "http://localhost/", "<p>Hello, World</p>", nil},
	{http.MethodGet, "http://localhost/hmm/yes?blah=blah&something=somethingelse", "", nil},
	{http.MethodGet, "http://localhost", "<p>Hello, World</p>",
		map[string]string{
			"Accept": "application/text",
			"Random": "application/text",
		},
	},
	{http.MethodGet, "http://localhost/", `{"message": "Good morning"}`,
		map[string]string{
			"Accept": "application/json",
		},
	},
	{http.MethodGet, "http://localhost/", `{"message": "Good morning"}`,
		map[string]string{
			"Accept": "text/json,application/json,text/html",
			"Other":  "hmm",
		},
	},
	{http.MethodPost, "http://localhost/", ``, nil},
}

func TestHelloCases(t *testing.T) {
	for index, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.url, nil)
		if tt.headers != nil {
			for key, value := range tt.headers {
				req.Header.Add(key, value)
			}
		}

		w := httptest.NewRecorder()
		helloHandler(w, req)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		actual := string(body)
		if actual != tt.expectedBody {
			t.Errorf("Hello(%d): expected vs actual \n[%s]\n[%s]", index, tt.expectedBody, actual)
		}
	}

}

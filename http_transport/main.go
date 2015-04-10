package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

var mock = flag.String("mock", "", "Doesn't perform real access when 'proxy' or 'mock' is specified")

func main() {
	flag.Parse()
	client := http.DefaultClient

	if *mock == "proxy" {
		server := newServer()
		defer server.Close()

		client.Transport = &http.Transport{
			// Proxy to httptest.Server which created above
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse(server.URL)
			},
		}
	} else if *mock == "mock" {
		client.Transport = newMockTransport()
	}

	resp, err := client.Get("http://ifconfig.co/all.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("GET http://ifconfig.co/all.json")
	fmt.Println(string(body))
}

// Create a HTTP server to return mocked response
func newServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w,
			`{
    "Accept-Encoding": [
        "proxy"
    ],
    "User-Agent": [
        "proxy"
    ],
    "X-Ip-Country": [
        "Japan(Proxy)"
    ],
    "X-Real-Ip": [
        "192.168.1.1"
    ]
}`)
	}))
}

type mockTransport struct{}

func newMockTransport() http.RoundTripper {
	return &mockTransport{}
}

// Implement http.RoundTripper
func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Create mocked http.Response
	response := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
	}
	response.Header.Set("Content-Type", "application/json")

	responseBody :=
		`{
    "Accept-Encoding": [
        "mock"
    ],
    "User-Agent": [
        "mock"
    ],
    "X-Ip-Country": [
        "Japan(Mock)"
    ],
    "X-Real-Ip": [
        "192.168.1.1"
    ]
}`
	response.Body = ioutil.NopCloser(strings.NewReader(responseBody))
	return response, nil
}

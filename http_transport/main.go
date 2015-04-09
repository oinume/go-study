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

var mock = flag.String("mock", "", "help message for long")

func main() {
	flag.Parse()

	client := http.DefaultClient

	if *mock == "proxy" {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
        defer server.Close()

		client.Transport = &http.Transport{
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

type mockTransport struct{}

func newMockTransport() http.RoundTripper {
    return &mockTransport{}
}

// http.Canceler interfaceの実装
func (t *mockTransport) CancelRequest(*http.Request) {
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    resp := &http.Response{Request: req}
    resp.StatusCode = http.StatusOK
    resp.Header = make(http.Header)
    resp.Header.Set("Content-Type", "application/json")
    jsonString :=
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
    resp.Body = ioutil.NopCloser(strings.NewReader(jsonString))
    return resp, nil
}


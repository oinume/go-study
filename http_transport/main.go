package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "net/url"
)

var mock = flag.String("mock", "", "help message for long")

func main() {
    flag.Parse()

    client := http.DefaultClient

    if *mock == "proxy" {
        server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.WriteHeader(200)
            w.Header().Set("Content-Type", "application/json")
            fmt.Fprintln(w, "all.json")
        }))

        client.Transport = &http.Transport{
            Proxy: func(req *http.Request) (*url.URL, error) {
                return url.Parse(server.URL)
            },
        }
    }

    resp, err := client.Get("http://ifconfig.co/all.json")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    fmt.Printf("%s\n", string(body))
}

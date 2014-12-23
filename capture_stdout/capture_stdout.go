package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
)

func print() {
    fmt.Println("hello world")
}

func main() {
    old := os.Stdout // keep backup of the real stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    print()

    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

    // back to normal state
    w.Close()
    os.Stdout = old // restoring the real stdout
    out := <-outC

    // reading our temp stdout
    fmt.Printf("previous output = '%s'\n", out)
}

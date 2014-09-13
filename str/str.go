package main

import (
	"fmt"
)

type JSONString string

func (s JSONString) Test() JSONString {
	return s + "_test"
}

func main() {
	var s1 string = "hello"
	var s2 string = string(s1)
	var s3 JSONString = (JSONString)(s2)
	fmt.Println(s3.Test())
}

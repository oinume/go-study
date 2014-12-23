// error interfaceを実装している場合は fmt.Printf でError()とString()どっちが呼ばれるかの確認
package main

import (
	"fmt"
)

type myerror struct {
	message string
}

func (e myerror) Error() string {
	fmt.Println("Error()")
	return e.message
}

func (e myerror) String() string {
	fmt.Println("String()")
	return e.message
}

func main() {
	err := fmt.Errorf("Hello")
	fmt.Println(err)

	myerr := myerror{"broken"}
	fmt.Printf("%s\n", myerr)
}

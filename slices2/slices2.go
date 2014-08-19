package main

import (
	"fmt"
)

func main() {
	ref()
	iterate()
}

func ref() {
	fmt.Println("ref()")
	var users []string = []string{"test0", "test1", "test2", "test3"}
	fmt.Println(users[1:3]) // test1, test2
	fmt.Println(users[1:])  // test1, test2, test3
	fmt.Println(users[:2])  // test0, test1
}

func iterate() {
	fmt.Println("iterate()")
	users := []string{"test0", "test1", "test2", "test3"}
	for i, user := range users {
		fmt.Printf("[%d] = %s\n", i, user)
	}
}

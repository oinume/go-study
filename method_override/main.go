// メソッドの継承(override)
package main

import (
	"fmt"
	"errors"
)

func main() {
	user1 := User{1, "oinume", "test"}
	result, err := user1.Login()
	fmt.Printf("%10s: Login() = %v, %v\n", user1.Name, result, err)

	user2 := BannedUser{User{2, "restricted", "test"}}
	result2, err2 := user2.Login()
	fmt.Printf("%10s: Login() = %v, %v\n", user2.Name, result2, err2)
}

type User struct {
	Id int
	Name string
	Password string
}

func (*User) Login() (bool, error) {
	return true, nil
}

type BannedUser struct {
	User
}

// User.Login()をオーバライド
func (*BannedUser) Login() (bool, error) {
	return false, errors.New("You are banned.")
}

//func (*BannedUser) Login() (bool) {
//	return false
//}

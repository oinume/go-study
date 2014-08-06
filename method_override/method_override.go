// メソッドの継承(override)
package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("--- Login() ---")
	user1 := User{1, "oinume", "test"}
	result, err := user1.Login()
	fmt.Printf("%10s: Login() = %v, %v\n", user1.Name, result, err)

	user2 := BannedUser{User{2, "banned", "test"}, "Too many spams"}
	result2, err2 := user2.Login()
	fmt.Printf("%10s: Login() = %v, %v\n", user2.Name, result2, err2)

	fmt.Println("--- String() ---")
	fmt.Println(user1.String())
	fmt.Println(user2.String())
}


type User struct {
	Id int
	Name string
	Password string
}

func (user *User) Login() (bool, error) {
	return true, nil
}

func (user *User) String() string {
	return fmt.Sprintf("Id:%d, Name:%s", user.Id, user.Name)
}

type BannedUser struct {
	User
	Reason string
}

// User.Login()をオーバーライド
func (user *BannedUser) Login() (bool, error) {
	return false, errors.New("You are banned.")
}

// User.String()をオーバーライドしつつ中で呼び出す
func (user *BannedUser) String() string {
	return user.User.String() + ", Reason:" + user.Reason
}

//func (*BannedUser) Login() (bool) {
//	return false
//}

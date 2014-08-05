package main

import (
	"fmt"
)

type User struct {
	Id int
	Name string
}

type BannedUser struct {
	User
	BannedReason string
}

type PBannedUser struct {
	*User
	BannedReason string
}

func main() {
	user := BannedUser{User{1, "oinume"}, "Too many spam"}
	fmt.Println(user.Name)

	puser := PBannedUser{&User{1, "oinume"}, "Too many spam"}
	fmt.Println(puser.User.Name)
}

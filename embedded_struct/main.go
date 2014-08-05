package main

import (
	"fmt"
)

type User struct {
	Id int
	Name string
}

// 通常の埋め込み構造体
type BannedUser struct {
	User
	BannedReason string
}

// ポインタ型の埋め込み構造体
type PBannedUser struct {
	*User
	BannedReason string
}

func main() {
	user := BannedUser{User{1, "oinume"}, "Too many spam"}
	fmt.Println(user.Name)

	puser := PBannedUser{&User{1, "oinume"}, "Too many spam"}
	// なぜか puser.User.Name でも参照可能
	fmt.Println(puser.User.Name)
}

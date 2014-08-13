package main

import (
	"fmt"
)

type User struct {
	Id int
	Name string
}

func (u *User) String() string {
	return fmt.Sprintf("%d:%s", u.Id, u.Name)
}

// 通常の埋め込み構造体
type BannedUser struct {
	User
	BannedReason string
}

func (u *BannedUser) GetBannedReason() string {
	return u.BannedReason
}

// ポインタ型の埋め込み構造体
type PBannedUser struct {
	*User
	BannedReason string
}

func (u *PBannedUser) GetBannedReason() string {
	return u.BannedReason
}

func main() {
	user := BannedUser{User{1, "oinume"}, "Too many spam"}
	fmt.Println(user.Name)

	puser := PBannedUser{&User{1, "oinume"}, "Too many spam"}
	// User.Nameが参照できる
	fmt.Println(puser.Name)
	// なぜか puser.User.Name でも参照可能
	fmt.Println(puser.User.Name)

	// Mixin: User.String()を呼び出す
	fmt.Println(puser.String())

	var u *User = &PBannedUser{&User{1, "oinume"}, "Too many spam"}
}

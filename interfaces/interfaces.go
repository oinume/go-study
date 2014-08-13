package main

import "fmt"

// 「JSON化できる」インターフェース
type JSONable interface {
	JSON() string
}

type User struct {
	Id   int
	Name string
}

// JSON()メソッドを実装
func (s *User) JSON() string {
	return fmt.Sprintf(`{ "Id": %d, "Name": "%s" }`, s.Id, s.Name)
}

type AdminUser struct {
	User
	Admin bool
}

// User.JSON()をオーバーライド
func (s *AdminUser) JSON() string {
	return fmt.Sprintf(`{ "Id": %d, "Name": "%s", "Admin": %v }`, s.Id, s.Name, s.Admin)
}


func main() {
	// JSONable を実装しているのでJSONable型に代入できる
	var user JSONable = &User{1, "oinume"}
	fmt.Println(user.JSON())

	// AdminUserもJSONableを実装している
	var adminUser JSONable = &AdminUser{User{0, "admin"}, true}
	fmt.Println(adminUser.JSON())

	// Type assertion
	jsonable, ok := adminUser.(JSONable)
	if ok {
		fmt.Printf("JSON(): %s\n", jsonable.JSON())
	} else {
		fmt.Printf("Not JSONable\n")
	}
}

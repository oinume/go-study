package main

import "fmt"

type JSONable interface {
	ToJSON() string
}

type User struct {
	Id   int
	Name string
}

func (s *User) ToJSON() string {
	return fmt.Sprintf(`{ "Id": %d, "Name": "%s" }`, s.Id, s.Name)
}

type AdminUser struct {
	User
	Admin bool
}

// オーバーライド
func (s *AdminUser) ToJSON() string {
	return fmt.Sprintf(`{ "Id": %d, "Name": "%s", "Admin": %v }`, s.Id, s.Name, s.Admin)
}


func main() {
	// JSONable を実装しているので代入JSONableに代入できる
	var user JSONable = &User{1, "oinume"}
	fmt.Println(user.ToJSON())

	// AdminUserもJSONableを実装している
	var adminUser JSONable = &AdminUser{User{0, "admin"}, true}
	fmt.Println(adminUser.ToJSON())
}

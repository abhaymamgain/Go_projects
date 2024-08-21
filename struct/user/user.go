package user

import "fmt"

type Admin struct {
	email    string
	password string
	User     User
}
type User struct {
	Name string
	Age  int
}

func NewAdmin(email string, password string, name string, age int) Admin {
	return Admin{email: email, password: password, User: CreateUser(name, age)}
}
func CreateUser(name string, age int) User {
	return User{Name: name, Age: age}
}
func (u *User) PrintUser() {
	fmt.Printf("Name: %v\nAge: %v\n", u.Name, u.Age)
}

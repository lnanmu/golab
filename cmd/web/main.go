package main 

import "fmt"

type User struct {
	name string
	age int 
}

func NewUser(name string, age int ) User {
	return  User{
		name: name,
		age: age,
	}
}

func main() {
	u := make([]User,0)
	u1 := NewUser("hong", 12)
	u2 := NewUser("ming", 18)
	u = append(u, u1)
	u = append(u, u2)
	for _, v := range u {
		if v.age > 15 {
			fmt.Println(v)
		}
	}

}
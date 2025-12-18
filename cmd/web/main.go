package main

import "fmt"

// 数组、切片、map、结构体、通道、指针、接口、函数
// 通道、指针、接口、函数 剩余这些
type user struct {
	id   int
	name string
	age  int
}

// 两种方式的对比，直接修改的方式是复制出来一份，修改的是副本，原数据不变
func main() {
	u1 := user{id: 1, name: "张三", age: 20}
	u2 := u1
	fmt.Println("原始数据u1", u1, "原始数据u2", u2)

	u1.age = 30
	fmt.Println("原始数据u1", u1, "原始数据u2", u2)
	// 通过指针的方式修改，修改的是内存地址，指向的是原始数据，
	u1ptr := &u1
	u1ptr.name = "李四"
	fmt.Println("原始数据u1", u1, "原始数据u2", u2)
	// 结构体本身没有遍历，只有在多个元素的数据下遍历才有意义。
	userlist := []user{u1, u2}
	fmt.Println(userlist)
	Userlist := [2]user{u1, u2}
	fmt.Println(Userlist)

	for i, v := range userlist {
		fmt.Println(i, v)
	}
}

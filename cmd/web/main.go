package main

import "fmt"

// 数组、切片、map、结构体、通道、指针、接口、函数


var s []int = []int{1, 2, 3, 4, 5}
var index int = 3
func hi() {
	s = append(s[:index],s[index+1:]...)

}

func main() {
	fmt.Println("改变前的数组：", s)
	hi()
	fmt.Println("改变后的数组", s)

}

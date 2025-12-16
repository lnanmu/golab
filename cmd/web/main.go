package main

import (
	"fmt"
	"unicode/utf8"
)

var (
	name      string = "jiusi"
	age       int    = 18
	bf        string = "yangchenguang"
	isMarried bool   = false
)

func main() {
	fmt.Println(utf8.RuneCountInString(name))
	fmt.Println(len(name))
	for i, v := range bf {
		fmt.Printf("类型：%T, 索引：%d, 值：%c", bf, i, v)
	}
	fmt.Printf("%d", 10+20)
	fmt.Printf("%d", 100-50)
	fmt.Printf("%d", 3*3)
	fmt.Printf("%d", 20/5)
	fmt.Println("好了，开始要注意了哦")
	fmt.Printf("%d", 9%2)
}
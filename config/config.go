package config

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

func hi() {
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
	fmt.Sprintf(name)
	fmt.Sprintln(isMarried)
}

func hello() {
	son := [4]string{"dong", "nan", "xi", "bei"}
	num := [4]int{100, 200, 300, 400}
	lan := [...]string{"go", "java", "python"}
	for k, v := range son {
		fmt.Println("son: k, v", k, v)
	}

	for i, v := range num {
		fmt.Println(i, v)
	}
	fmt.Println(len(lan))
	for i, k := range lan {
		if i <= len(lan) {
			fmt.Println("missing you", k)
			break
		}
	}
}

//切片的常见应用:创建、增、删、改、查、复制

func qiepian() {
	//创建切片
	s := make([]int, 0, 10)
	for i := 10; i <= 100; i += 10 {
		// 增加元素
		s = append(s, i)
	}
	fmt.Println(s)
	// 更改元素
	for i, v := range s {
		if v == 30 {
			s[i] = 30000
		}
	}
	fmt.Println(s)

	// 删除元素
	y := cap(s)

	s1 := make([]int, 0, y)

	for _, v := range s {
		if v == 30000 {
			continue
		}
		s1 = append(s1, v)

	}
	s = s1
	fmt.Println(s)
	//查询
	for k, v := range s {
		fmt.Printf("k:%d, v:%d\n", k, v)
	}

	s2 := make([]int, 2, 2)
	// copy的结果是有dts的长度决定的
	copy(s2, s)
	fmt.Println(s2, s)
}

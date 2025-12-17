package main

import "fmt"

// 数组、切片、map、结构体、通道、指针、接口、函数

//切片的常见应用:创建、增、删、改、查、复制

func main() {
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

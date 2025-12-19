package main

import "fmt"

func news(s []int) map[int]bool {
	m := make(map[int]bool)
	// s1 := []int{}
	for _, v := range s {
		m[v] = true
	}
	fmt.Println(m)

	return m
}

// map中的不能重复是只map的key必须是唯一的！！！！ 还是要了解底层的远离，否则不知道怎么写

func news1(m map[int]bool) []int {
	s := []int{}
	for k, _ := range m {
		s = append(s, k)
	}
	return s
}

func main() {
	s := []int{3, 5, 3, 7, 5, 9, 7, 10}
	s1 := news(s)
	fmt.Println(s1)
}

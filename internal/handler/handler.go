package handler

import "fmt"

func Hello() {
	fmt.Println("Hello, This is handler")
}

var s []int = []int{5, 12, 9, 27, 18, 3, 22}

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func max10(s []int) []int {
	s1 := make([]int, len(s), cap(s))
	for _, v := range s {
		if v > 10 {
			s1 = append(s1, v)
		}
	}
	return s1
}

// 找到最小的
func min(s []int) int {
	minnumber := s[0]
	for _, v := range s {
		if v < minnumber {
			minnumber = v
		}
	}
	return minnumber
}

// 找到最大的
func max(s []int) int {
	maxnumber := s[0]
	for _, v := range s {
		if v > maxnumber {
			maxnumber = v
		}
	}
	return maxnumber
}

//找到最小的和最大的，这两个函数都有一种思想，就是初始化条件的思想，先假设第一个就是最小的或者最大的，然后遍历数组。如果存在更小的那么就更新最小的，如果没有更小的就不更新，就继续遍历，直到遍历结束。 每一次都是经过对比过的 每一次遍历都是在上一次的对比的结果。。

func mainkk() {
	he := sum(s)

	maxn := max(s)

	s2 := max10(s)

	minn := min(s)

	fmt.Println(he, maxn, s2, minn)
}

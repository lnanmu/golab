package service

import "fmt"

func serviceinfo() {
	fmt.Println("hi")
}




var a [7]int = [7]int{18, 33, 27, 9, 33, 12, 27}

// func max(a [7]int) int {
// 	x := a[0]
// 	for _, v := range a {
// 		if v > x {
// 			x = v

// 		}

// 	}
// 	return x
// }

//	func index(x int) int {
//		y := 0
//		for i := 0; i < len(a); i++ {
//			if a[i] == x {
//				y = i
//				break
//			}
//		}
//		return y
//	}
// 一个简单的小练习，索引和数字，和解耦的思想，range和索引的使用，各种各的用处。先设计出来有几个变量来存储结果，
func findindexandnumber(a []int) (index, number int) {
	index = 0
	number = a[index]
	for i := 1; i < len(a); i++ {
		if a[i] > number {
			index = i
			number = a[i]
		}
	}
	return
}

func mainhhhhh() {
	index, number := findindexandnumber(a[:])
	fmt.Println(index, number)
	
}

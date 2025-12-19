package main 

import "fmt"

type Student struct {
	ID int
	Name string 
	Score []int // 是一个切片，存放成绩

}

var s1 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 返回学生的三科目的平均分数
func scoreing(s Student) int{
	sum := 0
	for _, v := range s.Score{
		sum += v
	}
	so := sum / len(s.Score)
	return so
}
// 返回一个切片中所有的偶数，创建一个新的切片
func oushu(s []int) []int {
	s2 := make([]int,0, len(s))
	for _, v := range s {
		if v % 2 ==0 {
			s2 = append(s2, v)
		}
	}
	return s2
}

// 查找20以内所有的质数字 
func zhishu() []int {
	s := make([]int,0, 20)
	// 外层控制，是寻找质数的范围
	for i := 2; i <= 20; i++ {
		// 内层控制，判断i是否为质数
		flag := true
		// 因为如果是i偶数，那么出了2以外，一定不是质数，那么就只剩下奇数的部分了，奇数部分只需要判断到i的一半就可以了，因为大于一半的数字不可能是因子，因为如果是因子，那么商就会是小于2，这样就不是整除了
		// 如果出现了能整除的数字，那么就不是质数，结束内层循环就可以
		for j := 2; j <= i/2; j++ {
			if i % j ==0 {
				flag = false
				break
			}
		}
		// 外层是一次一次的循环，所以使用标志控制是否添加。
		if flag {
			s = append(s, i)
		}
	}

	return s
}
// 我改如何理解flag这种操作，这里引入一个状态追的思想，因为整个程序的运行过程是动态的，我们需要根据不同的状态来做出不同的操作，所以要通过一个表达式来进行判断这种状态是否满足我们的要求，所以我们引入了一个标记位，我们通过判断标记位的的值来判断我们改进行如何操作，有了标记位之后，就可以有一个比较清晰的表达式，这样更简单，好好体会一下这个思想，标记位的变化是随着数据或者状态的变化而变化。
func zhishu2 () []int  {
	s := make([]int, 0, 20)
	for i := 37; i <= 200; i++ {
		flag := true
		for j := 2; j <= i/2; j++ {
			if i % j ==0 {
				flag = false
				break
			}
		}
		if flag {
			s = append(s, i)
		}
	}
	return s
}

// ++的使用 记住了
func zifununber (s string) map[rune]int {
	m := make(map[rune] int)
	for _, v := range s {
		m[v] ++
	}
	return m
} 



func main() {
	s := zhishu()
	fmt.Println(s)
	s2 := zhishu2()
	fmt.Println(s2)
	m := zifununber("abcabc")
	fmt.Println(m)

}
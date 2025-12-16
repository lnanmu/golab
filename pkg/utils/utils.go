package utils

import "fmt"

var name string = "kata"
var age int = 18

const Pi float32 = 3.14

const (
	paywx = iota
	payzfb
	payxj
	payyhk
)

func sum(a, b int) int {
	return a + b
}

func calareayuan(x float32) float32 {
	return x * x * Pi
}

func calzhouchang(x, y float32) float32 {
	return (x + y) * 2
}

func clazhouchangyuan(x float32) float32 {
	return 2 * Pi * x
}

func hello() {
	a := 10
	b := 20
	c := sum(a, b)
	fmt.Println(calareayuan(2))
	fmt.Println(calzhouchang(1, 2))
	fmt.Println(calareayuan(3))
	fmt.Println(c)
	fmt.Println(name, age)
	p := 0
	switch p {
	case paywx:
		fmt.Println("weixin")
	case payzfb:
		fmt.Println("支付宝")
	}
}

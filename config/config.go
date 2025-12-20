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

func maip() {
	// 创建一个map
	m := make(map[string]int, 3)
	m["chang"] = 18
	m["kuan"] = 20
	fmt.Println(m)

	//新增一个元素
	m["gao"] = 30
	fmt.Println(m)

	// 修改一个元素
	m["chang"] = 10
	fmt.Println(m)

	// 查一个元素的第一种方式，不是很推荐
	fmt.Println(m["yuan"])
	//查找元素的第二种方式，推荐,这里的ok是一个bool,存在就是true，不存在就是false
	v, ok := m["yuan"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("sorry bucunzai", ok)
	}
	// 删除元素
	delete(m, "kuan")
	fmt.Println("查看删除之后的元素", m)

	// 遍历元素
	for k, v := range m {
		fmt.Printf("key:%s, vaule:%d\n", k, v)
	}
}






// 数组、切片、map、结构体、通道、指针、接口、函数
// 通道、指针、接口、函数 剩余这些
type user struct {
	id   int
	name string
	age  int
}
// 要习惯有中间过程。
// 两种方式的对比，直接修改的方式是复制出来一份，修改的是副本，原数据不变
func mainjjj() {
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

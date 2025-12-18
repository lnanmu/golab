package db


var s []int = []int{1, 2, 3, 4, 5}
var index int = 3
// 删除数组里指定位置的元素，注意append的用法，第二个位置上是元素，所以使用...拆解切片
func hi() {
	s = append(s[:index],s[index+1:]...)

}
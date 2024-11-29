package main

import (
	"fmt"
)

type MyStruct struct {
	Data []int
	age  int
}

func modifyStruct(s MyStruct) {
	s.Data[0] = 42 // 修改引用类型字段的底层数据
	s.age = 33     // 修改普通类型不会修改原始数据
}

func main() {
	s := MyStruct{Data: []int{1, 2, 3}}
	modifyStruct(s)
	fmt.Println(s.Data) // 输出: [42, 2, 3]
}

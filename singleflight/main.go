package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

func getData(id int64) string {
	fmt.Println("query...")
	time.Sleep(10 * time.Second) // 模拟一个比较耗时的操作
	return "sealball~~~~"
}

func main() {
	g := new(singleflight.Group)

	// 第1次调用
	go func() {
		v1, _, shared := g.Do("getData", func() (interface{}, error) {
			ret := getData(1)
			return ret, nil
		})
		fmt.Printf("1st call: v1:%v, shared:%v\n", v1, shared)
	}()

	time.Sleep(2 * time.Second)

	// 第2次调用（第1次调用已开始但未结束）
	v2, _, shared := g.Do("getData", func() (interface{}, error) {
		ret := getData(1)
		return ret, nil
	})
	fmt.Printf("2nd call: v2:%v, shared:%v\n", v2, shared)
}

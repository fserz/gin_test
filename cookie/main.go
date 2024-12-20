package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie") // 获取Cookie
		if err != nil {
			cookie = "NotSet"
			// 设置Cookie
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})
	router.Run()
}

//func maxSpending(values [][]int) int64 {
//	m, n := len(values), len(values[0])
//	pq := &Heap{}
//	for i := 0; i < m; i++ {
//		//heap.Push(pq, []int{values[i][n-1], i, n - 1)
//		heap.Push(pq)
//	}
//	return 11
//}
//
//type Heap [][]int
//
//func (h Heap) Len() int {
//	return len(h)
//}
//
//func (h Heap) Less(i, j int) bool {
//	if h[i][0] == h[j][0] {
//		return h[i][1] < h[j][1]
//	}
//	return h[i][0] < h[j][0]
//}
//
//func (h Heap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *Heap) Push(x interface{}) {
//	*h = append(*h, x.([]int))
//}
//
//func (h *Heap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}

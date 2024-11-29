package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func sayHello(w http.ResponseWriter, r *http.Request) {
//	file, _ := os.ReadFile("./hello.txt")
//	_, _ = fmt.Fprintln(w, string(file))
//}
//
//func main() {
//	http.HandleFunc("/hello", sayHello)
//	err := http.ListenAndServe(":9095", nil)
//	if err != nil {
//		fmt.Printf("http serve failed, err:%v\n", err)
//		return
//	}
//}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello golang!",
	})
}

// 使用gin
func main() {
	r := gin.Default()
	r.GET("/hello", sayHello)
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.PATCH("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PATCH",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	// 启动
	r.Run()
}

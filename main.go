package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
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

func sayHi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hi seal~",
	})
}

// 使用gin

func main() {
	r := gin.Default()

	r.Any("/allMethod", func(c *gin.Context) {
		logrus.Info("seal ball~~~")
	})

	// 路由组
	userGroup := r.Group("/user").Use(ginBodyLogMiddleware())
	userGroup.GET("/hello", sayHello)
	userGroup.GET("/hi", sayHi)

	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
		//c.HTML()
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

	r.GET("/test", func(c *gin.Context) {
		// 重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	})

	r.LoadHTMLFiles("./404.html") // 加载模板文件

	// 为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.GET("/test2", ginBodyLogMiddleware(), func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println("get name from context: ", name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello ball ball ~~~~",
		})
	})

	// 启动
	r.Run(":9090")
}

//func main() {
//	g := errgroup.Group{}
//	var urls = []string{
//		"https://github.com",
//		"http://www.liwenzhou.com",
//		"http://www.baidu.com",
//		"http://www.yixieqitawangzhi.com",
//		"http://www.yixieqitaw",
//	}
//
//	for _, url := range urls {
//		url := url // 注意此处声明新的变量
//		// 启动一个goroutine去获取url内容
//		g.Go(func() error {
//			resp, err := http.Get(url)
//			if err == nil {
//				fmt.Printf("获取%s成功\n", url)
//				resp.Body.Close()
//			}
//			return err // 返回错误
//		})
//	}
//	if err := g.Wait(); err != nil {
//		// 处理可能出现的错误
//		fmt.Println(err)
//	}
//	fmt.Println("所有goroutine运行结束功")
//}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("name", "seal")
		// 使用 bodyLogWriter 记录响应内容
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 检查状态码并记录响应体
		statusCode := c.Writer.Status()
		if statusCode == 200 {
			logrus.Errorf("Response body: %s", blw.body.String())
		}
	}
}

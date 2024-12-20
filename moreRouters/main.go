package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//借助errgroup.Group或者自行开启两个goroutine分别启动两个服务
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	//var wg sync.WaitGroup
	//errorChan := make(chan error, 2) // 用于接收 Goroutine 中的错误
	//
	//// 启动 server01
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	log.Println("Starting server01 on :8080")
	//	if err := server01.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		errorChan <- err
	//	}
	//}()
	//
	//// 启动 server02
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	log.Println("Starting server02 on :8081")
	//	if err := server02.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		errorChan <- err
	//	}
	//}()
	//
	//// 等待 Goroutine 结束或错误发生
	//go func() {
	//	wg.Wait()
	//	close(errorChan)
	//}()
	//
	//// 检查错误
	//for err := range errorChan {
	//	log.Fatalf("Error occurred: %v", err)
	//}
	//
	//log.Println("Both servers stopped.")
}

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2. 解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v", err)
		return
	}
	name := "seal ball"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err :%v", err)
		return
	}

	// 3. 渲染模板

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP serve failed err: ", err)
		return
	}
}

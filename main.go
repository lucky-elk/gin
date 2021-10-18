package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, r)
	//解析模板
	t, err := template.ParseFiles("./hello.temp")
	fmt.Println(t, err)
	if err != nil {
		fmt.Println("parse template failed,err:%v", err)
		return
	}
	//渲染模板
	err = t.Execute(w, "小王")
	if err != nil {
		fmt.Println("render template failed,err:%v", err)
		return
	}
}

func main() {
	//http.HandleFunc("/",sayHello)
	//err := http.ListenAndServe(":9000",nil)
	//fmt.Println(err)
	//if err != nil{
	//	fmt.Println("HTTP server start failed,err:%v",err)
	//	return
	//}

	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r := gin.Default()
	//r.GET("/hello", func(c *gin.Context) {
	//	// c.JSON：返回JSON格式的数据
	//	c.JSON(200, gin.H{
	//		"message": "Hello world!",
	//	})
	//})
}

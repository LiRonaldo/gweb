package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
实现ServeHTTP 方法，定义属于自己的请求处理器 handler。
*/
type Engine struct {
}

//重写ServeHTTP方法
func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
	case "/hellogWeb":
		for k, v := range req.Header {
			fmt.Fprintf(resp, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(resp, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	//创建Engine
	engine := new(Engine)
	//将engine 对象放进去。会自动调用engine 重写的ServeHTTP
	log.Fatal(http.ListenAndServe(":8888", engine))
}

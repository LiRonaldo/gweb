package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//处理底层请求/
	http.HandleFunc("/", indexHandler)
	//处理/hellogweb请求
	http.HandleFunc("/hellogWeb", helloGwebHandler)
	//监听9999端口//fatal 打印，并且退出。但是http.ListenAndServe会一直监听不会执行完,所以log.Fatal 不出问题的话永远执行不到。
	log.Fatal(http.ListenAndServe(":9999", nil))
}
func indexHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
}
func helloGwebHandler(resp http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(resp, "Header[%q] = %q\n", k, v)
	}
}

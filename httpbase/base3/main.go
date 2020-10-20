package main

import (
	"fmt"
	"gweb/httpbase/gweb"
	"net/http"
)

func main() {
	r := gweb.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	r.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})
	r.RUN(":8888")
}

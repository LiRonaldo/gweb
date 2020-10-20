package gweb

import (
	"fmt"
	"net/http"
)

/**
定义底层的方法
*/
type HandlerFunc func(http.ResponseWriter, *http.Request)

/**
定义engine对象
*/
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: map[string]HandlerFunc{}}
}
func (engine *Engine) AddRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

//get请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.AddRoute("GET", pattern, handler)
}

//post请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.AddRoute("POST", pattern, handler)
}

//启动
func (engine *Engine) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

//engine 重写了ServeHTTP
func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(resp, req)
	} else {
		fmt.Fprintf(resp, "404 NOT FOUND: %s\n", req.URL)
	}
}

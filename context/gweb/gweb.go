package gweb

import "net/http"

type HandlerFunc func(*Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}
func (engine *Engine) AddRouter(method string, patten string, handler HandlerFunc) {
	engine.router.AddRouter(method, patten, handler)
}
func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	engine.AddRouter("GET", pattern, handlerFunc)
}
func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	engine.AddRouter("POST", pattern, handlerFunc)
}
func (engine *Engine) RUN(addr string) error {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	c := NewContext(resp, req)
	engine.router.Handle(c)
}

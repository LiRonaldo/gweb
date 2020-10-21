package gweb

import (
	"log"
	"net/http"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}
func (r *Router) AddRouter(method string, patten string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, patten)
	key := method + "-" + patten
	r.handlers[key] = handler
}
func (r *Router) Handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

package gee

import "net/http"

type HandlerFunc func(*Context)

type (
	RouteGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouteGroup
		engine      *Engine
	}
	Engine struct {
		*RouteGroup
		Router *Router
		group  []*RouteGroup
	}
)

func New() *Engine {
	engine := &Engine{Router: NewRouter()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.group = []*RouteGroup{engine.RouteGroup}
	return engine
}

func (g *RouteGroup) Group(prefix string) *RouteGroup {
	engine := g.engine
	newGroup := &RouteGroup{prefix: g.prefix + prefix, parent: g, engine: engine}
	engine.group = append(engine.group, newGroup)
	return newGroup
}
func (g *RouteGroup) AddRouter(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	g.engine.Router.addRoute(method, pattern, handler)
}

func (g *RouteGroup) GET(pattern string, handler HandlerFunc) {
	g.AddRouter("GET", pattern, handler)
}
func (group *RouteGroup) POST(pattern string, handler HandlerFunc) {
	group.AddRouter("POST", pattern, handler)
}

func (engine *Engine) RUN(addr string) error {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	c := NewContext(resp, req)
	engine.Router.handle(c)
}

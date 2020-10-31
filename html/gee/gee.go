package gee

import (
	"html/template"
	"net/http"
	"path"
	"strings"
)

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
		Router        *Router
		group         []*RouteGroup
		htmlTemplates *template.Template
		funcMap       template.FuncMap
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
	var middlewares []HandlerFunc
	for _, group := range engine.group {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(resp, req)
	c.handlers = middlewares
	c.engine = engine
	engine.Router.handle(c)
}
func (group *RouteGroup) Use(middlewars ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewars...)
}
func (group *RouteGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(context *Context) {
		file := context.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			context.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(context.Resp, context.Req)
	}
}

func (group *RouteGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	aa := template.New("")
	ss := aa.Funcs(engine.funcMap)
	bb, err := ss.ParseGlob(pattern)
	engine.htmlTemplates = template.Must(bb, err)
}

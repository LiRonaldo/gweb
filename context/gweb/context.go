package gweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
编写context，上下文。里边集成req，resp 等等。并集成josn方法。和基本方法。
*/
type H map[string]interface{}

type Context struct {
	Resp       http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

func NewContext(resp http.ResponseWriter, req *http.Request) *Context {
	return &Context{Resp: resp, Req: req, Path: req.URL.Path, Method: req.Method}
}
func (context *Context) PostForm(key string) string {
	return context.Req.FormValue(key)
}
func (context *Context) Query(key string) string {
	return context.Req.URL.Query().Get(key)
}
func (context *Context) Status(code int) {
	context.StatusCode = code
	context.Resp.WriteHeader(code)
}
func (context *Context) SetHeader(key string, value string) {
	context.Resp.Header().Set(key, value)
}
func (context *Context) String(code int, format string, values ...interface{}) {
	context.SetHeader("Content-Type", "text/plain")
	context.Status(code)
	context.Resp.Write([]byte(fmt.Sprintf(format, values)))
}
func (context *Context) JSON(code int, obj interface{}) {
	context.SetHeader("Context-Type", "application/json")
	context.Status(code)
	encoder := json.NewEncoder(context.Resp)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.Resp, err.Error(), 500)
	}
}
func (context *Context) Date(code int, date []byte) {
	context.Status(code)
	context.Resp.Write(date)
}
func (context *Context) HTML(code int, html string) {
	context.SetHeader("Content-Type", "text/html")
	context.Status(code)
	context.Resp.Write([]byte(html))
}

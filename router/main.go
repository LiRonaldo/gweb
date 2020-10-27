package main

import (
	"gweb/router/base"
	"net/http"
)

func main() {
	r := base.New()
	r.GET("/", func(c *base.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *base.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	/*r.GET("/hello/:name", func(c *base.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})*/

	/*r.GET("/assets/*filepath", func(c *base.Context) {
		c.JSON(http.StatusOK, base.H{"filepath": c.Param("filepath")})
	})*/

	r.RUN(":9999")
}

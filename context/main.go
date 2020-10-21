package main

import (
	"gweb/context/gweb"
	"net/http"
)

/**
新增 engine 和 router 和context 的结合。设计还是比较巧妙的。
*/
func main() {
	r := gweb.New()
	r.GET("/", func(c *gweb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.POST("/hello", func(c *gweb.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gweb.Context) {
		c.JSON(http.StatusOK, gweb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.RUN(":9999")
}

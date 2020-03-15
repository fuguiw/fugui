package main

import (
	"fugui"
	"net/http"
)

func main() {
	r := fugui.New()

	r.GET("/", func(c *fugui.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *fugui.Context) {
		// expect /hello?name=haha
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *fugui.Context) {
		// expect /hello/haha
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *fugui.Context) {
		c.JSON(http.StatusOK, fugui.H{"filepath": c.Param("filepath")})
	})

	r.Run(":4444")
}

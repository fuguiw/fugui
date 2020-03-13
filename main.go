package main

import (
	"fugui"
	"net/http"
)

func main() {
	r := fugui.New()

	r.GET("/", func(c *fugui.Context) {
		c.HTML(http.StatusOK, "<h1>hello fugui</h1>")
	})

	r.GET("/hello", func(c *fugui.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *fugui.Context) {
		c.JSON(http.StatusOK, fugui.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":4444")
}

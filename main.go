package main

import (
	"fmt"
	"fugui"
	"html/template"
	"log"
	"net/http"
	"time"
)

func onlyForV2() fugui.HandlerFunc {
	return func(c *fugui.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("%d %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

type student struct {
	Name string
	Age  int8
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := fugui.New()
	r.Use(fugui.Logger())

	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "haha", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *fugui.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *fugui.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", fugui.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *fugui.Context) {
			// expect /hello?name=wuha
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *fugui.Context) {
			// expect /hello/wuha
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *fugui.Context) {
			c.JSON(http.StatusOK, fugui.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":4444")
}

package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func render_template(c *gin.Context, tmpl *template.Template, data interface{}) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(c.Writer, data); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func render_index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	router := gin.New()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", render_index)

	router.Run(":3000")
}

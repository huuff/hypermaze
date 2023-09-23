package main

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
  app := newApplication()

  r := gin.Default()
  r.SetFuncMap(template.FuncMap {
    "directionToString": directionToString,
  })
  r.LoadHTMLGlob("templates/**/*")

  app.initRouter(r)

  r.Run()
}

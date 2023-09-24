package main

import (
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
  app := newApplication()

  r := gin.Default()
  r.SetFuncMap(template.FuncMap {
    "directionToString": directionToString,
    "toLower": strings.ToLower,
  })
  r.LoadHTMLGlob("templates/**/*")

  app.initRouter(r)

  r.Run()
}

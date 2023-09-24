package main

import (
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/cmd/web/util"
)

func main() {
  app := newApplication()

  r := gin.Default()
  r.SetFuncMap(template.FuncMap {
    "directionToString": util.DirectionToString,
    "toLower": strings.ToLower,
    "directionToKey": util.DirectionToKey,
  })
  r.LoadHTMLGlob("templates/**/*")

  app.initRouter(r)

  r.Run()
}

package main

import (
  "github.com/gin-gonic/gin"
  "text/template"
  "xyz.haff/maze/cmd/web/util"
  "strings"
)

func initTemplates(router *gin.Engine) {
  router.SetFuncMap(template.FuncMap {
    "directionToString": util.DirectionToString,
    "toLower": strings.ToLower,
    "directionToKey": util.DirectionToKey,
  })
  router.LoadHTMLGlob("templates/**/*")
}

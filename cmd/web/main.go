package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
  app := newApplication()

  r := gin.Default()
  r.HTMLRender = newRenderer()
  app.initRouter(r)

  r.Run()
}

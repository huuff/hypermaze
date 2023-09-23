package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
  app := newApplication()

  r := gin.Default()
  r.LoadHTMLGlob("templates/**/*")
  app.initRouter(r)

  r.Run()
}

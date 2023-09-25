package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
  mazes := generateMazes()

  r := gin.Default()
  initTemplates(r)
  initRouter(r, mazes)

  r.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/cmd/web/handlers"
)

func (app application) index(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html.gotmpl", gin.H {
    "Mazes": app.mazes,
  })
}

func (app application) initRouter(router *gin.Engine) http.Handler {
  router.GET("/", app.index)

  mazeHandler := handlers.MazeHandler { Mazes: app.mazes }

  router.GET("/mazes/:level/ascii", mazeHandler.Ascii)
  router.GET("/mazes/:level", mazeHandler.Maze)

  roomHandler := handlers.RoomHandler { Mazes: app.mazes }

  // TODO: I'd like to have :x,:y but gin-gonic doesn't allow it... what do I do?
  router.GET("/mazes/:level/room/:x/:y", roomHandler.Room)
  router.GET("/mazes/:level/room/:x/:y/minimap", roomHandler.Minimap)

  router.Static("/static", "./static")
  
  return router
}

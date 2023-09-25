package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/cmd/web/handlers"
)


func (app application) initRouter(router *gin.Engine) http.Handler {
  mazeHandler := handlers.MazeHandler { Mazes: app.mazes }
  roomHandler := handlers.RoomHandler { Mazes: app.mazes }
  mazeListHandler := handlers.MazeListHandler { Mazes: app.mazes }

  router.Static("/static", "./static")
  router.GET("/", func(c *gin.Context) {
    c.Redirect(http.StatusFound, "/mazes")
  })

  router.GET("/mazes", mazeListHandler.MazeList)
  router.GET("/mazes/:level/ascii", mazeListHandler.Ascii)

  router.GET("/mazes/:level", mazeHandler.Maze)

  // TODO: I'd like to have :x,:y but gin-gonic doesn't allow it... what do I do?
  router.GET("/mazes/:level/room/:x/:y", roomHandler.Room)
  router.GET("/mazes/:level/room/:x/:y/minimap", roomHandler.Minimap)

  
  return router
}

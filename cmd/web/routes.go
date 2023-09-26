package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/cmd/web/handlers"
	"xyz.haff/maze/pkg/maze"
)


func initRouter(router *gin.Engine, mazes []*maze.Maze) {
  mazeHandler := handlers.MazeHandler { Mazes: mazes }
  roomHandler := handlers.RoomHandler { Mazes: mazes }
  mazeListHandler := handlers.MazeListHandler { Mazes: mazes }

  router.Static("/static", "./static")
  router.GET("/", func(c *gin.Context) {
    c.Redirect(http.StatusFound, "/mazes")
  })

  router.GET("/mazes", mazeListHandler.MazeList)
  router.GET("/mazes/:level/ascii", mazeListHandler.Ascii)

  router.GET("/mazes/:level", mazeHandler.Maze)

  router.GET("/mazes/:level/room/:x/:y", roomHandler.Room)
  router.GET("/mazes/:level/room/:x/:y/minimap", roomHandler.Minimap)

}

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/cmd/web/handlers"
	"xyz.haff/maze/cmd/web/util"
	"xyz.haff/maze/pkg/ascii"
)

func (app application) index(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html.gotmpl", gin.H {
    "Mazes": app.mazes,
  })
}

func (app application) mazeAscii(c *gin.Context) {
  level, err := strconv.Atoi(c.Param("level"))
  if err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

  if level >= len(app.mazes) {
    c.String(http.StatusNotFound, "")
    return
  }

  maze := app.mazes[level]

  c.HTML(http.StatusOK, "maze-ascii-view.html.gotmpl", gin.H {
    "View": ascii.View(*maze, nil),
  })
}

func (app application) maze(c *gin.Context) {
  level, err := strconv.Atoi(c.Param("level"))

  if err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

  if level >= len(app.mazes) {
    c.String(http.StatusNotFound, "")
    return
  }

  maze := app.mazes[level]

  data := gin.H {
    "Level": level,
    "View": ascii.View(*maze, nil),
    "Maze": maze,
  }

  if util.IsHxRequest(c) {
    c.HTML(http.StatusOK, "maze-partial.html.gotmpl", data)
  } else {
    c.HTML(http.StatusOK, "maze.html.gotmpl", data)
  }
}

func (app application) initRouter(router *gin.Engine) http.Handler {
  router.GET("/", app.index)
  router.GET("/mazes/:level/ascii", app.mazeAscii)
  router.GET("/mazes/:level", app.maze)

  roomHandler := handlers.RoomHandler { Mazes: app.mazes }

  // TODO: I'd like to have :x,:y but gin-gonic doesn't allow it... what do I do?
  router.GET("/mazes/:level/room/:x/:y", roomHandler.Room)
  router.GET("/mazes/:level/room/:x/:y/minimap", roomHandler.Minimap)

  router.Static("/static", "./static")
  
  return router
}

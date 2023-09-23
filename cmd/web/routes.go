package main

import (
	"net/http"
	"strconv"

	"xyz.haff/maze/pkg/ascii"
	"xyz.haff/maze/pkg/grid"
  "github.com/gin-gonic/gin"
)

func (app application) index(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html.gotmpl", gin.H {
    "Mazes": app.mazes,
  })
}

func (app application) minimap(c *gin.Context) {
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

  c.HTML(http.StatusOK, "minimap.html.gotmpl", gin.H {
    "Minimap": ascii.View(*maze, nil),
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
    "Minimap": ascii.View(*maze, nil),
    "Maze": maze,
  }

  if c.GetHeader("HX-Request") == "" {
    c.HTML(http.StatusOK, "maze.html.gotmpl", data)
  } else {
    c.HTML(http.StatusOK, "maze-partial.html.gotmpl", data)
  }
}

type RoomUri struct {
  Level int `uri:"level"`
  X int `uri:"x"`
  Y int `uri:"y"`
}

func (app application) room(c *gin.Context) {
  var params RoomUri
  if err := c.ShouldBindUri(&params); err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

  if params.Level >= len(app.mazes) {
    c.String(http.StatusNotFound, "")
    return
  }

  maze := app.mazes[params.Level]

  point := grid.Point { X: params.X, Y: params.Y }
  room, ok := maze.Rooms[point]

  if !ok {
    c.String(http.StatusNotFound, "")
    return
  }

  c.HTML(http.StatusOK, "room-partial.html.gotmpl", gin.H {
    "Room": room,
    "Level": params.Level,
  })
}

//func (app application) roomMinimap(c *gin.Context) {
  
//}

func (app application) initRouter(router *gin.Engine) http.Handler {
  router.GET("/", app.index)
  router.GET("/mazes/:level/minimap", app.minimap)
  router.GET("/mazes/:level", app.maze)
  // TODO: I'd like to have :x,:y but gin-gonic doesn't allow it... what do I do?
  router.GET("/mazes/:level/room/:x/:y", app.room)
  //router.GET("/mazes/:level/room/:x/:y/minimap", app.roomMinimap)

  router.Static("/static", "./static")
  
  return router
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/pkg/ascii"
	"xyz.haff/maze/pkg/grid"
	"xyz.haff/maze/pkg/maze"
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

  if isHxRequest(c) {
    c.HTML(http.StatusOK, "maze-partial.html.gotmpl", data)
  } else {
    c.HTML(http.StatusOK, "maze.html.gotmpl", data)
  }
}

type RoomUri struct {
  Level int `uri:"level"`
  X int `uri:"x"`
  Y int `uri:"y"`
}

func (app application) findRoom(roomUri RoomUri) (*maze.Maze, *maze.Room, error) {
  if roomUri.Level >= len(app.mazes) {
    return nil, nil, errors.New(fmt.Sprintf("Level %d does not exist", roomUri.Level))
  } 

  maze := app.mazes[roomUri.Level]

  location := grid.Point { X: roomUri.X, Y: roomUri.Y}

  room, ok := maze.Rooms[location]
  if !ok {
    return maze, nil, errors.New(fmt.Sprintf("Room %v not found on level %d", location, roomUri.Level))
  }

  return maze, room, nil
}

func (app application) room(c *gin.Context) {
  var params RoomUri
  if err := c.ShouldBindUri(&params); err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

  _, room, err := app.findRoom(params)
  if err != nil {
    c.String(http.StatusNotFound, err.Error())
    return
  }

  data := gin.H {
    "Room": room,
    "Level": params.Level,
  }
  
  if isHxRequest(c) {
    c.HTML(http.StatusOK, "room-partial.html.gotmpl", data)
  } else {
    c.HTML(http.StatusOK, "room.html.gotmpl", data)
  }
}

func (app application) minimap(c *gin.Context) {
  var params RoomUri
  if err := c.ShouldBindUri(&params); err != nil {
    c.String(http.StatusBadRequest, err.Error())
  }

  maze, room, err := app.findRoom(params)
  if err != nil {
    c.String(http.StatusNotFound, err.Error())
    return
  }

  c.HTML(http.StatusOK, "minimap.html.gotmpl", gin.H {
    "View": ascii.View(*maze, &room.Location),
    "Level": params.Level,
    "Room": room,
  })
}

func (app application) initRouter(router *gin.Engine) http.Handler {
  router.GET("/", app.index)
  router.GET("/mazes/:level/ascii", app.mazeAscii)
  router.GET("/mazes/:level", app.maze)
  // TODO: I'd like to have :x,:y but gin-gonic doesn't allow it... what do I do?
  router.GET("/mazes/:level/room/:x/:y", app.room)
  router.GET("/mazes/:level/room/:x/:y/minimap", app.minimap)

  router.Static("/static", "./static")
  
  return router
}

package handlers

import (
  "github.com/gin-gonic/gin"
  "xyz.haff/maze/pkg/maze"
  "xyz.haff/maze/pkg/grid"
  "xyz.haff/maze/pkg/ascii"
  "errors"
  "fmt"
  "net/http"
)

type RoomHandler struct {
  Mazes []*maze.Maze
}

type RoomUri struct {
  Level int `uri:"level"`
  X int `uri:"x"`
  Y int `uri:"y"`
}

func (handler RoomHandler) Room(c *gin.Context) {
  var params RoomUri
  if err := c.ShouldBindUri(&params); err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

  _, room, err := handler.findRoom(params)
  if err != nil {
    c.String(http.StatusNotFound, err.Error())
    return
  }

  data := gin.H {
    "Room": room,
    "Level": params.Level,
  }

  if c.GetHeader("HX-Target") == "room" {
    c.HTML(http.StatusOK, "room.html.gotmpl", data)
  } else {
    c.HTML(http.StatusOK, "page-room.html.gotmpl", data)
  }
  
}

func (handler RoomHandler) Minimap(c *gin.Context) {
  var params RoomUri
  if err := c.ShouldBindUri(&params); err != nil {
    c.String(http.StatusBadRequest, err.Error())
  }

  maze, room, err := handler.findRoom(params)
  if err != nil {
    c.String(http.StatusNotFound, err.Error())
    return
  }

  c.HTML(http.StatusOK, "maze-ascii.html.gotmpl", gin.H {
    "View": ascii.View(*maze, &room.Location),
    "Level": params.Level,
    "Room": room,
  })
}

func (handler RoomHandler) findRoom(roomUri RoomUri) (*maze.Maze, *maze.Room, error) {
  if roomUri.Level >= len(handler.Mazes) {
    return nil, nil, errors.New(fmt.Sprintf("Level %d does not exist", roomUri.Level))
  } 

  maze := handler.Mazes[roomUri.Level]

  location := grid.Point { X: roomUri.X, Y: roomUri.Y}

  room, ok := maze.Rooms[location]
  if !ok {
    return maze, nil, errors.New(fmt.Sprintf("Room %v not found on level %d", location, roomUri.Level))
  }

  return maze, room, nil
}

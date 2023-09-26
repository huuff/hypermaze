package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"xyz.haff/maze/cmd/web/util"
	"xyz.haff/maze/pkg/ascii"
	"xyz.haff/maze/pkg/grid"
	"xyz.haff/maze/pkg/maze"
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

  maze, room, err := handler.findRoom(params)
  if err != nil {
    c.String(http.StatusNotFound, err.Error())
    return
  }

  data := gin.H {
    "Room": room,
    "Level": params.Level,
    "IsEntrance": maze.Entrance.Location.Equals(room.Location),
    "IsExit": maze.Exit.Location.Equals(room.Location),
  }

  c.Header("Vary", "HX-Target")

  etagExtraData := gin.H {
    "HX-Target": c.GetHeader("HX-Target"),
  }
  if etagMatch := util.SetAndCheckEtag(c, lo.Assign[string, any](data, etagExtraData)); etagMatch {
    return
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

  data := gin.H {
    "View": ascii.View(*maze, &room.Location),
    "Level": params.Level,
    "Room": room,
  }

  if etagMatch := util.SetAndCheckEtag(c, data); etagMatch {
    return
  }

  c.HTML(http.StatusOK, "maze-ascii.html.gotmpl", data)
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

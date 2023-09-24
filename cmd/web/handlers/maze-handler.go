package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/pkg/maze"
	"xyz.haff/maze/pkg/ascii"
)

type MazeHandler struct {
  Mazes []*maze.Maze
}

func (handler MazeHandler) Ascii(c *gin.Context) {
  level, err := strconv.Atoi(c.Param("level"))
  if err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

  if level >= len(handler.Mazes) {
    c.String(http.StatusNotFound, "")
    return
  }

  maze := handler.Mazes[level]

  c.HTML(http.StatusOK, "maze-ascii-view.html.gotmpl", gin.H {
    "View": ascii.View(*maze, nil),
  })
}

func (handler MazeHandler) Maze(c *gin.Context) {
  level, err := strconv.Atoi(c.Param("level"))

  if err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

  if level >= len(handler.Mazes) {
    c.String(http.StatusNotFound, "")
    return
  }

  maze := handler.Mazes[level]

  data := gin.H {
    "Level": level,
    "View": ascii.View(*maze, nil),
    "Maze": maze,
  }

  c.HTML(http.StatusOK, "maze.html.gotmpl", data)
}

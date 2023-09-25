package handlers

import (
  "xyz.haff/maze/pkg/maze"
  "github.com/gin-gonic/gin"
  "strconv"
  "xyz.haff/maze/pkg/ascii"
  "net/http"
)

type MazeListHandler struct {
  Mazes []*maze.Maze
}

func (handler MazeListHandler) Ascii(c *gin.Context) {
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

  c.HTML(http.StatusOK, "maze-ascii.html.gotmpl", gin.H {
    "View": ascii.View(*maze, nil),
  })
}

func (handler MazeListHandler) MazeList(c *gin.Context) {
  c.HTML(http.StatusOK, "page-index.html.gotmpl", gin.H {
    "Mazes": handler.Mazes,
  })
}

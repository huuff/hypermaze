package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/cmd/web/util"
	"xyz.haff/maze/pkg/ascii"
	"xyz.haff/maze/pkg/maze"
)

type MazeHandler struct {
  Mazes []*maze.Maze
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

  util.AddDefaultCacheHeaders(c)
  if etagMatch := util.SetAndCheckEtag(c, data); etagMatch {
    return
  }

  c.HTML(http.StatusOK, "page-maze.html.gotmpl", data)
}

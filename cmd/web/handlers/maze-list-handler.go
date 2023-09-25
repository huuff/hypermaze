package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/cmd/web/util"
	"xyz.haff/maze/pkg/ascii"
	"xyz.haff/maze/pkg/maze"
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

  data := gin.H {
    "View": ascii.View(*maze, nil),
  }

  if etagMatch := util.SetAndCheckEtag(c, data); etagMatch {
    return
  }

  c.HTML(http.StatusOK, "maze-ascii.html.gotmpl", data)
}

func (handler MazeListHandler) MazeList(c *gin.Context) {
  data := gin.H {
    "Mazes": handler.Mazes,
  }
  
  if etagMatch := util.SetAndCheckEtag(c, data); etagMatch {
    return
  }

  c.HTML(http.StatusOK, "page-maze-list.html.gotmpl", data)
}

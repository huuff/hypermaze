package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/cmd/web/util"
	"xyz.haff/maze/pkg/maze"
)

type WinHandler struct {
  Mazes []*maze.Maze
}

func (handler WinHandler) Win(c *gin.Context) {
  level, err := strconv.Atoi(c.Param("level"))

  if err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

  if level >= len(handler.Mazes) {
    c.String(http.StatusNotFound, fmt.Sprintf("There is no maze %d", level))
    return
  }

  var nextLevel *int
  if level < len(handler.Mazes)-1 {
    nextLevel = new(int)
    *nextLevel = level + 1
  }

  data := gin.H {
    "Level": level,
    "NextLevel": nextLevel,
  }

  util.AddDefaultCacheHeaders(c)
  if etagMatch := util.SetAndCheckEtag(c, data); etagMatch {
    return
  }

  c.HTML(http.StatusOK, "page-win.html.gotmpl", data)
}

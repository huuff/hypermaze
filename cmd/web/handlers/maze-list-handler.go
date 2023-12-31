package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
  "strings"

	"github.com/gin-gonic/gin"
	"xyz.haff/maze/cmd/web/hal"
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

  util.AddDefaultCacheHeaders(c)
  if etagMatch := util.SetAndCheckEtag(c, data); etagMatch {
    return
  }

  c.HTML(http.StatusOK, "maze-ascii.html.gotmpl", data)
}

func (handler MazeListHandler) MazeList(c *gin.Context) {
  accepted := c.GetHeader("Accept")
  if strings.Contains(accepted, "application/json") ||
     strings.Contains(accepted, "application/hal+json") {
    handler.MazeListHAL(c)
  } else if strings.Contains(accepted, "text/html") {
    handler.MazeListHTML(c)
  } else {
    c.String(http.StatusNotAcceptable, "Can only provide application/hal+json or text/html")
  }
}

func (handler MazeListHandler) MazeListHTML(c *gin.Context) {
  data := gin.H {
    "Mazes": handler.Mazes,
  }
  
  util.AddDefaultCacheHeaders(c)
  if etagMatch := util.SetAndCheckEtag(c, data); etagMatch {
    return
  }

  c.HTML(http.StatusOK, "page-maze-list.html.gotmpl", data)
}

// TODO: I definitely should separate HTML handlers from HAL handlers
// TODO: The listed maze may be just an embedded maze?
// TODO: Links to the entrance of the maze?
type halListedMaze struct {
  Links hal.Links `json:"_links"` 
  Level int `json:"level"`
  Height int `json:"height"`
  Width int `json:"width"`
}
type halMazeListEmbedded struct {
  Mazes []halListedMaze `json:"mazes"`
}
type halMazeList struct {
  Links hal.Links `json:"_links"`
  Embedded halMazeListEmbedded `json:"_embedded"`
}
func (handler MazeListHandler) MazeListHAL(c *gin.Context) {
  response := halMazeList {
    Links: hal.Links {
      Self: hal.Link {
        Href: "/mazes",
      },
    },
    Embedded: halMazeListEmbedded {
      Mazes: make([]halListedMaze, 0, len(handler.Mazes)),
    },
  }

  for level, maze := range handler.Mazes {
    halMaze := halListedMaze {
      Links: hal.Links {
        Self: hal.Link {
          Href: fmt.Sprintf("/mazes/%d", level),
        },
      },
      Level: level,
      Height: maze.Grid.Height,
      Width: maze.Grid.Width,
    }
    response.Embedded.Mazes = append(response.Embedded.Mazes, halMaze ) 
  }

  c.Header("Content-Type", "application/hal+json")
  bytes, err := json.Marshal(response)

  if err != nil {
    c.String(http.StatusInternalServerError, err.Error())
    return
  }

  c.Writer.Write(bytes)
}

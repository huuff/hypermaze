package util

import (
  "github.com/gin-gonic/gin"
  "xyz.haff/maze/pkg/direction"
  "fmt"
)

/**
 * Whether this request was made by HTMX, which indicates an HTML fragment
 * is requested, instead of a full page
 */
func IsHxRequest(c *gin.Context) bool {
  return c.GetHeader("HX-Request") != ""
}

func DirectionToString(d direction.Direction) string {
  switch (d) {
    case direction.North:
      return "North"
    case direction.South:
      return "South"
    case direction.East:
      return "East"
    case direction.West:
      return "West"
  }
  panic(fmt.Sprintf("Weird direction %d", d))
}

func DirectionToKey(d direction.Direction) string {
  switch (d) {
    case direction.North:
      return "ArrowUp"
    case direction.South:
      return "ArrowDown"
    case direction.West:
      return "ArrowLeft"
    case direction.East:
      return "ArrowRight"
  }
  panic(fmt.Sprintf("Weird direction %d", d))
}

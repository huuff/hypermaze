package maze

import (
	"fmt"

  "xyz.haff/maze/pkg/grid"
	"github.com/samber/lo"
)

// These are the outermost points created only for showing exterior walls and displaying the exit, they can't have any other connections
func isExteriorPoint(g grid.Grid, p grid.Point) bool {
  return p.X == 0  || p.Y == 0 || p.X == g.Width*2 || p.Y == g.Height*2
}

// These are interspersed to display connections. They don't actually belong to the maze
func isConnectionPoint(p grid.Point) bool {
  return (p.X != 0 && p.X%2 == 0) || (p.Y != 0 && p.Y%2 == 0)
}

func (m Maze) AsciiView() {
  for y := range lo.Range((m.Grid.Height * 2) + 1) {
    for x := range lo.Range((m.Grid.Width * 2) + 1) {
      p := grid.Point { x, y }
      if isExteriorPoint(m.Grid, p) {
        fmt.Print("#")
      } else if isConnectionPoint(p) {
        fmt.Print("%")
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Print("\n")
  }
}

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

// Gets the points that are arround a connection point
func surroundingPoints(p grid.Point) [4]grid.Point {
  if !isConnectionPoint(p) {
    panic(fmt.Sprintf("Called `surroundingPoints` on %v, which is not a connection point", p))
  }

  result := [4]grid.Point {
    North.From(p),
    East.From(p),
    South.From(p),
    West.From(p),
  }
  return result
}

func (m Maze) AsciiView() {
  for y := range lo.Range((m.Grid.Height * 2) + 1) {
    for x := range lo.Range((m.Grid.Width * 2) + 1) {
      p := grid.Point { x, y }
      if isExteriorPoint(m.Grid, p) {
        fmt.Print("#")
      } else if isConnectionPoint(p) {
        //surroundings := surroundingPoints(p)
        //fmt.Printf("\nSurroundings of %v:\n%v\n", p, surroundings)
        fmt.Printf("%")
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Print("\n")
  }
}

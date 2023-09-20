package boundary

import (
	"xyz.haff/maze/pkg/direction"
	"xyz.haff/maze/pkg/grid"
  "github.com/samber/lo"
)

type Boundary struct {
  Location grid.Point
  Direction direction.Direction
}

func FindAll(g grid.Grid) []Boundary {
  // TODO: Exact sizes for this?
  result := make([]Boundary, 0)
  
  for x := range lo.Range(g.Width) {
    result = append(result, Boundary { grid.Point { x, 0 }, direction.North})
    result = append(result, Boundary { grid.Point { x, g.Height-1 }, direction.South})
  }

  for y := range lo.Range(g.Height) {
    result = append(result, Boundary { grid.Point { 0, y }, direction.West})
    result = append(result, Boundary { grid.Point { g.Width-1, y }, direction.East})
  }

  return result
}

package ascii

import (
  "xyz.haff/maze/pkg/grid"
)

/*
  Points in the view are "expanded" (coordinates are multiplied by 2)
  to make space in the view for purely representational elements
  (walls, etc.). These helpers encapsulate the "expanded"
  point type and allow working with it
*/

type expandedPoint grid.Point
func (p expandedPoint) asPoint() grid.Point {
  return grid.Point {
    X: p.X,
    Y: p.Y,
  }
}

func asExpanded(p grid.Point) expandedPoint {
  return expandedPoint {
    X: p.X,
    Y: p.Y,
  }
}


func unexpand(p expandedPoint) grid.Point {
  return grid.Point {
    X: (p.X-1)/2,
    Y: (p.Y-1)/2,
  }
}


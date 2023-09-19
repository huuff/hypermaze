package ascii

import (
  "xyz.haff/maze/pkg/grid"
)
// TODO: Most of this should be private right?

/*
  Points in the view are "expanded" (coordinates are multiplied by 2)
  to make space in the view for purely representational elements
  (walls, etc.). These helpers encapsulate the "expanded"
  point type and allow working with it
*/

type ExpandedPoint grid.Point
func (p ExpandedPoint) AsPoint() grid.Point {
  return grid.Point {
    X: p.X,
    Y: p.Y,
  }
}

func AsExpanded(p grid.Point) ExpandedPoint {
  return ExpandedPoint {
    X: p.X,
    Y: p.Y,
  }
}


func unexpand(p ExpandedPoint) grid.Point {
  return grid.Point {
    X: (p.X-1)/2,
    Y: (p.Y-1)/2,
  }
}


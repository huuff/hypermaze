package maze

import (
	"bytes"
	"fmt"

	"github.com/samber/lo"
	"xyz.haff/maze/pkg/grid"
	"xyz.haff/maze/pkg/direction"
)

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

func (m Maze) AsciiView() string {
  var buf bytes.Buffer

  for y := range lo.Range((m.Grid.Height * 2) + 1) {
    for x := range lo.Range((m.Grid.Width * 2) + 1) {
      p := ExpandedPoint { x, y }
      if isExterior(m.Grid, p) {
        buf.WriteString(m.exteriorView(p))
      } else if isConnection(p) {
        buf.WriteString(m.connectionView(p))
      } else {
        buf.WriteString(" ")
      }
    }

    if y < m.Grid.Height*2 {
      // Add a newline to each row before the last
      buf.WriteString("\n")
    }
  }

  return buf.String()
}

// These are the outermost points created only for showing exterior walls and displaying the exit, they can't have any other connections
func isExterior(g grid.Grid, p ExpandedPoint) bool {
  return p.X == 0  || p.Y == 0 || p.X == g.Width*2 || p.Y == g.Height*2
}

func (m Maze) exteriorView(p ExpandedPoint) string {
  if p.X%2==0 && p.Y%2==0 {
    return "#"
  }

  unexpanded := unexpand(p)

  isLeftBoundary := unexpanded.X == 0
  isRightBoundary := unexpanded.X == m.Grid.Width-1
  isTopBoundary := unexpanded.Y == 0
  isBottomBoundary := unexpanded.Y == m.Grid.Height-1

  var expectedDirection direction.Direction
  switch {
    case isLeftBoundary:
      expectedDirection = direction.West
    case isRightBoundary:
      expectedDirection = direction.East
    case isTopBoundary:
      expectedDirection = direction.North
    case isBottomBoundary:
      expectedDirection = direction.South
  }
  
  if m.Entrance.Location.Equals(unexpanded) && m.Entrance.Direction == expectedDirection {
    return "a"
  } else if m.Exit.Location.Equals(unexpanded) && m.Exit.Direction == expectedDirection {
    return "e"
  } else {
    return "#"
  }
}

// These are interspersed to display connections. They don't actually belong to the maze
func isConnection(p ExpandedPoint) bool {
  return (p.X != 0 && p.X%2 == 0) || (p.Y != 0 && p.Y%2 == 0)
}

var horizontalDirections []direction.Direction = []direction.Direction { direction.West, direction.East }
var verticalDirections []direction.Direction = []direction.Direction { direction.North, direction.South }
func (m Maze) connectionView(p ExpandedPoint) string {
  if !isConnection(p) {
    panic(fmt.Sprintf("Called `connectionPointView` on %v, which is not a connection point", p))
  }

  if p.X%2 == 0 && p.Y%2 == 0 {
    // Always just a wall
    return "#"
  } else if p.X % 2 == 0 && m.isOpenInDirections(p, horizontalDirections){
    return " "
  } else if p.Y % 2 == 0  && m.isOpenInDirections(p, verticalDirections){
    return " "
  }


  return "#"
}

func (m Maze) isOpenInDirections(p ExpandedPoint, directions []direction.Direction) bool {
  for _, direction := range directions {
    pointInDirection := unexpand(AsExpanded(direction.From(p.AsPoint())))

    roomInDirection, exists := m.Rooms[pointInDirection]
    if exists && roomInDirection.IsOpenTowards(direction.Inverse()) {
      return true
    }
  }

  return false
}

/*
  Since we expand the map (multiplying by 2 and adding 1 to sizes)
  in order to print it, this method returns the "unexpanded" point,
  which corresponds to the actual point in the map
*/
func unexpand(p ExpandedPoint) grid.Point {
  return grid.Point {
    X: (p.X-1)/2,
    Y: (p.Y-1)/2,
  }
}


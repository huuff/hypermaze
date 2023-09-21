package ascii

import (
	"bytes"
	"fmt"

	"github.com/samber/lo"
	"xyz.haff/maze/pkg/grid"
	"xyz.haff/maze/pkg/direction"
	"xyz.haff/maze/pkg/maze"
)

func View(m maze.Maze) string {
  var buf bytes.Buffer

  for y := range lo.Range((m.Grid.Height * 2) + 1) {
    for x := range lo.Range((m.Grid.Width * 2) + 1) {
      p := expandedPoint { x, y }
      if isExterior(m.Grid, p) {
        buf.WriteString(exteriorView(m, p))
      } else if isConnection(p) {
        buf.WriteString(connectionView(m, p))
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
func isExterior(g grid.Grid, p expandedPoint) bool {
  return p.X == 0  || p.Y == 0 || p.X == g.Width*2 || p.Y == g.Height*2
}

func exteriorView(m maze.Maze, p expandedPoint) string {
  if p.X%2==0 && p.Y%2==0 {
    return "#"
  }

  unexpanded := unexpand(p)

  isLeftBoundary := p.X == 0
  isRightBoundary := p.X == (m.Grid.Width*2)
  isTopBoundary := p.Y == 0
  isBottomBoundary := p.Y == (m.Grid.Height*2)

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
func isConnection(p expandedPoint) bool {
  return (p.X != 0 && p.X%2 == 0) || (p.Y != 0 && p.Y%2 == 0)
}

var horizontalDirections []direction.Direction = []direction.Direction { direction.West, direction.East }
var verticalDirections []direction.Direction = []direction.Direction { direction.North, direction.South }
func connectionView(m maze.Maze, p expandedPoint) string {
  if !isConnection(p) {
    panic(fmt.Sprintf("Called `connectionPointView` on %v, which is not a connection point", p))
  }

  if p.X%2 == 0 && p.Y%2 == 0 {
    // Always just a wall
    return "#"
  } else if p.X % 2 == 0 && isOpenInDirections(m, p, horizontalDirections){
    return " "
  } else if p.Y % 2 == 0  && isOpenInDirections(m, p, verticalDirections){
    return " "
  }


  return "#"
}

func isOpenInDirections(m maze.Maze, p expandedPoint, directions []direction.Direction) bool {
  for _, direction := range directions {
    pointInDirection := unexpand(asExpanded(direction.From(p.asPoint())))

    roomInDirection, exists := m.Rooms[pointInDirection]
    if exists && roomInDirection.IsOpenTowards(direction.Inverse()) {
      return true
    }
  }

  return false
}


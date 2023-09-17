package maze

import (
	"bytes"
	"fmt"

	"github.com/samber/lo"
	"xyz.haff/maze/pkg/grid"
	"xyz.haff/maze/pkg/direction"
)

// TODO: Print entrance and exit
func (m Maze) AsciiView() string {
  var buf bytes.Buffer

  for y := range lo.Range((m.Grid.Height * 2) + 1) {
    for x := range lo.Range((m.Grid.Width * 2) + 1) {
      p := grid.Point { x, y }
      if isExteriorPoint(m.Grid, p) {
        buf.WriteString("#")
      } else if isConnectionPoint(p) {
        if m.isOpen(p) {
          buf.WriteString(" ")
        } else {
          buf.WriteString("#")
        }
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
func isExteriorPoint(g grid.Grid, p grid.Point) bool {
  return p.X == 0  || p.Y == 0 || p.X == g.Width*2 || p.Y == g.Height*2
}

// These are interspersed to display connections. They don't actually belong to the maze
func isConnectionPoint(p grid.Point) bool {
  return (p.X != 0 && p.X%2 == 0) || (p.Y != 0 && p.Y%2 == 0)
}


func (m Maze) isOpen(p grid.Point) bool {
  if !isConnectionPoint(p) {
    panic(fmt.Sprintf("Called `isOpen` on %v, which is not a connection point", p))
  }

  if p.X%2 == 0 && p.Y%2 == 0 {
    // Always just a wall
    return false
  } else if p.X % 2 == 0 && m.isOpenInDirections(p, []direction.Direction { direction.West, direction.East }){
    return true
  } else if p.Y % 2 == 0  && m.isOpenInDirections(p, []direction.Direction { direction.North, direction.South }){
    return true
  }


  return false
}

func (m Maze) isOpenInDirections(p grid.Point, directions []direction.Direction) bool {
  for _, direction := range directions {
    pointInDirection := direction.From(p)
    pointInDirection = grid.Point { 
      X: (pointInDirection.X-1)/2,
      Y: (pointInDirection.Y-1)/2,
    }

    roomInDirection, exists := m.Rooms[pointInDirection]
    if exists && roomInDirection.IsOpenTowards(direction.Inverse()) {
      return true
    }
  }

  return false
}


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


var directions [4]Direction = [4]Direction { North, East, South, West}

func (m Maze) isOpen(p grid.Point) bool {
  if !isConnectionPoint(p) {
    panic(fmt.Sprintf("Called `isOpen` on %v, which is not a connection point", p))
  }

  if p.X%2 == 0 && p.Y%2 == 0 {
    // Always just a wall
    return false
    // TODO: Please DRY
  } else if p.X % 2 == 0 {
    for _, direction := range []Direction { East, West } {
      pointInDirection := direction.From(p)
      pointInDirection = grid.Point { 
        X: (pointInDirection.X-1)/2,
        Y: (pointInDirection.Y-1)/2,
      }

      roomInDirection, exists := m.Rooms[pointInDirection]
      if exists && roomInDirection.IsOpenTowards(direction.inverse()) {
        return true
      }
    }
  } else if p.Y % 2 == 0 {
    for _, direction := range []Direction { North, South } {
      pointInDirection := direction.From(p)
      pointInDirection = grid.Point { 
        X: (pointInDirection.X-1)/2,
        Y: (pointInDirection.Y-1)/2,
      }

      roomInDirection, exists := m.Rooms[pointInDirection]
      if exists && roomInDirection.IsOpenTowards(direction.inverse()) {
        return true
      }
    }

  }


  return false
}

func (m Maze) AsciiView() {
  for y := range lo.Range((m.Grid.Height * 2) + 1) {
    for x := range lo.Range((m.Grid.Width * 2) + 1) {
      p := grid.Point { x, y }
      if isExteriorPoint(m.Grid, p) {
        fmt.Print("#")
      } else if isConnectionPoint(p) {
        if m.isOpen(p) {
          fmt.Print(" ")
        } else {
          fmt.Print("%")
        }
      } else {
        fmt.Print(" ")
      }
    }
    fmt.Print("\n")
  }
}

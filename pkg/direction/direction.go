package direction

import (
  "fmt"
  "xyz.haff/maze/pkg/grid"
)

type Direction byte

const (
  North Direction = iota
  East
  West
  South
)

func (d Direction) Inverse() Direction {
  switch d {
    case North:
      return South
    case East:
      return West
    case South:
      return North
    case West:
      return East
  }

  panic(fmt.Sprintf("No inverse found for direction %d", d))
}

func (d Direction) From(p grid.Point) grid.Point {
  switch d {
    case North:
      return grid.Point { p.X, p.Y - 1}
    case East:
      return grid.Point { p.X + 1, p.Y }
    case South:
      return grid.Point { p.X, p.Y + 1}
    case West:
      return grid.Point { p.X - 1, p.Y }
  } 

  panic(fmt.Sprintf("No next point found for direction %d", d))
}

func Between(p1 grid.Point, p2 grid.Point) Direction {
  switch {
  case p2.X == (p1.X - 1):
    return West
  case p2.Y == (p1.Y - 1):
    return North
  case p2.X == (p1.X + 1):
    return East
  case p2.Y == (p1.Y + 1):
    return South
  }

  panic(fmt.Sprintf("No cardinal direction between %v and %v!", p1, p2))

}


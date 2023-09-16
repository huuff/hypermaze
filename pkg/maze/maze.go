package maze

import (
	"github.com/samber/lo"
	"xyz.haff/maze/pkg/grid"
)

type Direction byte

const (
  North Direction = iota
  East
  West
  South
)

type Room struct {
  connections map[Direction]*Room
}

func NewRoom() *Room {
  return &Room {
    connections: make(map[Direction]*Room),
  }
}

type Maze struct {
  rooms map[grid.Point]*Room
}

func NewMaze(g grid.Grid, edges []grid.Point) *Maze {
  rooms := make(map[grid.Point]*Room) 

  for x := range lo.Range(g.Width) {
    for y := range lo.Range(g.Height) {

    }
  }
  return &Maze { rooms: rooms }
}

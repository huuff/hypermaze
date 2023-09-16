package maze

import (
	"github.com/samber/lo"
	"xyz.haff/maze/pkg/generators"
	"xyz.haff/maze/pkg/grid"
)

type Room struct {
  Location grid.Point
  Connections map[Direction]*Room
}

func newRoom(location grid.Point) *Room {
  return &Room {
    Location: location,
    Connections: make(map[Direction]*Room),
  }
}

func (thisRoom *Room) addConnection(otherRoom *Room) {
  direction := directionBetween(thisRoom.Location, otherRoom.Location)
  thisRoom.Connections[direction] = otherRoom
  otherRoom.Connections[direction.inverse()] = thisRoom
} 

func (room Room) IsOpenTowards(d Direction) bool {
  _, ok := room.Connections[d]

  return ok
}

type Maze struct {
  Grid grid.Grid
  Rooms map[grid.Point]*Room
}

func NewMaze(g grid.Grid, edges generators.Passages) *Maze {
  rooms := make(map[grid.Point]*Room) 

  for x := range lo.Range(g.Width) {
    for y := range lo.Range(g.Height) {
      location := grid.Point { x, y }
      rooms[location] = newRoom(location)
    }
  }

  for _, edge := range edges {
    r1 := rooms[edge[0]]
    r2 := rooms[edge[1]]
    r1.addConnection(r2)
  }

  return &Maze { 
    Grid: g,
    Rooms: rooms,
  }
}

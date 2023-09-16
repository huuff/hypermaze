package maze

import (
	"testing"

	"xyz.haff/maze/pkg/grid"
)

func point(x, y int) grid.Point {
  return grid.Point { x, y }
}

func fakeRoom(x, y int) *Room {
  return &Room {
    Location: point(x, y), 
    Connections: map[Direction]*Room{},
  }
}

func TestAsciiView(t *testing.T) {
  midLeftRoom := fakeRoom(1, 0)
  midCenterRoom := fakeRoom(1, 1)
  midRightRoom := fakeRoom(1, 1)

  midLeftRoom.Connections[East] = midCenterRoom
  midCenterRoom.Connections[West] = midLeftRoom
  midCenterRoom.Connections[East] = midRightRoom
  midRightRoom.Connections[West] = midLeftRoom

  // Should give a view like
  /*
    #######
    # % % #
    #%%%%%#
    # % % #
    #%%%%%#
    # % % #
    #######
  */

  maze := Maze {
    Grid: grid.Grid { Height: 3, Width: 3 },
    Rooms: map[grid.Point]*Room {
      point(0, 0): fakeRoom(0, 0),
      point(0, 1): fakeRoom(0, 1),
      point(0, 2): fakeRoom(0, 2),
      point(1, 0): midLeftRoom,
      point(1, 1): midCenterRoom,
      point(1, 2): midRightRoom,
      point(1, 3): fakeRoom(1, 3),
      point(2, 0): fakeRoom(2, 0),
      point(2, 1): fakeRoom(2, 1),
      point(2, 2): fakeRoom(2, 2),
    },
  }

  maze.AsciiView()
}

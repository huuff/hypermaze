package maze

import (
	"strings"
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

// TODO: Add another test with some vertical passages
func TestAsciiView(t *testing.T) {
  // ARRANGE
  midLeftRoom := fakeRoom(0, 1)
  midCenterRoom := fakeRoom(1, 1)
  midRightRoom := fakeRoom(2, 1)

  midLeftRoom.Connections[East] = midCenterRoom
  midCenterRoom.Connections[West] = midLeftRoom
  midCenterRoom.Connections[East] = midRightRoom
  midRightRoom.Connections[West] = midLeftRoom

  maze := Maze {
    Grid: grid.Grid { Height: 3, Width: 3 },
    Rooms: map[grid.Point]*Room {
      point(0, 0): fakeRoom(0, 0),
      point(1, 0): fakeRoom(1, 0),
      point(2, 0): fakeRoom(2, 0),
      point(0, 1): midLeftRoom,
      point(1, 1): midCenterRoom,
      point(2, 1): midRightRoom,
      point(0, 2): fakeRoom(0, 2),
      point(1, 2): fakeRoom(1, 2),
      point(2, 2): fakeRoom(2, 2),
    },
  }

  expected := strings.TrimSpace(`
#######
# % % #
#%%%%%#
#     #
#%%%%%#
# % % #
#######
`)

  // ACT
  actual := maze.AsciiView()

  // ASSERT
  if actual != expected {
    t.Fatalf("\nExpected:\n%s\nGot:\n%s", expected, actual)
  }
}
package ascii

import (
	"strings"
	"testing"

	"xyz.haff/maze/pkg/boundary"
	"xyz.haff/maze/pkg/direction"
	"xyz.haff/maze/pkg/maze"
	"xyz.haff/maze/pkg/grid"
)

func point(x, y int) grid.Point {
  return grid.Point { x, y }
}

func fakeRoom(x, y int) *maze.Room {
  return &maze.Room {
    Location: point(x, y), 
    Connections: map[direction.Direction]*maze.Room{},
  }
}

// TODO: Add another test with some vertical passages
// TODO: Add another with entrances/exits at the corners
// TODO: Add a test with a highlight
func TestAsciiView(t *testing.T) {
  // ARRANGE
  midLeftRoom := fakeRoom(0, 1)
  midCenterRoom := fakeRoom(1, 1)
  midRightRoom := fakeRoom(2, 1)

  midLeftRoom.Connections[direction.East] = midCenterRoom
  midCenterRoom.Connections[direction.West] = midLeftRoom
  midCenterRoom.Connections[direction.East] = midRightRoom
  midRightRoom.Connections[direction.West] = midLeftRoom

  maze := maze.Maze {
    Grid: grid.Grid { Height: 3, Width: 3 },
    Rooms: map[grid.Point]*maze.Room {
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
    Entrance: boundary.Boundary {
      Location: point(0, 1),
      Direction: direction.West,
    },
    Exit: boundary.Boundary {
      Location: point(2, 1),
      Direction: direction.East,
    },
  }

  expected := strings.TrimSpace(`
#######
# # # #
#######
a     e
#######
# # # #
#######
`)

  // ACT
  actual := View(maze, nil)

  // ASSERT
  if actual != expected {
    t.Fatalf("\nExpected:\n%s\nGot:\n%s", expected, actual)
  }
}

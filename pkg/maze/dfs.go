package maze

type DfsMazeGenerator struct {
  grid Grid
  visited []Point
  unvisited []Point
  edges [][2]Point
}

func NewDfsMazeGenerator(grid Grid) DfsMazeGenerator {
  unvisited := make([]Point, grid.width * grid.height)

  for x := 0; x < grid.width; x++ {
    for y := 0; y < grid.height; y++ {
      unvisited = append(unvisited, Point { x, y})
    }
  }

  return DfsMazeGenerator {
    grid: grid,
    unvisited: unvisited,
    visited: make([]Point, grid.width * grid.height),
    edges: make([][2]Point, 0),
  }
}

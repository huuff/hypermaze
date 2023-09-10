package maze

import (
	"fmt"
	"math/rand"
	"time"
)

type DfsMazeGenerator struct {
  random *rand.Rand
  grid Grid
  visited []Point
  unvisited []Point
  edges [][2]Point
}

func NewDfsMazeGenerator(grid Grid) DfsMazeGenerator {
  randomSource := rand.NewSource(time.Now().Unix())
  random := rand.New(randomSource)

  unvisited := make([]Point, grid.Width * grid.Height)
  for x := 0; x < grid.Width; x++ {
    for y := 0; y < grid.Height; y++ {
      unvisited = append(unvisited, Point { x, y})
    }
  }

  return DfsMazeGenerator {
    random: random,
    grid: grid,
    unvisited: unvisited,
    visited: make([]Point, grid.Width * grid.Height),
    edges: make([][2]Point, 0),
  }
}

func (gen *DfsMazeGenerator) unvisitedNeighbours(p Point) []Point {
  neighbours := gen.grid.Neighbours(p)
  unvisitedNeighbours := make([]Point, 0)

  for _, neighbour := range neighbours {
    visited := Contains(gen.visited, neighbour)
    if !visited {
      unvisitedNeighbours = append(unvisitedNeighbours, neighbour)
    }
  }

  return unvisitedNeighbours
}

func (gen *DfsMazeGenerator) GenerateMaze() {
  startingPoint := Point { 0, 0 }
  gen.randomizedDfs(startingPoint)
}


func (gen *DfsMazeGenerator) randomizedDfs(p Point) {
  gen.visited = append(gen.visited, p)

  for len(gen.unvisitedNeighbours(p)) > 0 {
    // TODO: Try to not call this method twice
    neighbours := gen.unvisitedNeighbours(p)
    randomIndex := gen.random.Intn(len(neighbours))
    next := neighbours[randomIndex]
    fmt.Printf("Next: %v\n", next)

    gen.edges = append(gen.edges, [2]Point { p, next })
    gen.randomizedDfs(next)
  }
}

package generators

import (
	"fmt"
	"math/rand"
	"time"
  "slices"
  "xyz.haff/maze/pkg/grid"
)

type DfsMazeGenerator struct {
  random *rand.Rand
  grid grid.Grid
  visited []grid.Point
  unvisited []grid.Point
  edges [][2]grid.Point
}

func NewDfsMazeGenerator(g grid.Grid) DfsMazeGenerator {
  randomSource := rand.NewSource(time.Now().Unix())
  random := rand.New(randomSource)

  unvisited := make([]grid.Point, g.Width * g.Height)
  for x := 0; x < g.Width; x++ {
    for y := 0; y < g.Height; y++ {
      unvisited = append(unvisited, grid.Point { x, y})
    }
  }

  return DfsMazeGenerator {
    random: random,
    grid: g,
    unvisited: unvisited,
    visited: make([]grid.Point, g.Width * g.Height),
    edges: make([][2]grid.Point, 0),
  }
}

func (gen *DfsMazeGenerator) unvisitedNeighbours(p grid.Point) []grid.Point {
  neighbours := gen.grid.Neighbours(p)
  unvisitedNeighbours := make([]grid.Point, 0)

  for _, neighbour := range neighbours {
    visited := slices.Contains(gen.visited, neighbour)
    if !visited {
      unvisitedNeighbours = append(unvisitedNeighbours, neighbour)
    }
  }

  return unvisitedNeighbours
}

func (gen *DfsMazeGenerator) GenerateMaze() {
  startingPoint := grid.Point { 0, 0 }
  gen.randomizedDfs(startingPoint)
}


func (gen *DfsMazeGenerator) randomizedDfs(p grid.Point) {
  gen.visited = append(gen.visited, p)

  for len(gen.unvisitedNeighbours(p)) > 0 {
    // TODO: Try to not call this method twice
    neighbours := gen.unvisitedNeighbours(p)
    randomIndex := gen.random.Intn(len(neighbours))
    next := neighbours[randomIndex]
    fmt.Printf("Next: %v\n", next)

    gen.edges = append(gen.edges, [2]grid.Point { p, next })
    gen.randomizedDfs(next)
  }
}

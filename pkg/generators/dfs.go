package generators

import (
	"fmt"
	"math/rand"
	"time"
  "slices"
  "xyz.haff/maze/pkg/grid"
  "github.com/samber/lo"
)

type DfsPassageGenerator struct {
  random *rand.Rand
  grid grid.Grid
  visited []grid.Point
  unvisited []grid.Point
  passages [][2]grid.Point
}

func NewDfsPassageGenerator(g grid.Grid) DfsPassageGenerator {
  randomSource := rand.NewSource(time.Now().Unix())
  random := rand.New(randomSource)

  unvisited := make([]grid.Point, g.Width * g.Height)
  for x := range lo.Range(g.Width) {
    for y := range lo.Range(g.Height) {
      unvisited = append(unvisited, grid.Point { x, y})
    }
  }

  return DfsPassageGenerator {
    random: random,
    grid: g,
    unvisited: unvisited,
    visited: make([]grid.Point, g.Width * g.Height),
    passages: make(Passages, 0),
  }
}

func (gen *DfsPassageGenerator) unvisitedNeighbours(p grid.Point) []grid.Point {
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

func (gen *DfsPassageGenerator) GeneratePassages() Passages {
  startingPoint := grid.Point { 0, 0 }
  gen.randomizedDfs(startingPoint)
  return gen.passages
}


func (gen *DfsPassageGenerator) randomizedDfs(p grid.Point) {
  gen.visited = append(gen.visited, p)

  for len(gen.unvisitedNeighbours(p)) > 0 {
    // TODO: Try to not call this method twice
    neighbours := gen.unvisitedNeighbours(p)
    randomIndex := gen.random.Intn(len(neighbours))
    next := neighbours[randomIndex]
    fmt.Printf("Next: %v\n", next)

    gen.passages = append(gen.passages, [2]grid.Point { p, next })
    gen.randomizedDfs(next)
  }
}

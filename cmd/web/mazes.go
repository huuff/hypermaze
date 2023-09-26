package main 

import (
	"xyz.haff/maze/pkg/generators"
	"xyz.haff/maze/pkg/grid"
	"xyz.haff/maze/pkg/maze"

	"github.com/samber/lo"
)

const mazeAmount = 5

var mazeSizes []int = []int { 3, 5, 7, 10, 12, 15, 17, 20, 22, 25 }

// TODO: Maybe just one function to generate a maze and call it several times?

func generateMazes() []*maze.Maze {
  return lo.Map(mazeSizes, func(size, index int) *maze.Maze {
    grid := grid.Grid { Width: size, Height: size }
    dfs := generators.NewDfsPassageGenerator(grid)
    passages := dfs.GeneratePassages()
    return maze.NewMaze(grid, passages)
  })
}


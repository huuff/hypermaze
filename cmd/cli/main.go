package main

import (
	"fmt"

	"xyz.haff/maze/pkg/generators"
	"xyz.haff/maze/pkg/grid"
	"xyz.haff/maze/pkg/maze"
)

func main() {
  grid := grid.Grid { Width: 10, Height: 10 }
  dfs := generators.NewDfsPassageGenerator(grid)
  passages := dfs.GeneratePassages()
  maze := maze.NewMaze(grid, passages)
  asciiView := maze.AsciiView()

  fmt.Println(asciiView)
} 


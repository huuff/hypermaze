package main

import (
  "xyz.haff/maze/pkg/grid"
  "xyz.haff/maze/pkg/generators"
)

func main() {
  grid := grid.Grid { Width: 5, Height: 5 }
  dfs := generators.NewDfsPassageGenerator(grid)
  dfs.GeneratePassages()
} 


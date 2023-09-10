package main

import (
  "maze/pkg/maze"
)

func main() {
  grid := maze.Grid { Width: 5, Height: 5 }
  dfs := maze.NewDfsMazeGenerator(grid)
  dfs.GenerateMaze()
} 


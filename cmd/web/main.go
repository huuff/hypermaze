package main

import (
	"fmt"
)

func main() {
  for level, maze := range Mazes {
    fmt.Println()
    fmt.Printf("Level %d:\n", level + 1)
    fmt.Println(maze.AsciiView())
  }
}

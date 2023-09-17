package main

import (
	"xyz.haff/maze/pkg/maze"
)

type application struct {
  mazes []*maze.Maze
  templates *templateCache
}

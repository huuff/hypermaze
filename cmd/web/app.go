package main

import (
	"html/template"

	"xyz.haff/maze/pkg/maze"
)

type application struct {
  mazes []*maze.Maze
  templates *template.Template
}

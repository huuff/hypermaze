package main

import (
	"log"
	"os"

	"xyz.haff/maze/pkg/maze"
)

type application struct {
  mazes []*maze.Maze
  infoLog *log.Logger
  errorLog *log.Logger
}

func newApplication() application {
  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  return application {
    mazes: generateMazes(),
    infoLog: infoLog,
    errorLog: errorLog,
  }
}


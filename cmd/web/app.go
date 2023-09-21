package main

import (
	"log"
	"net/http"
	"os"

	"xyz.haff/maze/pkg/maze"
)

type application struct {
  mazes []*maze.Maze
  templates *templateCache
  infoLog *log.Logger
  errorLog *log.Logger
}

func newApplication() application {
  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  return application {
    mazes: generateMazes(),
    templates: newTemplateCache(),
    infoLog: infoLog,
    errorLog: errorLog,
  }
}

func (app application) serverError(w http.ResponseWriter, err error) {
  app.errorLog.Println(err)
  w.WriteHeader(http.StatusInternalServerError)
  w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

func (app application) badRequest(w http.ResponseWriter, err error) {
  app.errorLog.Println(err)
  w.WriteHeader(http.StatusBadRequest)
  w.Write([]byte(http.StatusText(http.StatusBadRequest)))
}

func (app application) notFound(w http.ResponseWriter) {
  w.WriteHeader(http.StatusNotFound)
  w.Write([]byte(http.StatusText(http.StatusNotFound)))
}

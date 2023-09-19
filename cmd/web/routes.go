package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"xyz.haff/maze/pkg/ascii"
)

func (app application) index(w http.ResponseWriter, r *http.Request) {
  err := app.templates.pages["index.html.gotmpl"].Execute(w, map[string]any {
    "Mazes": app.mazes,
  })

  if err != nil {
    // TODO: Actual log for the error
    fmt.Println(err.Error())
  }
}

func (app application) minimap(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  level, err := strconv.Atoi(vars["level"])
  if err != nil {
    // TODO: Actual log for the error
    fmt.Println(err.Error())
  }

  err = app.templates.partials.ExecuteTemplate(w, "minimap.html.gotmpl", map[string]any {
    "Minimap": ascii.View(*app.mazes[level]),
  })

  if err != nil {
    // TODO: Actual log for the error
    fmt.Println(err.Error())
  }
}

func (app application) maze(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  level, err := strconv.Atoi(vars["level"])

  if err != nil {
    // TODO: Actual log for the error
    fmt.Println(err.Error())
    return
  }

  if level >= len(app.mazes) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(http.StatusText(http.StatusNotFound)))
    return
  }

  maze := app.mazes[level]

  err = app.templates.partials.ExecuteTemplate(w, "maze-summary.html.gotmpl", map[string]any {
      "Level": level,
      "Minimap": ascii.View(*maze),
      "Maze": maze,
  })

  if err != nil {
    // TODO: Actual log for the error
    fmt.Println(err.Error())
  }
}

func (app application) routes() http.Handler {
  router := mux.NewRouter()
  router.HandleFunc("/", app.index)
  router.HandleFunc("/mazes/{level}/minimap", app.minimap)
  router.HandleFunc("/mazes/{level}", app.maze)
  router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
  
  return router
}

package main

import (
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
    app.serverError(w, err)
  }
}

func (app application) minimap(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  level, err := strconv.Atoi(vars["level"])
  if err != nil {
    app.serverError(w, err)
  }

  err = app.templates.partials.ExecuteTemplate(w, "minimap.html.gotmpl", map[string]any {
    "Minimap": ascii.View(*app.mazes[level]),
  })

  if err != nil {
    app.serverError(w, err)
    return
  }
}

// TODO: Make it able to return either a full-page or a fragment depending on the HX-Swap
func (app application) maze(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  level, err := strconv.Atoi(vars["level"])

  if err != nil {
    app.serverError(w, err)
    return
  }

  if level >= len(app.mazes) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(http.StatusText(http.StatusNotFound)))
    return
  }

  maze := app.mazes[level]

  // TODO: maze-summary doesn't seem like a good name
  err = app.templates.partials.ExecuteTemplate(w, "maze-summary.html.gotmpl", map[string]any {
      "Level": level,
      "Minimap": ascii.View(*maze),
      "Maze": maze,
  })

  if err != nil {
    app.serverError(w, err)
    return
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

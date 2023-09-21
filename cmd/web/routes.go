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
    app.serverError(w, err)
    return
  }
}

func (app application) minimap(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  level, err := strconv.Atoi(vars["level"])
  if err != nil {
    app.badRequest(w, err)
    return
  }

  err = app.templates.partials.ExecuteTemplate(w, "minimap.html.gotmpl", map[string]any {
    "Minimap": ascii.View(*app.mazes[level]),
  })

  if err != nil {
    app.serverError(w, err)
    return
  }
}

func (app application) maze(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  level, err := strconv.Atoi(vars["level"])

  if err != nil {
    app.badRequest(w, err)
    return
  }

  if level >= len(app.mazes) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(http.StatusText(http.StatusNotFound)))
    return
  }

  maze := app.mazes[level]

  data := map[string]any {
    "Level": level,
    "Minimap": ascii.View(*maze),
    "Maze": maze,
  }

  if r.Header.Get("HX-Request") == "" {
    err = app.templates.pages["maze.html.gotmpl"].Execute(w, data)
  } else {
    err = app.templates.partials.ExecuteTemplate(w, "maze-partial.html.gotmpl", data)
  }

  if err != nil {
    app.serverError(w, err)
    return
  }
}

func (app application) room(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  level, err := strconv.Atoi(vars["level"])
  
  if err != nil {
    app.badRequest(w, err)
    return
  }

  roomX, err := strconv.Atoi(vars["x"])

  if err != nil {
    app.badRequest(w, err)
  }

  roomY, err := strconv.Atoi(vars["y"])

  if err != nil {
    app.badRequest(w, err)
  }

  fmt.Printf("(%d): (%d, %d)", level, roomX, roomY)
}

func (app application) routes() http.Handler {
  router := mux.NewRouter()
  router.HandleFunc("/", app.index)
  router.HandleFunc("/mazes/{level}/minimap", app.minimap)
  router.HandleFunc("/mazes/{level}", app.maze)
  router.HandleFunc("/mazes/{level}/room/{x},{y}", app.maze)
  router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
  
  return router
}

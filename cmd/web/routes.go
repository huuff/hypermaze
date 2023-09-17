package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

  err = app.templates.pages["minimap.html.gotmpl"].Execute(w, map[string]any {
    "Level": level,
    "Minimap": app.mazes[level].AsciiView(),
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
  router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
  
  return router
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app application) index(w http.ResponseWriter, r *http.Request) {
  ts, err := template.ParseFiles("./cmd/web/templates/index.html.gotmpl")

  if err != nil {
    // TODO: Actual log for the error
    fmt.Println(err.Error())
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
    return
  }

  ts.Execute(w, map[string]any {
    "Mazes": app.mazes,
  })
}

func (app application) minimap(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  level, err := strconv.Atoi(vars["level"])
  if err != nil {
    // TODO: Actual log for the error
    fmt.Println(err.Error())
  }

  ts, err := template.ParseFiles("./cmd/web/templates/minimap.html.gotmpl")

  if err != nil {
    // TODO: Actual log for the error
    fmt.Println(err.Error())
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
    return
  }

  ts.Execute(w, map[string]any {
    "Level": level,
    "Minimap": app.mazes[level].AsciiView(),
  })
}

func (app application) routes() http.Handler {
  router := mux.NewRouter()
  router.HandleFunc("/", app.index)
  router.HandleFunc("/mazes/{level}/minimap", app.minimap)
  
  return router
}

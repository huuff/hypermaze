package main

import (
  "net/http"
  "html/template"
  "fmt"
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

func (app application) routes() http.Handler {
  mux := http.NewServeMux()
  mux.HandleFunc("/", app.index)
  
  return mux
}

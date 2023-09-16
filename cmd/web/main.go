package main

import (
	"fmt"

  "net/http"
  "html/template"
  "xyz.haff/maze/pkg/maze"
)

func index(w http.ResponseWriter, r *http.Request) {
  ts, _ := template.ParseFiles("./cmd/web/templates/index.html.gotmpl")

  templateData := struct { Mazes []*maze.Maze }{ Mazes: Mazes }
  ts.Execute(w, templateData)
}

func main() {
  for level, maze := range Mazes {
    fmt.Println()
    fmt.Printf("Level %d:\n", level + 1)
    fmt.Println(maze.AsciiView())
  }

  mux := http.NewServeMux()
  mux.HandleFunc("/", index)

  server := http.Server {
    Addr: ":8080",
    Handler: mux,
  }

  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }
}

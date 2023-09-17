package main

import (
  "net/http"
)

func main() {
  app := application {
    mazes: generateMazes(),
  }

  server := http.Server {
    Addr: ":8080",
    Handler: app.routes(),
  }

  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }
}

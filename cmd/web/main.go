package main

import (
	"fmt"
	"net/http"
)

func main() {
  app := newApplication()

  port := "8080"

  server := http.Server {
    Addr: ":" + port,
    Handler: app.routes(),
  }

  fmt.Printf("Starting server on port %s\n", port)
  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }
}

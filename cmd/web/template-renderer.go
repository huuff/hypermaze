package main

import (
	"fmt"
	"path/filepath"
  "strings"

	"github.com/gin-contrib/multitemplate"
)

func newRenderer() multitemplate.Renderer {
  r := multitemplate.NewRenderer()
  
  pages, err := filepath.Glob("templates/pages/*")
  if err != nil {
    panic(err)
  }
  for _, pagePath := range pages {
    name := fmt.Sprintf("pages/%s", baseName(pagePath))
    fmt.Printf("Loading page template: %s\n", name)
    r.AddFromFiles(name, "templates/base.html.gotmpl", pagePath)
  }

  partials, err := filepath.Glob("templates/partials/*")
  if err != nil {
    panic(err)
  }

  for _, partialPath := range partials {
    name := fmt.Sprintf("partials/%s", baseName(partialPath))
    fmt.Printf("Loading partial template: %s\n", name)
    r.AddFromFiles(name, partialPath)
  }


  return r
}


func baseName(path string) string {
  pathParts := strings.Split(path, "/")
  fileParts := strings.Split(pathParts[2], ".")
  return fileParts[0]
}

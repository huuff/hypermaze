package main

import (
	"html/template"
	"path/filepath"
)

type templateCache struct {
  pages map[string]*template.Template
  partials *template.Template
}

func newTemplateCache() *templateCache {
  pagesCache := map[string]*template.Template{}

  pages, err := filepath.Glob("./cmd/web/templates/pages/*.gotmpl")

  if err != nil {
    panic(err)
  }

  for _, page := range pages {
    name := filepath.Base(page)

    ts, err := template.New(name).ParseGlob("./cmd/web/templates/partials/*.gotmpl")

    if err != nil {
      panic(err)
    }

    ts, err = ts.ParseFiles(page)

    if err != nil {
      panic(err)
    }

    pagesCache[name] = ts
  }  

  partials, err := template.ParseGlob("./cmd/web/templates/partials/*.gotmpl")

  if err != nil {
    panic(err)
  }

  return &templateCache {
    pages: pagesCache,
    partials: partials,
  }
}

package main

import (
  "html/template"
)

func newTemplateCache()*template.Template {
  ts, err := template.ParseGlob("./cmd/web/templates/*.gotmpl")
  if err != nil {
    panic(err)
  }

  return ts
}

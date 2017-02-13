package main

import (
  "fmt"
  "net/http"
  "os"

  "github.com/attic-labs/noms/go/spec"
)

func handler(w http.ResponseWriter, r *http.Request) {
  sp, err := spec.ForDataset("http://localhost:8000::todo")
  if err != nil {
    fmt.Fprintf(os.Stderr, "Could not create dataset: %s\n", err)
    fmt.Fprintf(w, "Coud not create dataset: %s\n", err)
    return
  }
  defer sp.Close()

  if _, ok := sp.GetDataset().MaybeHeadValue(); !ok {
    fmt.Fprintf(w, "head is empty\n")
  }
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}

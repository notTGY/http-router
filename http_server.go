package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World!")
}

func main() {
  http.HandleFunc("/", handler)
  fmt.Println("Starting server on :80")
  if err := http.ListenAndServe(":80", nil); err != nil {
      fmt.Println(err)
  }
}

package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World!")
}

func main() {
  fmt.Println("Hello, World!")
  http.HandleFunc("/", handler)
  fmt.Println("Starting server on :8000")
  if err := http.ListenAndServe(":8000", nil); err != nil {
      fmt.Println(err)
  }
}

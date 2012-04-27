package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "os"
  "time"
)

type Message struct {
  Hostname string
  Cwd string
  PageSize int
  Time string
}

func setMessage() []byte {
  name, _ := os.Hostname()
  cwd, _ := os.Getwd()
  page := os.Getpagesize()
  t := time.Now().Format("20060102150405")
  m := Message{name, cwd, page, t}
  b, _ := json.Marshal(m)
  return b
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s", setMessage())
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}

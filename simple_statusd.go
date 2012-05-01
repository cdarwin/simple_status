package main

import (
  "encoding/json"
  "fmt"
  "io"
  "net/http"
  "os"
  "time"
)

type Message struct {
  Host string
  Load string
  Time string
}

func host() string {
  host, err := os.Hostname()
  if err != nil {
    return fmt.Sprint(err)
  }
  return host
}

func load() string {
  f, err := os.Open("/proc/loadavg")
  if err != nil {
    return fmt.Sprint(err)
  }
  var r io.Reader
  r = f
  var a, b, c, d, e string
  fmt.Fscanf(r, "%s %s %s %s %s", &a, &b, &c, &d, &e)
  return fmt.Sprintf("%s %s %s %s %s", a, b, c, d, e)
}

func now() string {
  return time.Now().Format("2006 01/02 1504-05")
}

func message() []byte {
  m := Message{host(), load(), now()}
  b, err := json.Marshal(m)
  if err != nil {
    return nil
  }
  return b
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s", message())
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}

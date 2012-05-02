package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
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
  b, err := ioutil.ReadFile("/proc/loadavg")
  if err != nil {
    return fmt.Sprint(err)
  }
  return fmt.Sprintf("%s", b[:len(b)-1])
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

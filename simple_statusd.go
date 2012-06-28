package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "os/exec"
  "time"
)

type Message struct {
  Host string
  Load string
  Rams string
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

func ram() string {
  free := exec.Command("free", "-m")
  grep := exec.Command("awk", `/cache:/{print$3}`)
  fout, err := free.StdoutPipe()
  if err != nil {
    return fmt.Sprint(err)
  }
  free.Start()
  grep.Stdin = fout
  gout, err := grep.Output()
  if err != nil {
    return fmt.Sprint(err)
  }
  return fmt.Sprintf("%sMB", gout)
}

func now() string {
  return time.Now().Format("2006 01/02 1504-05")
}

func message() []byte {
  m := Message{host(), load(), ram(), now()}
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

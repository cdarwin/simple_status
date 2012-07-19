package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	port  = flag.String("p", ":8080", "http service address")
	token = flag.String("t", "", "http auth token")
	tls   bool
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
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return fmt.Sprint(err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)
	b := make([]byte, 100)
	var free, total string
	for line, isPrefix, err := bufReader.ReadLine(); err != io.EOF; line, isPrefix, err = bufReader.ReadLine() {
		b = append(b, line...)

		if !isPrefix {
			switch {
			case bytes.Contains(b, []byte("MemFree")):
				s := bytes.Fields(b)
				free = string(s[1])
			case bytes.Contains(b, []byte("MemTotal")):
				s := bytes.Fields(b)
				total = string(s[1])
			}
			b = b[:0]
		}
	}
	return fmt.Sprintf("%s/%s", free, total)
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
	flag.Parse()
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

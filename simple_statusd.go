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
	return string(b[:len(b)-1])
}

func ram() string {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return fmt.Sprint(err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)
	b := make([]byte, 0, 100)
	var free, total string
	for line, isPrefix, err := bufReader.ReadLine(); err != io.EOF; line, isPrefix, err = bufReader.ReadLine() {
		if err != nil {
			return fmt.Sprint(err)
		}
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
	w.Write(message())
}

func main() {
	port  := flag.String("p", ":8080", "http service address")
	token := flag.String("t", "", "http auth token")
	tls := flag.Bool("ssl", false, "TLS boolean flag")
	flag.Parse()

	url := "/"
	if *token != "" {
		url += *token
	}
	http.HandleFunc(url, handler)
	switch *tls {
	case false:
		err := http.ListenAndServe(*port, nil)
		if err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	case true:
		err := http.ListenAndServeTLS(*port, "cert.pem", "key.pem", nil)
		if err != nil {
			log.Fatal("ListenAndServeTLS:", err)
		}
	}
}

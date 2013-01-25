package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Ram struct {
	Free  string `json:"free"`
	Total string `json:"total"`
}

func ramHandler(w http.ResponseWriter, r *http.Request) {
	R := ram()
	m := Ram{R.Free, R.Total}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(auth(b, r.FormValue("token")))
}

func ram() Ram {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Fatal("os.Open: ", err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)
	b := make([]byte, 0, 100)
	var r Ram
	for line, isPrefix, err := bufReader.ReadLine(); err != io.EOF; line, isPrefix, err = bufReader.ReadLine() {
		if err != nil {
			log.Fatal("bufReader.ReadLine: ", err)
		}
		b = append(b, line...)

		if !isPrefix {
			switch {
			case bytes.Contains(b, []byte("MemFree")):
				s := bytes.Fields(b)
				r.Free = string(s[1])
			case bytes.Contains(b, []byte("MemTotal")):
				s := bytes.Fields(b)
				r.Total = string(s[1])
			}
			b = b[:0]
		}
	}
	return r
}

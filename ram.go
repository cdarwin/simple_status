package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

type Ram struct {
	Free  string `json:"free"`
	Total string `json:"total"`
}

func ram() interface{} {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return "Unsupported"
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

package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

type Ram struct {
	Free  int `json:"free"`
	Total int `json:"total"`
}

func ram() interface{} {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return "Unsupported"
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)
	b := make([]byte, 0, 100)
	var free, total int

	for line, isPrefix, err := bufReader.ReadLine(); err != io.EOF; line, isPrefix, err = bufReader.ReadLine() {
		if err != nil {
			log.Fatal("bufReader.ReadLine: ", err)
		}
		b = append(b, line...)

		if !isPrefix {
			switch {
			case bytes.Contains(b, []byte("MemFree")):
				free = toInt(bytes.Fields(b)[1])
			case bytes.Contains(b, []byte("MemTotal")):
				total = toInt(bytes.Fields(b)[1])
			}
			b = b[:0]
		}
	}
	return Ram{free, total}
}

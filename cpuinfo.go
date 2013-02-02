package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
)

type CpuInfo struct {
	Processors int `json:"processors"`
	Siblings   int `json:"siblings"`
	Cores      int `json:"cores"`
}

func cpuinfo() interface{} {
	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "Unsupported"
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)
	b := make([]byte, 0, 100)
	var procs, sibs, cores int

	for line, isPrefix, err := bufReader.ReadLine(); err != io.EOF; line, isPrefix, err = bufReader.ReadLine() {
		if err != nil {
			log.Fatal("bufReader.ReadLine: ", err)
		}
		b = append(b, line...)

		if !isPrefix {
			switch {
			case bytes.Contains(b, []byte("processor")):
				procs, _ = strconv.Atoi(string(bytes.Fields(b)[2]))
			case bytes.Contains(b, []byte("siblings")):
				sibs, _ = strconv.Atoi(string(bytes.Fields(b)[2]))
			case bytes.Contains(b, []byte("cores")):
				cores, _ = strconv.Atoi(string(bytes.Fields(b)[3]))
			}
			b = b[:0]
		}
	}
	return CpuInfo{procs + 1, sibs, cores}
}

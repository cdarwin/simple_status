package main

import (
	"io/ioutil"
	"log"
)

type Load struct {
	Avg1 string `json:"avg1"`
	Avg2 string `json:"avg2"`
	Avg3 string `json:"avg3"`
}

func load() Load {
	b, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		log.Fatal(err)
	}
	var l Load
	l.Avg1 = string(b[0:4])
	l.Avg2 = string(b[5:9])
	l.Avg3 = string(b[10:14])
	return l
}

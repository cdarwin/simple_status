package main

import (
	"io/ioutil"
)

type Load struct {
	Avg1 string `json:"avg1"`
	Avg2 string `json:"avg2"`
	Avg3 string `json:"avg3"`
}

func load() interface{} {
	b, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return "Unsupported"
	}
	return Load{string(b[0:4]), string(b[5:9]), string(b[10:14])}
}

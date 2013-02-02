package main

import (
	"encoding/json"
)

type System struct {
	Host string      `json:"host"`
	Disk interface{} `json:"disk"`
	Cpu  interface{} `json:"cpuinfo"`
	Load interface{} `json:"load"`
	Ram  interface{} `json:"ram"`
	Time string      `json:"time"`
}

func system(s string) []byte {
	m := System{host(), disk(s), cpuinfo(), load(), ram(), now()}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil
	}
	return b
}

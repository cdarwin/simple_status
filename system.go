package main

import (
	"encoding/json"
)

type System struct {
	Host string      `json:"host"`
	Disk interface{} `json:"disk"`
	Load interface{} `json:"load"`
	Ram  interface{} `json:"ram"`
	Time string      `json:"time"`
}

func system(s string) []byte {
	m := System{host(), disk(s), load(), ram(), now()}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil
	}
	return b
}

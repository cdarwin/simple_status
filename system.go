package main

import (
	"encoding/json"
	"net/http"
)

type System struct {
	Host string `json:"host"`
	Load Load   `json:"load"`
	Ram  Ram    `json:"ram"`
	Time string `json:"time"`
}

func systemHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(auth(system(), r.FormValue("token")))
}

func system() []byte {
	m := System{host(), load(), ram(), now()}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil
	}
	return b
}

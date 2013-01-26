package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
)

type Shell struct {
	Output string `json:"out"`
}

func shellHandler(w http.ResponseWriter, r *http.Request) {
	var m interface{}
	if *token != "" {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		S := shell(r.FormValue("exec"))
		m = Shell{S.Output}
	} else {
		m = "You must set an auth token to use the shell endpoint"
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}
	w.Write(auth(b, r.FormValue("token")))
}

func shell(s string) Shell {
	out, err := exec.Command("/bin/sh", "-c", s).Output()
	if err != nil {
		log.Println(err)
	}
	var sh Shell
	sh.Output = string(out)
	return sh
}

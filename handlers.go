package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s string
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		if r.FormValue("token") != *token {
			w.Write([]byte("Unauthorized"))
			return
		}
		switch r.FormValue("disk") {
		case "":
			s = "/"
		default:
			s = r.FormValue("disk")
		}
		fn(w, r, s)
	}
}

func doMarshall(m interface{}) []byte {
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println(err)
	}
	return b
}

func diskHandler(w http.ResponseWriter, r *http.Request, s string) {
	w.Write(doMarshall(disk(s)))
}

func hostHandler(w http.ResponseWriter, r *http.Request, s string) {
	w.Write(doMarshall(host()))
}

func loadHandler(w http.ResponseWriter, r *http.Request, s string) {
	w.Write(doMarshall(load()))
}

func ramHandler(w http.ResponseWriter, r *http.Request, s string) {
	w.Write(doMarshall(ram()))
}

func cpuHandler(w http.ResponseWriter, r *http.Request, s string) {
	w.Write(doMarshall(cpuinfo()))
}

func systemHandler(w http.ResponseWriter, r *http.Request, s string) {
	w.Write(system(s))
}

func shellHandler(w http.ResponseWriter, r *http.Request, s string) {
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
	w.Write(doMarshall(m))
}

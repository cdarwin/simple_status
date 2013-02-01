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

func hostHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.MarshalIndent(host(), "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(auth(b, r.FormValue("token")))
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	L := load()
	m := Load{L.Avg1, L.Avg2, L.Avg3}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(auth(b, r.FormValue("token")))
}

func ramHandler(w http.ResponseWriter, r *http.Request) {
	R := ram()
	m := Ram{R.Free, R.Total}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(auth(b, r.FormValue("token")))
}

func systemHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(auth(system(), r.FormValue("token")))
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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

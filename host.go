package main

import (
	"fmt"
	"os"
)

func hostHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.MarshalIndent(host(), "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(auth(b, r.FormValue("token")))
}

func host() string {
	host, err := os.Hostname()
	if err != nil {
		return fmt.Sprint(err)
	}
	return host
}

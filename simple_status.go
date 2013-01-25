package main

import (
	"flag"
	"log"
	"net/http"
)

var token *string

func auth(m []byte, t string) []byte {
	if t != *token {
		return []byte("Unauthorized")
	}
	return m
}

func main() {
	port := flag.String("p", ":8080", "http service address")
	token = flag.String("t", "", "http auth token")
	tls := flag.Bool("ssl", false, "TLS boolean flag")
	flag.Parse()

	base := "/1/api"
	http.HandleFunc(base+"/system", systemHandler)
	http.HandleFunc(base+"/system/ram", ramHandler)
	http.HandleFunc(base+"/system/load", loadHandler)
	http.HandleFunc(base+"/shell", shellHandler)

	switch *tls {
	case false:
		err := http.ListenAndServe(*port, nil)
		if err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	case true:
		err := http.ListenAndServeTLS(*port, "cert.pem", "key.pem", nil)
		if err != nil {
			log.Fatal("ListenAndServeTLS:", err)
		}
	}
}

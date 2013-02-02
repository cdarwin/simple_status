package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var token *string

func toInt(b []byte) (i int) {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		log.Println("Failed to convert string to int")
	}
	return
}

func toFloat(b []byte) (f float64) {
	f, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		log.Println("Failed to convert string to float")
	}
	return
}

func main() {
	port := flag.String("p", ":8080", "http service address")
	token = flag.String("t", "", "http auth token")
	tls := flag.Bool("ssl", false, "TLS boolean flag")
	flag.Parse()

	base := "/1/api"
	http.HandleFunc(base+"/system", makeHandler(systemHandler))
	http.HandleFunc(base+"/system/ram", makeHandler(ramHandler))
	http.HandleFunc(base+"/system/load", makeHandler(loadHandler))
	http.HandleFunc(base+"/system/host", makeHandler(hostHandler))
	http.HandleFunc(base+"/system/disk", makeHandler(diskHandler))
	http.HandleFunc(base+"/system/cpuinfo", makeHandler(cpuHandler))
	http.HandleFunc(base+"/shell", makeHandler(shellHandler))

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

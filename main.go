package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := 8080
	if v := os.Getenv("PORT"); v != "" {
		var err error
		port, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("env PORT must be int. %s", v)
		}
	}
	flag.IntVar(&port, "port", port, "application listen port")
	flag.Parse()
	log.Printf("start :%d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), http.NotFoundHandler()))
}

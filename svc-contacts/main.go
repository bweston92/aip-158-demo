package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()

	s, err := getStoreFromEnvironment()
	if err != nil {
		log.Fatalf("unable to construct store: %s", err)
	}

	a := newApp(s)
	startAdminConsole(a)

	log.Printf("--\n")
	log.Printf("Add data to this application using: http://localhost:3100/")
	log.Printf("--\n")

	c := make(chan struct{})
	<-c
}

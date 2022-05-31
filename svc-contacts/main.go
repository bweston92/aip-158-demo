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

	if err := startAdminConsole(s); err != nil {
		log.Fatalf("unable to start admin console: %s", err)
	}

	if err := startPublicAPI(s); err != nil {
		log.Fatalf("unable to start admin console: %s", err)
	}

	log.Printf("--\n")
	log.Printf("Add data to this application using: http://localhost:3100/")
	log.Printf("--\n")

	c := make(chan struct{})
	<-c
}

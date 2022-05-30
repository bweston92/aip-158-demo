package main

import "log"

func main() {
	a := newApp(nil)
	startAdminConsole(a)

	log.Printf("--\n")
	log.Printf("Add data to this application using: http://localhost:3100/")
	log.Printf("--\n")

	c := make(chan struct{})
	<-c
}

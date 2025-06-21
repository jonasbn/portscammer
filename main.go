package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const port string = ":8080"

func loop() {
	ln, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Error starting server on port: %s - %s\n", port, err)
	}
	log.Printf("Starting server on port: %s\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection on on port: %s - %s\n", port, err)
		}
		// REF: https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
		fmt.Print("\033[H\033[2J\007")
		log.Printf("Connection received from %s on port %s\n", conn.RemoteAddr().String(), port)

	}
}

// REF: https://stackoverflow.com/questions/48653502/interrupting-an-infinite-loop-in-golang-with-signals
func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Printf("Program terminated")
		os.Exit(1)
	}()
	loop()
}

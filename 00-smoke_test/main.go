package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	port := ":5001"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("could not start listener", err)
	}

	defer listener.Close()

	fmt.Printf("Server started on 0.0.0.0%s\n", port)
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("Error on request", err)
		}

		go handleRequest(connection)
	}
}

func handleRequest(connection net.Conn) {
	for {
		buffer := make([]byte, 1024)
		length, err := connection.Read(buffer)
		if err != nil {
			log.Println("Could not read request", err)
			connection.Close()
			return
		}

		fmt.Printf("Got %s", string(buffer))

		_, err = connection.Write(buffer[:length])
		if err != nil {
			log.Println("Could not write response", err)
			connection.Close()
			return
		}
	}
}
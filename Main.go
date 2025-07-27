package main

import (
    "fmt",
	"net",
	"io",
	"os"
)

func main() {

    fmt.Println("Listening on port :6379")

	// Create a new server (AKA listener)
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close() // Close connection once finished

	/*
		Create an infinite loop to receive requests from our client redis-cli
	*/
	for {
		buf := make([]byte, 1024)

		// Read message from client
		_, err = conn.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client", err.Error())
			os.Exit(1)
		}

		// ignore request and respond with OK regardless of the client message
		conn.Write([]byte("+OK]\r\n"))
	}




}


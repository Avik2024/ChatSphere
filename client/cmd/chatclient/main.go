package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "server:8080") // server is the name of the Docker service for the server
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Read from user input and send to server
	fmt.Println("Connected to the server. Type messages to send:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, err := fmt.Fprintf(conn, scanner.Text()+"\n")
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}


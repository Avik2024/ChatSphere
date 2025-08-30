package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Listen for incoming connections
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
		os.Exit(1)
	}
	defer listen.Close()
	fmt.Println("Server started on port 8080...")

	// Handle client connections
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleClient(conn) // Handle each client in a new goroutine
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	// Read messages from the client
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from client:", err)
			break
		}
		// Print received message from the client
		fmt.Printf("Message from client: %s\n", string(buffer[:n]))
	}
}

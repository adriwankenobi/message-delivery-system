package client

import (
	"fmt"
    "net"
)

/**
	CONNECT TO SERVER
*/
func Connect(protocol, host, port string) (net.Conn, error) {
    
    fmt.Println("Connecting to hub server...")
    
    // Connect to the socket
	return net.Dial(protocol, host + ":" + port)
}

/**
	SENDS REQUEST AND GETS THE RESPONSE
*/
func Send(conn net.Conn, text string) (string, error) {
	
	// Send request
	fmt.Fprintf(conn, text)
	
	// Get response
	// Make a buffer (1024KB) to hold incoming data
	buf := make([]byte, 1024000)
	
	// Read the incoming connection into the buffer
	bytesRead, err := conn.Read(buf)
	
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return "", err
	}
	
	// Build the message
	message := string(buf[:bytesRead])
	
	fmt.Println("Received message from hub: ", message)
	
	return message, nil
}
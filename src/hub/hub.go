package hub

import (
    "fmt"
    "net"
)

/**
	RUNS THE HUB
*/
func Run(protocol, host, port string) error {

    fmt.Println("Starting hub server...")
    
	// Listen for incoming connections
	listener, err := net.Listen(protocol, host + ":" + port)
	
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		return err
	}
    
	// Close the listener when the application stops
	defer listener.Close()
	
	fmt.Println("Listening on", CONN_HOST + ":" + CONN_PORT)
    
	for {
		// Accept an incoming connection
		conn, err := listener.Accept()
		
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return err
		}
		
        fmt.Println("Received message from " + conn.RemoteAddr().String() + " to " + conn.LocalAddr().String())
        
		// Handle connections in a new goroutine
		go handleRequest(conn)
	}
}

/**
	HANDLES REQUEST
*/
func handleRequest(conn net.Conn) error {

	// Make a buffer (1024KB) to hold incoming data
	buf := make([]byte, 1024000)
	
	// Read the incoming connection into the buffer
	bytesRead, err := conn.Read(buf)
	
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return err
	}
	
	// Build the message
	message := string(buf[:bytesRead])
	
	fmt.Println("Received message from client: ", message)
	
	// Build the response
	response := "echo " + message
	
	// Send the response
	conn.Write([]byte(response))
	
	// Close the connection
	conn.Close()
	
	return nil
}


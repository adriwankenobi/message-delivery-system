package client

import (
	"fmt"
    "net"
)

/** 
	MessageClient interface defines what methods the message client should have
*/
type MessageClient interface {
	
	// Connect to the hub listening on the specified address
	Connect(laddr string) error
	
	// Stop stops the server
	SendAndGet(request string) (string, error)
}

/** 
	Custom TCP MessageClient implementation
*/
const CONN_TYPE = "tcp"

type TcpMessageClient struct {

	// address
	laddr string
	
	// connection
	conn net.Conn
}

/**
	CONNECT TO SERVER
*/
func (t *TcpMessageClient) Connect(laddr string) error {
    
    fmt.Println("Connecting to hub server...")
    
    // Connect to the socket
	conn, err := net.Dial(CONN_TYPE, laddr)
	
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return err
	}
	
	t.laddr = laddr
	t.conn = conn
	
	return nil
}

/**
	SENDS REQUEST AND GETS THE RESPONSE
*/
func (t *TcpMessageClient) SendAndGet(request string) (string, error) {
	
	// Send request
	fmt.Fprintf(t.conn, request)
	
	// Get response
	// Make a buffer (1024KB) to hold incoming data
	buf := make([]byte, 1024000)
	
	// Read the incoming connection into the buffer
	bytesRead, err := t.conn.Read(buf)
	
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return "", err
	}
	
	// Build the message
	message := string(buf[:bytesRead])
	
	fmt.Println("Received message from hub:", message)
	
	return message, nil
}
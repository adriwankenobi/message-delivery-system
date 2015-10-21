package hub

import (
    "fmt"
    "net"
    "message-delivery-system/src/utils"
)

/** 
	MessageHub interface defines what methods the message hub should have
*/
type MessageHub interface {
	
	// Start the hub and binds it to the specified address
	Start(laddr string) error
	
	// ClientIDs Returns client ids for the currently connected clients
	//ClientIDs() []uint64
	
	// Stop stops the server
	Stop() error
}

/** 
	Custom TCP MessageHub implementation
*/
const CONN_TYPE = "tcp"

type TcpMessageHub struct {

	// address
	laddr string
	
	// listener
	ln net.Listener
	
	// map of connections
	conns map[uint64]net.Conn
	
	// quit channel
	quit chan bool
}

/**
	RUNS THE HUB
*/
func (t *TcpMessageHub) Start(laddr string) error {

    fmt.Println("Starting hub server...")
    
	// Listen for incoming connections
	ln, err := net.Listen(CONN_TYPE, laddr)
	
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return err
	}
    
    t.laddr = laddr
	t.ln = ln
	t.conns = make(map[uint64]net.Conn)
	t.quit = make(chan bool)
	
	// Schedule gracefull shutdown
	defer t.Stop()
	
	fmt.Println("Listening on", laddr)
    
	for {
		// Accept an incoming connection
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			
			// Continuing accepting connections unless quit signal is sent through channel
			select {
			case <- t.quit:
				return nil
			default:
			}
			
			continue
		}
        
        // Generate unique id and add connection to the pool
        id := utils.NextId()
		t.conns[id] = conn
		
		// Handle connections in a new goroutine
		go t.handleRequest(id)
	}
}

func (t *TcpMessageHub) Stop() error {

	fmt.Println("Shutting down hub server...")
	
	// Send quit signal
	close(t.quit)
	
	for id, conn := range t.conns {
		if conn != nil {
			fmt.Println("Closing connection", id)
			conn.Close()
			delete(t.conns, id)
		}
	}
	
	fmt.Println("Closing listener")
	return t.ln.Close()
}

/**
	HANDLES REQUEST
*/
func (t *TcpMessageHub) handleRequest(id uint64) error {
	
	for {
	
		// Make a buffer (1024KB) to hold incoming data
		buf := make([]byte, 1024000)
		
		// Read the incoming connection into the buffer
		bytesRead, err := t.conns[id].Read(buf)
		
		if err != nil {
			fmt.Println("Hub error reading:", err.Error())
			return err
		}
		
		// Build the message
		message := string(buf[:bytesRead])
		
		fmt.Println("Received message from client:", message)
		
		// Build the response
		response := "echo " + message
		
		// Send the response
		t.conns[id].Write([]byte(response))
	}
	
	return nil
}
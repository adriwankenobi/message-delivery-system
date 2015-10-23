package hub

import (
    "fmt"
    "net"
    "errors"
    "io"
    "message-delivery-system/src/utils"
    "message-delivery-system/src/messages"
    "message-delivery-system/src/services"
)

/** 
	MessageHub interface defines what methods the message hub should have
*/
type MessageHub interface {
	
	// Start the hub and binds it to the specified address
	Start(laddr string) error
	
	// ClientIDs Returns client ids for the currently connected clients
	ClientIDs() []uint64
	
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
	
	// connected
	running bool
}

/**
	RUNS THE HUB
*/
func (t *TcpMessageHub) Start(laddr string) error {

	if t.isRunning() {
	
		fmt.Println("Hub server is already running")
		return nil
		
	} else {
	
	    fmt.Println("Starting hub server...")
	    
		// Listen for incoming connections
		ln, err := net.Listen(CONN_TYPE, laddr)
		
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			return err
		}
	    
		// Schedule gracefull shutdown
		defer t.Stop()
		
		fmt.Println("Listening on", laddr)
		
		t.laddr = laddr
		t.ln = ln
		t.conns = make(map[uint64]net.Conn)
		t.quit = make(chan bool)
		t.running = true
	    
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
}

/**
	RETURNS CONNECTED CLIENTS
*/
func (t *TcpMessageHub) ClientIDs() []uint64 {
	var list []uint64
	for id, _ := range t.conns {
		list = append(list, id)
	}
	return list
}

/**
	STOPS THE HUB
*/
func (t *TcpMessageHub) Stop() error {

	if !t.isRunning() {
	
		fmt.Println("Hub server is already stopped")
		return nil
		
	} else {
	
		fmt.Println("Shutting down hub server...")
		
		// Send quit signal
		close(t.quit)
		
		for id, conn := range t.conns {
			if conn != nil {
				fmt.Println("Closing connection", id)
				err := conn.Close()
				if err != nil {
					return err
				}
				delete(t.conns, id)
			}
		}
		
		fmt.Println("Closing listener")
		
		err := t.ln.Close()
		if err != nil {
			return err
		} else {
			t.running = false
			return nil
		}
	}
}

/**
	IS RUNNING
*/
func (t *TcpMessageHub) isRunning() bool {
	return t.running
}

/**
	HANDLES REQUEST
*/
func (t *TcpMessageHub) handleRequest(id uint64) error {
	
	for {
		
		// Read request
		request, err := utils.Read(t.conns[id])

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading request:", err.Error())
			return err
		}
		
		fmt.Println("Request from client:", string(request))
		
		// Decode request
		decoded, err := messages.Decode(request)
		
		if err != nil {
			fmt.Println("Error decoding request:", err.Error())
			return err
		}
		
		// Handle request and build response
		var response messages.Response
		switch decoded.GetMessageType() {
	    case messages.EchoRequestMessage:
	        response = services.HandleEcho(decoded.(messages.EchoRequest), id)
	    case messages.IdRequestMessage:
	        response = services.HandleId(decoded.(messages.IdRequest), id)
	    case messages.ListRequestMessage:
	        response = services.HandleList(decoded.(messages.ListRequest), id, t.ClientIDs())
	    case messages.RelayRequestMessage:
	        response = services.HandleRelay(decoded.(messages.RelayRequest))
	    default:
	    	err := errors.New("Error handling: Cannot find handler for this message type")
	    	return err
	    }
		
		t.handleResponse(response)

	}
	
	return nil
}

func (t *TcpMessageHub) handleResponse(response messages.Response) error {

	// Encode response
	encoded := messages.Encode(response)
	
	// Send the responses
	for _, id := range response.GetReceiverIds() {

		err := utils.Write(t.conns[id], encoded)
		
		if err != nil { 
			fmt.Println("Error sending response:", err.Error())
			return err
		}
	}

	return nil
}
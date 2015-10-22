package client

import (
	"fmt"
    "net"
    "message-delivery-system/src/messages"
    "message-delivery-system/src/utils"
)

/** 
	MessageClient interface defines what methods the message client should have
*/
type MessageClient interface {
	
	// Connect to the hub listening on the specified address
	Connect(laddr string) error
	
	// Echo message
	Echo(text string) (string, error)
	
	// Identity message
	Identity() (uint64, error)
	
	// List message
	//List() ([]uint64, error)
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
	ECHO REQUEST
*/
func (t *TcpMessageClient) Echo(text string) (string, error) {
	
	request := message.EchoRequest{text}
	response, err := t.sendRequest(request)
	
	return response.(message.EchoResponse).Text, err
}

/**
	ID REQUEST
*/
func (t *TcpMessageClient) Identity() (uint64, error) {
	
	request := message.IdRequest{}
	response, err := t.sendRequest(request)
	
	return response.(message.IdResponse).Id, err
}

/**
	LIST REQUEST
*/
func (t *TcpMessageClient) List() ([]uint64, error) {
	
	request := message.ListRequest{}
	response, err := t.sendRequest(request)
	
	return response.(message.ListResponse).List, err
}

/**
	SEND REQUEST
*/
func (t *TcpMessageClient) sendRequest(request message.Message) (message.Message, error) {
	
	// Encode request
	encoded := message.Encode(request)
	
	// Send request
	err := utils.Write(t.conn, encoded)
	
	if err != nil { 
		fmt.Println("Error sending request:", err.Error())
		return nil, err
	}
	
	// Get response
	response, err := utils.Read(t.conn)
	
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return nil, err
	}
	
	// Decode response
	decoded, err := message.Decode(response)
		
	if err != nil {
		fmt.Println("Error decoding response:", err.Error())
		return nil, err
	}
	
	return decoded, nil
}
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
	List() ([]uint64, error)
	
	// Relay message
	Relay(receivers []uint64, payload []byte)
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
	
	request := messages.EchoRequest{text}
	t.sendRequest(request)
	response, err := t.WaitForResponse()
	
	return response.(messages.EchoResponse).Text, err
}

/**
	ID REQUEST
*/
func (t *TcpMessageClient) Identity() (uint64, error) {
	
	request := messages.IdRequest{}
	t.sendRequest(request)
	response, err := t.WaitForResponse()
	
	return response.(messages.IdResponse).Id, err
}

/**
	LIST REQUEST
*/
func (t *TcpMessageClient) List() ([]uint64, error) {
	
	request := messages.ListRequest{}
	t.sendRequest(request)
	response, err := t.WaitForResponse()
	
	return response.(messages.ListResponse).List, err
}

/**
	RELAY REQUEST
*/
func (t *TcpMessageClient) Relay(receivers []uint64, payload []byte) {
	
	request := messages.RelayRequest{receivers, payload}
	t.sendRequest(request)
}

/**
	SEND REQUEST
*/
func (t *TcpMessageClient) sendRequest(request messages.Message) error {
	
	// Encode request
	encoded := messages.Encode(request)
	
	// Send request
	err := utils.Write(t.conn, encoded)
	
	if err != nil { 
		fmt.Println("Error sending request:", err.Error())
		return err
	}
	
	return nil
}

/**
	WAIT FOR RESPONSE
*/
func (t *TcpMessageClient) WaitForResponse() (messages.Response, error) {
	
	// Get response
	response, err := utils.Read(t.conn)
	
	fmt.Println("Response from hub:", string(response))
	
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return nil, err
	}
	
	// Decode response
	decoded, err := messages.Decode(response)
		
	if err != nil {
		fmt.Println("Error decoding response:", err.Error())
		return nil, err
	}
	
	return decoded.(messages.Response), nil
}
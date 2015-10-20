package hub

import (
    "testing"
    "time"
    "message-delivery-system/src/client"
)

/**
	CONFIG
*/
const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func TestEcho(t *testing.T) {

	// Run hub
	go func() {
		err := Run(CONN_TYPE, CONN_HOST, CONN_PORT)
	
		if err != nil {
			t.Error("Server failed to run")
		}
	}()
	
	// Give the hub some time
	time.Sleep(1000 * time.Millisecond)
	
	// Create a client
	conn, err := client.Connect(CONN_TYPE, CONN_HOST, CONN_PORT)
	
	if err != nil {
		t.Error("Client failed to connect")
	}
	
	// Build message
	request := "example text"
	expectedResponse := "echo " + request; 
	
	// Send and receive text
	response, err := client.Send(conn, request)
	
	if err != nil {
		t.Error("Client failed to send text")
	}
	
	if response != expectedResponse  {
		t.Error("Server didn't response the echo message")
	}
}


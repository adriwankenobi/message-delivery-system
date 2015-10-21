package hub

import (
    "testing"
    "time"
    "message-delivery-system/src/client"
    "message-delivery-system/src/utils"
)

/**
	CONFIG
*/
const CONN_HOST = "localhost"
var CONN_LADDR = utils.GetLaddrFreeTCPPort(CONN_HOST)

func TestEcho(t *testing.T) {
	
	// Initialize the hub
	hub := new(TcpMessageHub)
	
	// Schedule cleaning up
	defer hub.Stop()
	
	// Start the hub
	go func() {
		err := hub.Start(CONN_LADDR)
	
		if err != nil {
			t.Error("Server failed to run")
		}
	}()
	
	// Give the hub some time to start
	time.Sleep(time.Second)
	
	// Start 5 clients
	for i := 0; i < 5; i++ {
		go func(id int) {
		
			// Initialize the client
			cli := new(client.TcpMessageClient)
	
			// Connect to the hub
			err := cli.Connect(CONN_LADDR)
	
			if err != nil {
				t.Error("Client failed to connect")
			}
			
			// Send 5 messages
			for i := 0; i < 5; i++ {
			
				// Build message
				request := "ping"
				expectedResponse := "echo " + request; 
				
				// Send request and get response
				response, err := cli.SendAndGet(request)
				
				if err != nil {
					t.Error("Client failed")
				}
				
				if response != expectedResponse  {
					t.Error("Reponse and expected response don't match")
				}
				
				// Wait a bit for the next message
				time.Sleep(100 * time.Millisecond)
			}	
			
		}(i)
	}		
	
	// Wait for the clients to complete
	time.Sleep(2 * time.Second)
}
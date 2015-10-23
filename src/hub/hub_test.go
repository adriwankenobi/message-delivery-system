package hub

import (
    "testing"
    "time"
    "strconv"
    "message-delivery-system/src/client"
    "message-delivery-system/src/utils"
    "message-delivery-system/src/messages"
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
	
	doTestMethods(t)
	doTestRelay(t)
	doTestSomeConcurrency(t)
	
	time.Sleep(2 * time.Second) // Wait a bit to complete
}

func doTestSomeConcurrency(t *testing.T) {

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
			
				testEcho(t, cli, id, i)
				
				// Wait a bit for the next message
				time.Sleep(100 * time.Millisecond)
			}
			
		}(i)
	}		
}

func testEcho(t *testing.T, cli *client.TcpMessageClient, id int, message int) {

	// Build message
	request := "ping$"+strconv.Itoa(id)+"$"+strconv.Itoa(message)
	expectedResponse := request;
				
	// Send Echo request
	response, err := cli.Echo(request)
				
	if err != nil {
		t.Error("Echo failed")
	}
				
	if response != expectedResponse  {
		t.Error("Response and expected response don't match")
	}
}

func doTestMethods(t *testing.T) {

	// Start 5 clients
	for i := 0; i < 5; i++ {
	
		// Initialize the client
		cli := new(client.TcpMessageClient)
	
		// Connect to the hub
		err := cli.Connect(CONN_LADDR)
	
		if err != nil {
			t.Error("Client failed to connect")
		}
			
		// Send 5 messages
		for j := 0; j < 5; j++ {
			
			testId(t, cli, i)
			
			// Wait a bit for the next message
			time.Sleep(100 * time.Millisecond)
			
			testList(t, cli, i)
				
			// Wait a bit for the next message
			time.Sleep(100 * time.Millisecond)
		}
	}		
}

func testId(t *testing.T, cli *client.TcpMessageClient, id int) {

	expectedResponse := uint64(id+1);
				
	// Send Id request
	response, err := cli.Identity()
				
	if err != nil {
		t.Error("Id failed")
	}
				
	if response != expectedResponse  {
		t.Error("Response and expected response don't match")
	}
}

func testList(t *testing.T, cli *client.TcpMessageClient, id int) {
				
	// Send List request
	response, err := cli.List()
				
	if err != nil {
		t.Error("List failed")
	}
	
	if len(response) != id {
		t.Error("Reponse and expected response don't match - wrong size")
	}	
	
	for j := 1; j < id+1; j++ {
		found := false
		for _, a := range response {
	        if a == uint64(j) {
	            found = true
	            break
	        }
	    }
		if !found {
			t.Error("Response and expected response don't match - id missing")
		}
	}
}

func doTestRelay(t *testing.T) {

	// Client 1
	cli1 := new(client.TcpMessageClient)
	cli1.Connect(CONN_LADDR)
	cli1.Identity()
	
	// Client 2
	cli2 := new(client.TcpMessageClient)
	cli2.Connect(CONN_LADDR)
	cliId2, _ := cli2.Identity()
	
	// Client 3
	cli3 := new(client.TcpMessageClient)
	cli3.Connect(CONN_LADDR)
	cliId3, _ := cli3.Identity()
	
	expectedResponse := "helloworld"
	
	// Set client 2 to wait
	go func() {
		response, err := cli2.WaitForResponse()
		if err != nil {
			t.Error("Relay to client 2 failed")
		}
		
		if string(response.(messages.RelayResponse).Body) != expectedResponse {
			t.Error("Response and expected response don't match")
		}
		
	}()
		
	// Set client 3 to wait
	go func() {
		response, err := cli3.WaitForResponse()
		if err != nil {
			t.Error("Relay to client 3 failed")
		}
		
		if string(response.(messages.RelayResponse).Body) != expectedResponse {
			t.Error("Response and expected response don't match")
		}
	}()
	
	// Send message from client 1 to client 2 and 3
	cli1.Relay([]uint64{cliId2, cliId3}, []byte(expectedResponse))
	
	time.Sleep(2 * time.Second) // Wait a bit to complete
		
}
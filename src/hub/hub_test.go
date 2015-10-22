package hub

import (
    "testing"
    "time"
    "strconv"
    "message-delivery-system/src/client"
    "message-delivery-system/src/utils"
)

/**
	CONFIG
*/
const CONN_HOST = "localhost"
var CONN_LADDR = utils.GetLaddrFreeTCPPort(CONN_HOST)

func Test(t *testing.T) {
	
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

	//doTestId(t)
	//doTestEcho(t)
	doTestList(t)
	
	// Wait for the tests to complete
	time.Sleep(2 * time.Second)
}

func doTestEcho(t *testing.T) {

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

func doTestId(t *testing.T) {

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

func doTestList(t *testing.T) {

	// Start 5 clients
	for i := 0; i < 3; i++ {
	
		// Initialize the client
		cli := new(client.TcpMessageClient)
	
		// Connect to the hub
		err := cli.Connect(CONN_LADDR)
	
		if err != nil {
			t.Error("Client failed to connect")
		}
			
		// Send 5 messages
		for j := 0; j < 1; j++ {
			
			testList(t, cli, i)
				
			// Wait a bit for the next message
			time.Sleep(100 * time.Millisecond)
		}
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
			t.Error("Reponse and expected response don't match - id missing")
		}
	}
}
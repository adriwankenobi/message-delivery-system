# message-delivery-system
Simple message delivery system

- Trying Go language! I like it so far!
- TCP server accepting (using a simple client for testing purposes)
- 1 - Echo message
- 2 - Identify yourself
- 3 - List clients connected
- 4 - Relay message to other clients (up to 1024KB binary-encoded payload and 255 clients)

# TODO
- Generic encoding function for Relay message
- Create a generic holder interface and two implementations for "Reader" and "Writer" in net package using "net.Conn" and "bytes.Buffer"
- Code preventing against wrong inputs
- And their tests
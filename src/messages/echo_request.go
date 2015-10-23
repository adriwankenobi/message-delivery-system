package messages

/**
	ECHO REQUEST
*/
type EchoRequest struct {
	
	// text
	Text string
}

/**
	CUSTOM CONSTRUCTOR
*/
func NewEchoRequest(data []byte) EchoRequest {
	return EchoRequest{string(data)}
}

/**
	GET MESSAGE TYPE
*/
func (e EchoRequest) GetMessageType() MessageType {
	return EchoRequestMessage
}

/**
	GET BINARY DATA
*/
func (e EchoRequest) GetData() []byte {
	return []byte(e.Text)
}
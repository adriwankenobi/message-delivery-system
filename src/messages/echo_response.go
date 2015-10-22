package message

/**
	ECHO RESPONSE
*/
type EchoResponse struct {
	
	// text
	Text string
}

/**
	CUSTOM CONSTRUCTOR
*/
func NewEchoResponse(data []byte) EchoResponse {
	return EchoResponse{string(data)}
}

/**
	GET MESSAGE TYPE
*/
func (e EchoResponse) GetMessageType() MessageType {
	return EchoResponseMessage
}

/**
	GET BINARY DATA
*/
func (e EchoResponse) GetData() []byte {
	return []byte(e.Text)
}
package messages

/**
	ECHO RESPONSE
*/
type EchoResponse struct {
	
	// text
	Text string
	
	// receiver
	Receiver uint64
}

/**
	CUSTOM CONSTRUCTOR
*/
func NewEchoResponse(data []byte) EchoResponse {
	return EchoResponse{Text:string(data)}
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

/**
	GET RECEIVERS
*/
func (e EchoResponse) GetReceiverIds() []uint64 {
	var receivers []uint64
	return append(receivers, e.Receiver)
}
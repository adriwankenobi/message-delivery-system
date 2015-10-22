package message

/**
	ID REQUEST
*/
type IdRequest struct {
	
}

/**
	CUSTOM CONSTRUCTOR
*/
func NewIdRequest(data []byte) IdRequest {
	return IdRequest{}
}

/**
	GET MESSAGE TYPE
*/
func (i IdRequest) GetMessageType() MessageType {
	return IdRequestMessage
}

/**
	GET BINARY DATA
*/
func (i IdRequest) GetData() []byte {
	return nil
}
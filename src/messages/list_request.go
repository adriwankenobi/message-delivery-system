package message

/**
	LIST REQUEST
*/
type ListRequest struct {
	
}

/**
	CUSTOM CONSTRUCTOR
*/
func NewListRequest(data []byte) ListRequest {
	return ListRequest{}
}

/**
	GET MESSAGE TYPE
*/
func (l ListRequest) GetMessageType() MessageType {
	return ListRequestMessage
}

/**
	GET BINARY DATA
*/
func (l ListRequest) GetData() []byte {
	return nil
}
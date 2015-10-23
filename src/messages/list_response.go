package messages

import (
	"bytes"
	"message-delivery-system/src/utils"
)

/**
	LIST RESPONSE
*/
type ListResponse struct {
	
	// list of ids
	List []uint64
	
	// receiver
	Receiver uint64
}

/**
	CUSTOM CONSTRUCTOR
*/
func NewListResponse(data []byte) ListResponse {
	buf := bytes.NewReader(data)
	list, _ := utils.ByteArrayToUint64List(buf, data)
	return ListResponse{List:list}
}

/**
	GET MESSAGE TYPE
*/
func (l ListResponse) GetMessageType() MessageType {
	return ListResponseMessage
}

/**
	GET BINARY DATA
*/
func (l ListResponse) GetData() []byte {
	buf := new(bytes.Buffer)
	data, _ := utils.Uint64ListToByteArray(buf, l.List)
	return data
}

/**
	GET RECEIVERS
*/
func (l ListResponse) GetReceiverIds() []uint64 {
	var receivers []uint64
	return append(receivers, l.Receiver)
}
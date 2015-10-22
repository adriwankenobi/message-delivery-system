package message

import (
	"strconv"
)

/**
	ID RESPONSE
*/
type IdResponse struct {
	
	// id
	Id uint64
}

/**
	CUSTOM CONSTRUCTOR
*/
func NewIdResponse(data []byte) IdResponse {
	
	id, _ := strconv.ParseUint(string(data), 10, 64)
	return IdResponse{id}
}

/**
	GET MESSAGE TYPE
*/
func (i IdResponse) GetMessageType() MessageType {
	return IdResponseMessage
}

/**
	GET BINARY DATA
*/
func (i IdResponse) GetData() []byte {
	
	dataString := strconv.FormatUint(i.Id, 10)
	
	return []byte(dataString)
}
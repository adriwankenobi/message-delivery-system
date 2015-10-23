package messages

import (
	"bytes"
	"message-delivery-system/src/utils"
)

/**
	RELAY RESPONSE
*/
type RelayResponse struct {
	
	// body message
	Body []byte
	
	// list of receiver ids
	Receivers []uint64
}

/**
	CUSTOM CONSTRUCTOR
*/
func NewRelayResponse(data []byte) RelayResponse {
	
	buf := bytes.NewReader(data)
	body, _ := utils.Read(buf)
	
	return RelayResponse{Body:body}
}

/**
	GET MESSAGE TYPE
*/
func (r RelayResponse) GetMessageType() MessageType {
	return RelayResponseMessage
}

/**
	GET BINARY DATA
*/
func (r RelayResponse) GetData() []byte {
	
	buf := new(bytes.Buffer)
	utils.Write(buf, r.Body)
	
	return buf.Bytes()
}

/**
	GET RECEIVERS
*/
func (r RelayResponse) GetReceiverIds() []uint64 {
	return r.Receivers
}
package messages

import (
	"bytes"
	"message-delivery-system/src/utils"
)

/**
	RELAY REQUEST
	TODO: Using binary encoding. Should a function parameter so any encoding could be used
*/
type RelayRequest struct {
	
	// list of receiver ids
	Receivers []uint64
	
	// body message
	Body []byte
}

/**
	CUSTOM CONSTRUCTOR
*/
func NewRelayRequest(data []byte) RelayRequest {
	
	buf := bytes.NewReader(data)
	receivers, _ := utils.ByteArrayToUint64List(buf, data)
	body, _ := utils.Read(buf)
	
	return RelayRequest{receivers, body}
}

/**
	GET MESSAGE TYPE
*/
func (r RelayRequest) GetMessageType() MessageType {
	return RelayRequestMessage
}

/**
	GET BINARY DATA
*/
func (r RelayRequest) GetData() []byte {
	
	buf := new(bytes.Buffer)
	utils.Uint64ListToByteArray(buf, r.Receivers)
	utils.Write(buf, r.Body)
	
	return buf.Bytes()
}
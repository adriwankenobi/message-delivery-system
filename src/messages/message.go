package message

import (
	"strconv"
	"strings"
	"errors"
)

/**
	MESSAGE TYPE
*/
type MessageType int
const (
        EchoRequestMessage MessageType = iota
        EchoResponseMessage
        IdRequestMessage
        IdResponseMessage
        ListRequestMessage
        ListResponseMessage
        RelayRequestMessage
        RelayResponseMessage
)

/**
	MESSAGE
*/
type Message interface {

	// Get the message type
	GetMessageType() MessageType
	
	// Gets the binary data
	GetData() []byte
	
}

const DELIM = "#"

/**
	ENCODE MESSAGE
*/
func Encode(message Message) []byte {
	
	// <MESSAGE_TYPE> + <DELIM> + <MESSAGE_DATA>
	
	messageType := strconv.Itoa(int(message.GetMessageType()))
	messageData := message.GetData()
	
	return append([]byte(messageType + DELIM), messageData...)
}

/**
	DECODE MESSAGE
*/
func Decode(encoded []byte) (Message, error) {
	
	//<MESSAGE_TYPE> + <DELIM> + <MESSAGE_DATA>
	
	messageParts := strings.Split(string(encoded), DELIM)
	
	if len(messageParts) != 2 {
		err := errors.New("Cannot decode message: wrong input data")
    	return nil, err
	}
	
	code, err := strconv.Atoi(messageParts[0])
	
	if err != nil {
		return nil, err
	}
	
	messageType := MessageType(code)
	messageData := []byte(messageParts[1])
	
	var message Message
	
	switch messageType {
    case EchoRequestMessage:
        message = NewEchoRequest(messageData)
    case EchoResponseMessage:
        message = NewEchoResponse(messageData)
    case IdRequestMessage:
        message = NewIdRequest(messageData)
    case IdResponseMessage:
        message = NewIdResponse(messageData)
    case ListRequestMessage:
        message = NewListRequest(messageData)
    case ListResponseMessage:
        message = NewListResponse(messageData)
    default:
    	err := errors.New("Cannot decode message: wrong type")
    	return nil, err
    }
    
	return message, nil
}





// MessageEncoder defines the message encoder
type MessageEncoder func(receivers []uint64, payload []byte) []byte
package services

import (
	"message-delivery-system/src/messages"
)

func HandleEcho(request messages.EchoRequest, id uint64) messages.EchoResponse {
	return messages.EchoResponse{request.Text, id}
}

func HandleId(request messages.IdRequest, id uint64) messages.IdResponse {
	return messages.IdResponse{id, id}
}

func HandleList(request messages.ListRequest, id uint64, clients []uint64) messages.ListResponse {
	var list []uint64
	for _, someId := range clients {
		if someId != id {
			list = append(list, someId)
		}
	}
	return messages.ListResponse{list, id}
}

func HandleRelay(request messages.RelayRequest) messages.RelayResponse {
	return messages.RelayResponse{request.Body, request.Receivers}
}
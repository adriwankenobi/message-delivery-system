package services

import (
	"message-delivery-system/src/messages"
)

func HandleEcho(request message.EchoRequest) message.EchoResponse {
	return message.EchoResponse{request.Text}
}

func HandleId(request message.IdRequest, id uint64) message.IdResponse {
	return message.IdResponse{id}
}

func HandleList(request message.ListRequest, id uint64, clients []uint64) message.ListResponse {
	var list []uint64
	for _, someId := range clients {
		if someId != id {
			list = append(list, someId)
		}
	}
	return message.ListResponse{list}
}
package message

import (
	"bytes"
	"encoding/binary"
)

/**
	LIST RESPONSE
*/
type ListResponse struct {
	
	// list of ids
	List []uint64
}

/**
	CUSTOM CONSTRUCTOR
*/
func NewListResponse(data []byte) ListResponse {
	
	// Define buffer
	buf := bytes.NewReader(data)
	
	// Read how many (fixed size int32)
	var size int32
	binary.Read(buf, binary.LittleEndian, &size)
	list := make([]uint64, int(size))
	
	// Read elements
	for i := 0; i < int(size); i++ {
		var e uint64
		binary.Read(buf, binary.LittleEndian, &e)
		list[i] = e
	}

	return ListResponse{list}
}

/**
	GET MESSAGE TYPE
*/
func (i ListResponse) GetMessageType() MessageType {
	return ListResponseMessage
}

/**
	GET BINARY DATA
*/
func (i ListResponse) GetData() []byte {
	
	// Define buffer
	buf := new(bytes.Buffer)
	
	// Write list len (fixed size int32)
	binary.Write(buf, binary.LittleEndian, int32(len(i.List)))

	// Write list elements
	for _, e := range i.List {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return buf.Bytes()
}
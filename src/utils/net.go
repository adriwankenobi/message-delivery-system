package utils

import (
	"net"
	"bytes"
	"errors"
	"encoding/binary"
)

const CONN_TYPE = "tcp"
const MAX_SIZE_STREAM = 1024000 // 1024KB
const MAX_ID_LIST_SIZE = 255

/**
	GET A LADDR USING A FREE TCP PORT
*/
func GetLaddrFreeTCPPort(host string) string {
	l, _ := net.Listen(CONN_TYPE, host+":0")
	defer l.Close()
	return l.Addr().String()
}

/**
	WRITE INTO GENERIC WRITER (net.Conn or *bytes.Reader)
	TODO: Very ugly! Replace switch with an interface and two implementations
*/
func Write(writer interface{}, data []byte) error {
	
	if (len(data) > MAX_SIZE_STREAM) {
		return errors.New("Stream data too large")
	}
	
	// Write size (fixed size int32)
	var err error
	switch conn := writer.(type) {
	case *bytes.Buffer:
	    err = binary.Write(conn, binary.LittleEndian, int32(len(data)))
	case net.Conn:
	    err = binary.Write(conn, binary.LittleEndian, int32(len(data)))
	}
	
	if err != nil {
		return err
	}
	
	// Write data
	
	switch conn := writer.(type) {
	case *bytes.Buffer:
	    _, err = conn.Write(data)
	case net.Conn:
	    _, err = conn.Write(data)
	}
	return err
}

/**
	READ FROM GENERIC READER (net.Conn or *bytes.Reader)
	TODO: Very ugly! Replace switch with an interface and two implementations
*/
func Read(reader interface{}) ([]byte, error) {

	// Read size (fixed size int32)
	var expectedBytes int32
	var err error
	switch conn := reader.(type) {
	case *bytes.Reader:
	    err = binary.Read(conn, binary.LittleEndian, &expectedBytes)
	case net.Conn:
	    err = binary.Read(conn, binary.LittleEndian, &expectedBytes)
	}
	
	if err != nil {
		return nil, err
	}
	
	if (expectedBytes > MAX_SIZE_STREAM) {
		return nil, errors.New("Stream data too large")
	}
	
	// Read data
	totalBytesRead := 0
	
	// Make a fixed buffer to hold incoming data
	// MAX length is 1024KB, but that's too much of a memory
	// Instead using 102400 B = 102,4 KB which is much more reasonable than a ~1MB buffer
	// This means ~10 loops in worst case
	buf := make([]byte, MAX_SIZE_STREAM/10)
		
	// Make the dynamic buffer for the total data
	var data bytes.Buffer
		
	for {
	
		// Read the incoming connection into the buffer
		var bytesRead int
		var err error
		switch conn := reader.(type) {
		case *bytes.Reader:
		    bytesRead, err = conn.Read(buf)
		case net.Conn:
		    bytesRead, err = conn.Read(buf)
		}
		
		totalBytesRead = totalBytesRead + bytesRead
		
		if err != nil {
			return nil, err
		}
		
		data.Write(buf[:bytesRead])
		
		if totalBytesRead >= int(expectedBytes) {
			break // Reached end of data stream
		}
	}
	
	return data.Bytes(), nil
}

/**
	 []byte to []uint64
*/
func ByteArrayToUint64List(buf *bytes.Reader, data []byte) ([]uint64, error) {
	
	// Read how many (fixed size int32)
	var size int32
	binary.Read(buf, binary.LittleEndian, &size)
	
	if (size > MAX_ID_LIST_SIZE) {
		return nil, errors.New("Too many users")
	}
	
	list := make([]uint64, int(size))
	
	// Read elements
	for i := 0; i < int(size); i++ {
		var e uint64
		binary.Read(buf, binary.LittleEndian, &e)
		list[i] = e
	}
	
	return list, nil
}

/**
	 []uint64 to []byte
*/
func Uint64ListToByteArray(buf *bytes.Buffer, data []uint64) ([]byte, error) {
	
	if (len(data) > MAX_ID_LIST_SIZE) {
		return nil, errors.New("Too many users")
	}
	
	// Write list len (fixed size int32)
	binary.Write(buf, binary.LittleEndian, int32(len(data)))

	// Write list elements
	for _, e := range data {
		binary.Write(buf, binary.LittleEndian, e)
	}
	
	return buf.Bytes(), nil
}

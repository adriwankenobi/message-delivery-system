package utils

import (
	"net"
	"bytes"
	"encoding/binary"
)

const CONN_TYPE = "tcp"

/**
	GET A LADDR USING A FREE TCP PORT
*/
func GetLaddrFreeTCPPort(host string) string {
	l, _ := net.Listen(CONN_TYPE, host+":0")
	defer l.Close()
	return l.Addr().String()
}

/**
	WRITE INTO CONNECTION
*/
func Write(conn net.Conn, data []byte) error {
	
	// Write size (fixed size int32)
	err := binary.Write(conn, binary.LittleEndian, int32(len(data)))
	
	if err != nil {
		return err
	}
	
	// Write data
	_, err = conn.Write(data)
	return err
}

/**
	READ FROM CONNECTION
*/
func Read(conn net.Conn) ([]byte, error) {
	
	// Read size (fixed size int32)
	var expectedBytes int32
	err := binary.Read(conn, binary.LittleEndian, &expectedBytes)
	
	if err != nil {
		return nil, err
	}
	
	// Read data
	totalBytesRead := 0
	
	// Make a fixed buffer to hold incoming data
	// 102400 B = 102,4 KB which is much more reasonable than a ~1MB buffer
	// This means 10 loops in worst case
	buf := make([]byte, 3)
		
	// Make the dynamic buffer for the total data
	var data bytes.Buffer
		
	for {
	
		// Read the incoming connection into the buffer
		bytesRead, err := conn.Read(buf)
		totalBytesRead = totalBytesRead + bytesRead
		
		if err != nil {
			return nil, err
		}
		
		data.Write(buf[:bytesRead])
		
		if totalBytesRead >= int(expectedBytes) {
			break
		}
	}
	
	return data.Bytes(), nil
}
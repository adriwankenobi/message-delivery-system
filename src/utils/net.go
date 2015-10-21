package utils

import (
	"net"
)

const CONN_TYPE = "tcp"

/**
	GETS A LADDR USING A FREE TCP PORT
*/
func GetLaddrFreeTCPPort(host string) string {
	l, _ := net.Listen(CONN_TYPE, host+":0")
	defer l.Close()
	return l.Addr().String()
}
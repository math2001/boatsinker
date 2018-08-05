package utils

import "net"

type Message struct {
	Data  map[string]interface{}
	From  *net.Conn
	Count int
}

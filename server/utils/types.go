package utils

import (
	"fmt"
	"net"
)

type Message struct {
	Data  map[string]interface{}
	From  *net.Conn
	Count int
}

type Error struct {
	Err  error
	Data interface{}
}

func (e *Error) String() {
	fmt.Printf("Error: %s\nData: %#v", e.Err, e.Data)
}

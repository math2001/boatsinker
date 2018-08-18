package utils

import (
	"fmt"
	"net"
)

// Message stores information about web socket message
// It is used to represent messages that have been received and messages that
// have to be sent
type Message struct {
	Data  map[string]interface{}
	From  *net.Conn
	Count int
}

func (m Message) String() string {
	return fmt.Sprintf("#%d %v", m.Count, m.Data)
}

// Error is used to pass an error around with some data
type Error struct {
	Err  error
	Data interface{}
}

func (e *Error) String() {
	fmt.Printf("Error: %s\nData: %#v", e.Err, e.Data)
}

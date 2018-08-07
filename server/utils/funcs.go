package utils

import (
	"log"
	"net"
)

// usage: MakeMap("first", someobject, "second", another, )
// returns map[string]interface{}{first: someobject, second, another}
func MakeMap(args ...interface{}) map[string]interface{} {
	var m = make(map[string]interface{})
	for i, arg := range args {
		if i%2 == 1 {
			continue
		}
		str, ok := arg.(string)
		if !ok {
			log.Fatal("Key should be string (every second argument)")
		}
		m[str] = args[i+1]
	}
	return m
}

// usage: NewMessage(conn, "kind", "something", "data": obj)
func NewMessage(conn *net.Conn, args ...interface{}) Message {
	return Message{
		From:  conn,
		Data:  MakeMap(args...),
		Count: 0,
	}
}

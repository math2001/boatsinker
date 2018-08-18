package utils

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

// MakeMap makes a map from a list of argument, where the keys are strings
// usage: MakeMap("first", someobject, "second", another)
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

// NewMessage creates a new message from
// usage: NewMessage(conn, "kind", "something", "data": obj)
func NewMessage(conn *net.Conn, args ...interface{}) Message {
	return Message{
		From:  conn,
		Data:  MakeMap(args...),
		Count: 0,
	}
}

// ErrorFrom concatenates a list of errors into one error.
func ErrorFrom(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	var b strings.Builder
	for _, err := range errs {
		fmt.Fprintln(&b, err)
	}
	return errors.New(b.String())
}

// Must returns makes sure the error is nil.
func Must(val interface{}, err error) interface{} {
	if err != nil {
		log.Fatal(err)
	}
	return val
}

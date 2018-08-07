package em

import (
	"fmt"
	"log"
)

// A simple event manager
// Bind listeners with on and trigger them with On

type callback func(interface{}) error

type EventManager struct {
	// purely for debug purposes
	name   string
	events map[string][]callback
}

func logevent(name string, args interface{}) {
	fmt.Printf("%s: ", name)
	if _, ok := args.(fmt.Stringer); ok {
		fmt.Printf("%s", args)
	} else {
		fmt.Printf("%#v", args)
	}
	fmt.Printf("\n")
}

func (e *EventManager) On(name string, cb callback) {
	_, ok := e.events[name]
	if !ok {
		e.events[name] = make([]callback, 0)
	}
	e.events[name] = append(e.events[name], cb)
}

func (e *EventManager) Emit(name string, args interface{}) []error {
	// handlers should return there is a user error. If it's a dev error,
	// it should panic *itself*.
	var errs []error
	logevent(name, args)
	callbacks, ok := e.events[name]
	if !ok {
		log.Printf("No handlers for the event '%s'", name)
		return errs
	}
	for _, cb := range callbacks {
		err := cb(args)
		if err != nil {
			errs = append(errs, err)
			log.Print("Error:", err)
		}
	}
	return errs
}

func NewEventManager(name string) *EventManager {
	return &EventManager{name: name, events: make(map[string][]callback)}
}

var em = NewEventManager("default")

func Emit(name string, args interface{}) []error {
	return em.Emit(name, args)
}

func On(name string, cb callback) {
	em.On(name, cb)
}

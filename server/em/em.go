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

func (e *EventManager) On(name string, cb callback) {
	_, ok := e.events[name]
	if !ok {
		e.events[name] = make([]callback, 0)
	}
	e.events[name] = append(e.events[name], cb)
}

func (e *EventManager) Emit(name string, args interface{}) []error {
	log.Printf("%#v %v", name, args)
	callbacks, ok := e.events[name]
	if !ok {
		return []error{fmt.Errorf("No handlers for the event %#v", name)}
	}
	var errs []error
	for _, cb := range callbacks {
		err := cb(args)
		if err != nil {
			errs = append(errs, err)
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

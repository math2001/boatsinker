package em

import (
	"fmt"
	"log"
	"strings"

	"github.com/math2001/boatsinker/server/utils"
)

// A simple event manager
// Bind listeners with on and trigger them with On

type callback func(interface{}) error

// EventManager is a simple pub/sub
type EventManager struct {
	// purely for debug purposes
	name   string
	events map[string][]callback
}

func logevent(name string, args interface{}) {
	var b strings.Builder
	fmt.Fprintf(&b, "%s: ", name)
	if _, ok := args.(fmt.Stringer); ok {
		fmt.Fprintf(&b, "%s", args)
	} else {
		fmt.Fprintf(&b, "%v", args)
	}
	log.Print(b.String())
}

// On adds a listener (a function) to an event
func (e *EventManager) On(name string, cb callback) {
	_, ok := e.events[name]
	if !ok {
		e.events[name] = make([]callback, 0)
	}
	e.events[name] = append(e.events[name], cb)
}

// Emit triggers every listener bound to the event with the given argument
func (e *EventManager) Emit(name string, args interface{}) error {
	// handlers should return there is a user error. If it's a dev error,
	// it should panic *itself*.
	var errs []error
	logevent(name, args)
	callbacks, ok := e.events[name]
	if !ok {
		log.Printf("No handlers for the event '%s'", name)
		return nil
	}
	for _, cb := range callbacks {
		err := cb(args)
		if err != nil {
			errs = append(errs, err)
			log.Print("Error:", err)
		}
	}
	return utils.ErrorFrom(errs)
}

// NewEventManager creates a new event manager
func NewEventManager(name string) *EventManager {
	return &EventManager{name: name, events: make(map[string][]callback)}
}

var em = NewEventManager("default")

// Emit calls the default event manager's Emit method
func Emit(name string, args interface{}) error {
	return em.Emit(name, args)
}

// On calls the default event manager's On method
func On(name string, cb callback) {
	em.On(name, cb)
}

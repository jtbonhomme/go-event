// Package event provides simple event emitter for Golang.
package event

import (
	"github.com/google/uuid"
)

// A Handler responds to an Event triggering.
type Handler interface {
	// On() is the callbck function called when an event is triggered.
	On(string)
	// ID() returns the internal handler ID.
	ID() uuid.UUID
}

// An Event is used to notify handlers with a given message.
// Every Event typically has internal list of event handlers.
// To be added to the internal handlers list, the function Register is used.
// To be removed to the internal handlers list, the function Unregister is used.
type Event struct {
	handlers []Handler
}

// Register adds an handler to an event handlers list.
func (e *Event) Register(h Handler) {
	e.handlers = append(e.handlers, h)
}

// Unregister removes an handler from an event handlers list.
func (e *Event) Unregister(handler Handler) {
	for i, h := range e.handlers {
		if h.ID() == handler.ID() {
			// Remove element i from handlers list.
			e.handlers[i] = e.handlers[len(e.handlers)-1]
			e.handlers[len(e.handlers)-1] = nil
			e.handlers = e.handlers[:len(e.handlers)-1]
			break
		}
	}
}

// Emit triggers all event handlers Run() function.
func (e *Event) Emit(event string) {
	for _, h := range e.handlers {
		h.On(event)
	}
}

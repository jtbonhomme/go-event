// Package event provides simple event emitter for Golang.
package event

import (
	"github.com/google/uuid"
)

// A Handler responds to an Event triggering.
type Handler interface {
	On(string)
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
			e.handlers[i] = e.handlers[len(e.handlers)-1] // Copy last element to index i.
			e.handlers[len(e.handlers)-1] = nil           // Erase last element (write zero value).
			e.handlers = e.handlers[:len(e.handlers)-1]   // Truncate slice.
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

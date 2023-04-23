package event_test

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/jtbonhomme/go-event"
)

type concreteHandler struct {
	count int
	id    uuid.UUID
}

func (e concreteHandler) On(eventType string) {
	if eventType == "incrCounter" {
		e.count++
	}
	fmt.Printf("handler %s has been notified of an event emission, new value is %d\n", e.id, e.count)
}

func (e concreteHandler) ID() uuid.UUID {
	return e.id
}

func (e concreteHandler) Value() int {
	return e.count
}

func ExampleEvent() {
	e := &event.Event{}
	h := &concreteHandler{
		id: uuid.New(),
	}
	e.Register(h)
	fmt.Printf("handler value is %d\n", h.Value())
	e.Emit("incrCounter")
	fmt.Printf("handler value is %d\n", h.Value())
}

package event_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"

	"github.com/jtbonhomme/go-event"
)

type testHandler struct {
	count int
	id    uuid.UUID
}

func (e *testHandler) On(eventType string) {
	if eventType == "incrCounter" {
		e.count++
	}
	fmt.Printf("handler %s has been notified of an event emission, new value is %d\n", e.id, e.count)
}

func (e *testHandler) ID() uuid.UUID {
	return e.id
}

func (e *testHandler) Value() int {
	return e.count
}

func TestEvent(t *testing.T) {
	e := &event.Event{}
	h1 := &testHandler{
		id: uuid.New(),
	}
	h2 := &testHandler{
		id: uuid.New(),
	}
	if h1.count != 0 {
		t.Errorf("handler expected value is 0, got %d", h1.count)
	}
	if h2.count != 0 {
		t.Errorf("handler expected value is 0, got %d", h1.count)
	}
	e.Register(h1)
	e.Emit("incrCounter")
	if h1.count != 1 {
		t.Errorf("handler expected value is 1, got %d", h1.count)
	}
	e.Register(h2)
	e.Emit("incrCounter")
	if h1.count != 2 {
		t.Errorf("handler expected value is 2, got %d", h1.count)
	}
	if h2.count != 1 {
		t.Errorf("handler expected value is 1, got %d", h2.count)
	}
	e.Unregister(h2)
	e.Emit("incrCounter")
	if h1.count != 3 {
		t.Errorf("handler expected value is 3, got %d", h1.count)
	}
	if h2.count != 1 {
		t.Errorf("handler expected value is 1, got %d", h2.count)
	}
	e.Unregister(h1)
	e.Emit("incrCounter")
	if h1.count != 3 {
		t.Errorf("handler expected value is 3, got %d", h1.count)
	}
	if h2.count != 1 {
		t.Errorf("handler expected value is 1, got %d", h2.count)
	}
}

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

func (e *testHandler) On(event string) {
	if event == "incrCounter" {
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
	h := &testHandler{
		id: uuid.New(),
	}
	e.Register(h)
	e.Emit("incrCounter")
	if h.count != 1 {
		t.Fatalf("handler expected value is 1, got %d", h.count)
	}
}

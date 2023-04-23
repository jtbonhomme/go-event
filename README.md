# go-event

![Build Status](https://github.com/jtbonhomme/go-event/workflows/Go/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/jtbonhomme/go-event)](https://pkg.go.dev/mod/github.com/jtbonhomme/go-event)

Simple Event observer library for Golang.

## Installation

```sh
go get github.com/jtbonhomme/go-event
```

## Events and Handlers

Observers will register an `handler` to an `event`.
When the event is emmitted, all registered handlers are called.

## Get Started

```
package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jtbonhomme/go-event"
)

type concreteHandler struct {
	count int
	id    uuid.UUID
}

func (e concreteHandler) On(event string) {
	if event == "incrCounter" {
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

func main() {
	e := &event.Event{}
	h := &concreteHandler{
		id: uuid.New(),
	}
	e.Register(h)
	fmt.Printf("handler value is %d\n", h.Value())
	e.Emit("incrCounter")
	fmt.Printf("handler value is %d\n", h.Value())
}
```
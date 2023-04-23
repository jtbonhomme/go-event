# go-event

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/jtbonhomme/go-event)](https://pkg.go.dev/mod/github.com/jtbonhomme/go-event)
[![Go Report Card](https://goreportcard.com/badge/github.com/jtbonhomme/go-event)](https://goreportcard.com/report/github.com/jtbonhomme/go-event)
[![License](https://img.shields.io/github/license/jtbonhomme/go-event.svg)](https://github.com/jtbonhomme/go-event/blob/master/LICENSE)
[![Release](https://img.shields.io/github/v/release/jtbonhomme/go-event.svg)](https://github.com/jtbonhomme/go-event/releases/)
[![Tests](https://github.com/jtbonhomme/go-event/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/jtbonhomme/go-event/actions/workflows/tests.yml)

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
package main

import "log"

type X struct {
	body string
}

type Message interface {
	Body() string
}

type LoggableMessage interface {
	log()
}

func (x *X) Body() string {
	return x.body
}

func (x *X) log() {
	log.Printf("logging body: %q\n", x.body)
}

func send(m Message) {
	log.Printf("sending: %v", m.Body())

	if l, ok := m.(LoggableMessage); ok {
		l.log()
	}
}

func main() {
	x := X{body: "hello, world"}

	send(&x)
}

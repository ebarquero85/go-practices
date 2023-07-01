package main

import (
	"fmt"
)

type OptFunc func(*Options)

type Options struct {
	host string
	id   int
	tls  bool
}

type Server struct {
	options Options
}

func defaultOptions() Options {
	return Options{
		host: "127.0.0.1",
		id:   10,
		tls:  false,
	}
}

func NewServer(opts ...OptFunc) *Server {

	o := defaultOptions()

	for _, fn := range opts {
		fn(&o)
	}

	return &Server{
		options: o,
	}
}

// Options functions
func withTLS(o *Options) {
	o.tls = true
}

func withHost(host string) OptFunc {
	return func(o *Options) {
		o.host = host
	}
}

func withId(id int) OptFunc {
	return func(o *Options) {
		o.id = id
	}
}

func main() {

	//server := NewServer() // Default values

	server := NewServer(
		withTLS,
		withHost("192.168.1.0"),
		withId(99),
	)

	fmt.Printf("%+v\n", server)

}

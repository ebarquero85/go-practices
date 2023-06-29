package main

import "fmt"

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type Server struct {
	Opts
}

type OpFunc func(*Opts)

// =======

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "one",
		tls:     false,
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withMaxConn(mc int) OpFunc {
	return func(opts *Opts) {
		opts.maxConn = mc
	}
}

func withID(id string) OpFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

// =======

// the slice of OpFunc (...OpFunc) give us the posibility of send 0, 1, 2 etc options
func newServer(opts ...OpFunc) *Server {

	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &Server{
		Opts: o,
	}
}

func main() {

	s := newServer(withTLS, withMaxConn(100), withID("two"))
	fmt.Printf("%v", s)

}

package math

import (
	"net/rpc"

	log "github.com/sirupsen/logrus"
)

const (
	Address = "/tmp/sample.sock"
	Network = "unix"
)

type Request struct {
	A, B int
}

type Server int

func (rs *Server) Plus(r Request, res *int) error {
	log.Printf("Request A: %d, B: %d", r.A, r.B)
	*res = r.A + r.B

	log.Printf("Result: %d", *res)
	return nil
}

func Plus(a, b int) (int, error) {
	client, err := rpc.DialHTTP(Network, Address)
	if err != nil {
		return 0, err
	}

	var reply int
	err = client.Call("Server.Plus", Request{A: a, B: b}, &reply)
	return reply, err
}

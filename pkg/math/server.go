package math

import (
	log "github.com/sirupsen/logrus"
	"net/rpc"
)

type Request struct {
	A, B int
}

type Server int

func (rs *Server) Plus(r Request, res *int) error {
	log.Printf("Request A: %d, B: %d", r.A, r.B)
	result := r.A + r.B
	res = &result

	log.Printf("Result: %d", result)
	return nil
}

func Plus(a,b int) (int, error){
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		return 0, err
	}

	var reply int
	err = client.Call("Server.Plus", Request{A: a, B:b}, &reply)
	return reply, err
}

package main

import (
	// "errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B string
}

type Quotient struct {
	Quo, Rem int
}

type Arith string

func (t *Arith) Multiply(args *Args, reply *string) error {
	fmt.Println(args.A)
	*reply = "PONG TO " + args.B
	return nil
}

func main() {
	// Create a new RPC server
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1235")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)

}


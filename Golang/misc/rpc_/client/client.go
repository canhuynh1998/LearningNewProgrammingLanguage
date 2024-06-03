package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B string
}


func main() {

	args := Args{"PING", "1235"}

	client, err := rpc.DialHTTP("tcp", "localhost"+":1235")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("Arith.Multiply", Args{"PING", "1235"}, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %s %s =%s\n", args.A, args.B, reply)
	// quotient := Quotient{}

	// divCall := client.Go("Arith.Divide", args, Quotient{ }, nil)
	// // replyCall := <-divCall.Done // will be equal to divCall
	// // check errors, print, etc.

	// fmt.Println(<-divCall.Done)
	// fmt.Printf("Arith: %d*%d=%d\n", quotient.Quo, quotient.Rem, reply)
}

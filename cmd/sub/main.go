package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func main() {
	//  Prepare our subscriber
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://127.0.0.1:8100")
	subscriber.SetSubscribe("B")
	fmt.Println("started subscriber")
	for {
		//  Read envelope with address
		address, _ := subscriber.Recv(0)
		fmt.Println("got address")
		//  Read message contents
		contents, _ := subscriber.Recv(0)
		fmt.Printf("[%s] %s\n", address, contents)
	}
}
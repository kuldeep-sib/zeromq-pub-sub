//
//  Pubsub envelope publisher.
//

package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"

	"time"
)

func main() {
	//  Prepare our publisher
	fmt.Println("starting publisher")
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect("tcp://127.0.0.1:5563")
	fmt.Println("started publisher2")
	for {
		//  Write two messages, each with an envelope and content
		publisher.Send("185.107.232.99", zmq.SNDMORE)
		publisher.Send("3:orange.fr", 0)
		publisher.Send("185_107_232_90", zmq.SNDMORE)
		publisher.Send("2:test.com", 0)
		fmt.Println("written msgs")
		time.Sleep(time.Second)
	}
}
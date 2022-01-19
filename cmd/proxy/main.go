package main

import (
	zmq "github.com/pebbe/zmq4"

	"log"
)

func main() {
	//  This is where the server sits
	frontend, _ := zmq.NewSocket(zmq.XSUB)
	defer frontend.Close()
	frontend.Bind("tcp://127.0.0.1:5563")

	//  This is our public endpoint for subscribers
	backend, _ := zmq.NewSocket(zmq.XPUB)
	defer backend.Close()
	backend.Bind("tcp://127.0.0.1:8100")

	log.Printf("started Proxy")
	//  Run the proxy until the user interrupts us
	err := zmq.Proxy(frontend, backend, nil)
	log.Fatalln("Proxy interrupted:", err)
}
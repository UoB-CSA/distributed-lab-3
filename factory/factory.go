package main

import ("flag"
	"fmt"
	"net"
	"pairbroker/stubs"
	"net/rpc")


type Factory struct {}

//TODO: Define a Multiply function to be accessed via RPC. 
//Check the previous weeks' examples to figure out how to do this.


func main(){
	pAddr := flag.String("ip", "127.0.0.1:8050", "IP and port to listen on")
	brokerAddr := flag.String("broker","127.0.0.1:8030", "Address of broker instance")
	flag.Parse()
	//TODO: You'll need to set up the RPC server, and subscribe to the running broker instance.
}

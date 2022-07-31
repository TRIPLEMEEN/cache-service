package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type bytes struct {
	Title string
	Body  string
}

func main() {
	var reply bytes

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	a := bytes{"First", "One byte"}
	client.Call("API.Get", a, &reply)

	fmt.Println("The byte: ", reply)

}

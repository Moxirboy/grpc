package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/grpc/proto"
)

var address = "localhost:3000"

func main() {
	for {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		c := proto.NewMessangerClient(conn)
		var message string
		fmt.Println("enter your name : ")
		fmt.Scanln(&message)
		response, err := c.Messager(context.Background(), &proto.Request{Message: message})
		if err != nil {
			panic(err)
		}
		fmt.Println(response)
	}
}

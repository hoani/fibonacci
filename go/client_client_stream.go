// client_client_stream.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	fibonacci "github.com/hoani/fibonacci/proto"
	"google.golang.org/grpc"
)

func getUserInteger() (int, error) {
	var indexStr string
	fmt.Print("enter an index (or nothing to stop): ")
	fmt.Scanln(&indexStr)
	return strconv.Atoi(indexStr)
}

func serverAddress() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "localhost:1337"
}

func main() {
	conn, err := grpc.Dial(serverAddress(),
		grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := fibonacci.NewFibonacciClient(conn)

	stream, err := c.SumIndicies(context.Background())
	if err != nil {
		log.Fatalf("call failed: %v", err)
	}

	for {
		index, err := getUserInteger()
		if err != nil {
			r, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("error closing %v", err)
			}
			log.Printf("Sum: %v", r.Value)
			break
		}
		increment := &fibonacci.Number{Value: int32(index)}
		err = stream.Send(increment)
		if err != nil {
			log.Fatalf("error sending %v", err)
		}
	}
}

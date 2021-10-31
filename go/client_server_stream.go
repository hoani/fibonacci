// client_server_stream.go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	fibonacci "github.com/hoani/fibonacci/proto"
	"google.golang.org/grpc"
)

func getUserInteger() int {
	var indexStr string
	fmt.Print("enter an index integer: ")
	fmt.Scanln(&indexStr)
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		log.Fatalf("entered invalid integer `%v`", indexStr)
	}
	return index
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

	index := getUserInteger()

	stream, err := c.GetSequence(
		context.Background(),
		&fibonacci.Number{Value: int32(index)},
	)
	if err != nil {
		log.Fatalf("call failed: %v", err)
	}
	i := 0
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		i += 1
		log.Printf("Fibonacci #%v: %v", i, r.Value)
	}
}

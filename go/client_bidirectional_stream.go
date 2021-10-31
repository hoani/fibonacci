// client_bidirectional_stream.go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	fibonacci "github.com/hoani/fibonacci/proto"
	"google.golang.org/grpc"
)

func getUserInteger() (int, error) {
	var indexStr string
	fmt.Print("enter the next index (or nothing to stop): ")
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

	stream, err := c.StreamSequence(context.Background())
	if err != nil {
		log.Fatalf("call failed: %v", err)
	}

	rxDone := make(chan struct{})
	go func() {
		index := 0
		for {
			r, err := stream.Recv()
			if err == io.EOF {
				close(rxDone)
				return
			}
			if err != nil {
				log.Fatalf("Receive error: %v", err)
			}
			index += 1
			log.Printf("Fibonacci #%v: %v", index, r.Value)
		}
	}()
	for {
		increment, err := getUserInteger()
		if err != nil {
			err := stream.CloseSend()
			if err != nil {
				log.Fatalf("error closing %v", err)
			}
			break
		}
		request := &fibonacci.Number{
			Value: int32(increment),
		}
		if err := stream.Send(request); err != nil {
			log.Fatalf("send failed: %v", err)
		}
		<-time.After(time.Second)
	}
	<-rxDone
}

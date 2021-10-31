// client_unary.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	request := &fibonacci.Number{Value: int32(index)}
	r, err := c.AtIndex(ctx, request)
	if err != nil {
		log.Fatalf("call failed: %v", err)
	}
	log.Printf("Fibonacci #%v: %v", index, r.Value)
}

// server.go
package main

import (
	"context"
	"io"
	"log"
	"net"

	fibonacci "github.com/hoani/fibonacci/proto"
	"google.golang.org/grpc"
)

func calculate(index int32) int32 {
	var a, b int32 = 0, 1
	for i := 0; i < int(index); i++ {
		a, b = b, a+b
	}
	return a
}

type server struct {
	fibonacci.UnimplementedFibonacciServer
}

func (s *server) AtIndex(
	ctx context.Context,
	in *fibonacci.Number,
) (*fibonacci.Number, error) {
	return &fibonacci.Number{Value: calculate(in.Value)}, nil
}

func (s *server) GetSequence(
	in *fibonacci.Number,
	stream fibonacci.Fibonacci_GetSequenceServer,
) error {
	var i int32 = 0
	for i = 0; i < in.Value; i++ {
		result := &fibonacci.Number{Value: calculate(i + 1)}
		if err := stream.Send(result); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) SumIndicies(
	stream fibonacci.Fibonacci_SumIndiciesServer,
) error {
	result := &fibonacci.Number{Value: 0}
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(result)
		}
		if err != nil {
			return err
		}
		result.Value += calculate(request.Value)
	}
}

func (s *server) StreamSequence(
	stream fibonacci.Fibonacci_StreamSequenceServer,
) error {
	var index int32 = 0
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		for i := 0; i < int(request.Value); i++ {
			index++
			result := &fibonacci.Number{
				Value: calculate(index),
			}
			if err := stream.Send(result); err != nil {
				return err
			}
		}
	}
}

func main() {
	l, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fibonacci.RegisterFibonacciServer(s, &server{})
	log.Printf("server listening at %v", l.Addr())
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

# client_unary.py
import grpc
import fibonacci_pb2
import fibonacci_pb2_grpc

if __name__ == '__main__':
    with grpc.insecure_channel('localhost:1337') as channel:
        stub = fibonacci_pb2_grpc.FibonacciStub(channel)
        value = int(input("enter an index integer: "))
        response = stub.AtIndex(fibonacci_pb2.Number(value=value))
    print(f"Fibonacci #{value}: {response.value}")

# client_server_stream.py
import grpc
import fibonacci_pb2
import fibonacci_pb2_grpc

if __name__ == '__main__':
    with grpc.insecure_channel('localhost:1337') as channel:
        stub = fibonacci_pb2_grpc.FibonacciStub(channel)
        value = int(input("enter an index integer: "))
        responses = stub.GetSequence(fibonacci_pb2.Number(value=value))
        index = 0
        for response in responses:
            index += 1
            print(f"Fibonacci #{index}: {response.value}")

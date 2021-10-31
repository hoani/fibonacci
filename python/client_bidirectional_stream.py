# client_bidirectional_stream.py
import grpc
from time import sleep
import fibonacci_pb2
import fibonacci_pb2_grpc

def nextIncrement():
    try:
        while(True):
            sleep(1)
            value = int(input("enter the next increment "+
                "(or nothing to stop): "))
            yield fibonacci_pb2.Number(value=value)
    except:
        print("done sending")

if __name__ == '__main__':
    with grpc.insecure_channel('localhost:1337') as channel:
        stub = fibonacci_pb2_grpc.FibonacciStub(channel)
        responses = stub.StreamSequence(nextIncrement())
        index = 0
        for response in responses:
            index += 1
            print(f"Fibonacci #{index}: {response.value}")

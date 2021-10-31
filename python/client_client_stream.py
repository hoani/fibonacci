# client_client_stream.py
import grpc
import fibonacci_pb2
import fibonacci_pb2_grpc

def sendIndicies():
    try:
        while(True):
            value = int(input("enter an index (or nothing to stop): "))
            yield fibonacci_pb2.Number(value=value)
    except:
        print("done sending")

if __name__ == '__main__':
    with grpc.insecure_channel('localhost:1337') as channel:
        stub = fibonacci_pb2_grpc.FibonacciStub(channel)
        response = stub.SumIndicies(sendIndicies())
    print(f"Sum: {response.value}")

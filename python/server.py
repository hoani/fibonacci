# server.py
from concurrent import futures

import grpc
import fibonacci_pb2
import fibonacci_pb2_grpc

class Fibonacci(fibonacci_pb2_grpc.FibonacciServicer):

    def calculate(self, index):
        a, b = 0, 1
        for _ in range(index):
            a, b = b, a + b
        return a

    def AtIndex(self, request : fibonacci_pb2.Number, context):
        result = self.calculate(request.value)
        return fibonacci_pb2.Number(value=result)

    def GetSequence(self, request, context):
        for index in range(request.value):
            result = self.calculate(index + 1)
            yield fibonacci_pb2.Number(value=result)

    def SumIndicies(self, requests, context):
        result = 0
        for request in requests:
            result += self.calculate(request.value)
        return fibonacci_pb2.Number(value=result)

    def StreamSequence(self, requests, context):
        index = 0
        for request in requests:
            for _ in range(request.value):
                index += 1
                result = self.calculate(index)
                yield fibonacci_pb2.Number(value=result)


def serve():
    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=10)
    )
    fibonacci_pb2_grpc.add_FibonacciServicer_to_server(
        Fibonacci(), server
    )
    server.add_insecure_port('[::]:1337')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()

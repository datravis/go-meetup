import grpc

import hello_pb2
import hello_pb2_grpc



channel = grpc.insecure_channel('localhost:9000')
stub = hello_pb2_grpc.HelloServiceStub(channel)

resp = stub.Hello(hello_pb2.HelloRequest(
    message="How's it going"
))
print("Received message %s" % (resp))
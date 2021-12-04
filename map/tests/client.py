from typing import Tuple
import grpc
from stubs import map_pb2 as pb2
from stubs import map_pb2_grpc as pb2_grpc

class MapClient:
    def __init__(self, host: str, port: str):
        self.channel = grpc.insecure_channel(f'{host}:{port}')
        self.stub = pb2_grpc.MapStub(self.channel)

    def resolve_coordination(self, latitude: str, longitude: str) -> Tuple[str, str, str]:
        response = self.stub.GetServerResponse(pb2.CoordinateRequest(latitude=latitude, longitude=longitude))
        print(response)
        return response.district, response.city, response.country
    
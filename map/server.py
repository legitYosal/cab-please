import grpc
from concurrent import futures
from stubs import map_pb2 as pb2
from stubs import map_pb2_grpc as pb2_grpc
from utils.nominatim import resolve_coordinates
import config

class MapService(pb2_grpc.MapServicer):
    def __init__(self, *args, **kwargs):
        pass
    
    def kebab_case(self, string: str) -> str:
        return string.replace(' ', '-').lower()

    def GetServerResponse(self, request, context):
        latitude = request.latitude
        longitude = request.longitude
        district, city, country = resolve_coordinates(latitude, longitude)
        return pb2.DistrictResponse(
            district=self.kebab_case(district), 
            city=self.kebab_case(city), 
            country=self.kebab_case(country)
        )

if __name__ == '__main__':
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_MapServicer_to_server(MapService(), server)
    server.add_insecure_port(f'[::]:{config.PORT}')
    server.start()
    print(f'Server started on :{config.PORT}.')
    server.wait_for_termination()
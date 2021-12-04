from typing import Tuple
import requests

def resolve_coordinates(latitude: str, longitude: str) -> Tuple[str, str, str]:
    res = requests.get(
        f'https://nominatim.openstreetmap.org/reverse?lat={latitude}&lon={longitude}&zoom=14&format=json&accept-language=en')
    if res.status_code != 200:
        print('nominatim not successful!')
    add = res.json()['address']
    return add['suburb'], add['city'] , add['country']

from typing import List
import redis
import json
from django.conf import settings

client = redis.Redis(
    host=settings.REDIS_HOST, 
    port=settings.REDIS_PORT, 
    db=settings.REDIS_DB
)

def set_surge_config(surge_config: List[List]) -> None:
    client.set(settings.SURGE_CONFIG_CACHE_KEY, json.dumps(surge_config))
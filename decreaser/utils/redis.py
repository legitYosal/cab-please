import redis
import config

client = redis.Redis(host=config.REDIS_HOST, port=6379, db=0)

def decrease_demand(demand_key: str):
    if client.exists(demand_key):
        print(client.get(demand_key))
        client.decr(demand_key)
        print(client.get(demand_key))
import os
from dotenv import load_dotenv

load_dotenv('.env')

REDIS_HOST = os.getenv('REDIS_HOST')
RABBITMQ_HOST = os.getenv('RABBITMQ_HOST')
print(RABBITMQ_HOST)
RABBITMQ_PORT = '5672'
DECREASE_QUEUE_NAME = 'decrease_demand_queue_name'
RABBITMQ_DELAYED_EXCHANGE_NAME = 'delayed_exchange_for_demand_decrease'
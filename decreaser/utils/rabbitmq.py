import pika
import config

connection = pika.BlockingConnection(pika.ConnectionParameters(
    host=config.RABBITMQ_HOST,
    port=config.RABBITMQ_PORT,
))
channel = connection.channel()

channel.exchange_declare(
    exchange=config.RABBITMQ_DELAYED_EXCHANGE_NAME, 
    exchange_type='x-delayed-message', durable=True,
    arguments={'x-delayed-type': 'direct'}    
)

channel.queue_declare(queue=config.DECREASE_QUEUE_NAME)

channel.queue_bind(exchange=config.RABBITMQ_DELAYED_EXCHANGE_NAME, queue=config.DECREASE_QUEUE_NAME)
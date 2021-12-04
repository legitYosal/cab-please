import pika
import config

connection = pika.BlockingConnection(pika.ConnectionParameters(
    host=config.RABBITMQ_HOST,
    port=config.RABBITMQ_PORT,
))
channel = connection.channel()
channel.queue_declare(queue=config.DecreaseDemandQueueName)
from utils.rabbitmq import channel, connection
import config
import pika


if __name__ == '__main__':
    channel.basic_publish(
        exchange=config.RABBITMQ_DELAYED_EXCHANGE_NAME,
        routing_key=config.DECREASE_QUEUE_NAME,
        properties=pika.BasicProperties(
            headers={'x-delay': 1000}
        ),
        body='Hello World!'
    )
    channel.close()
    connection.close()
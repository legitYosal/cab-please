import os
import sys
import config
from utils.rabbitmq import channel
from utils.redis import decrease_demand

def callback(ch, method, properties, body):
    print(" [x] Received %r" % body.decode())
    demand_key = body.decode()
    decrease_demand(demand_key)

if __name__ == '__main__':
    try:
        channel.basic_consume(queue=config.DecreaseDemandQueueName,
                        auto_ack=True,
                        on_message_callback=callback)
        print(' [*] Waiting for messages. To exit press CTRL+C')
        channel.start_consuming()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
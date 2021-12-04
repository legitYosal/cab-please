import os
from dotenv import load_dotenv

load_dotenv('.env')
PORT = os.getenv('SERVER_PORT')
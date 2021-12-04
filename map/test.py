from tests.client import MapClient
import config

def feature_test_resolve_coordination():
    client = MapClient('localhost', config.PORT)
    c = client.resolve_coordination('35.680725', '51.427646') 
    # this is district 12 tehran iran


if __name__ == '__main__':
    feature_test_resolve_coordination()
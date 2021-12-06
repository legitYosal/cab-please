import os
import django
import json

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'surgeapi.settings')
django.setup()

from thresholds.models import Threshold, DefaultLoaded
from thresholds.views import BaseThresholdAPI

def load_if_not_loaded():
    if not DefaultLoaded.objects.all().exists():
        with open('surge.config.default.json') as f:
            data = json.load(f)
            for threshold in data:
                Threshold.objects.create(**threshold)
                print('loaded:',threshold)
        BaseThresholdAPI.update_on_redis()        

if __name__ == '__main__':
    load_if_not_loaded()
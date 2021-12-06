import os
import django
import json

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'surgeapi.settings')
django.setup()

from thresholds.models import Threshold, DefaultLoaded

def load_if_not_loaded():
    if not DefaultLoaded.objects.all().exists():
        with open('surge.config.default.json') as f:
            data = json.load(f)
            for threshold in data:
                Threshold.objects.create(**threshold)
                print('loaded:',threshold)

if __name__ == '__main__':
    load_if_not_loaded()
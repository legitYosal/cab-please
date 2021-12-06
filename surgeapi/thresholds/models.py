from django.db import models

class DefaultLoaded(models.Model):
    """
        This is an unusal database table to only hold the fact that
        we had loaded the default thresholds or not.

        if len(table) > 0: we had loaded defaults
        else: we did not load defaults
    """
    loaded = models.BooleanField(default=True)

class Threshold(models.Model):
    request_threshold = models.IntegerField(unique=True)
    price_coefficient = models.FloatField()

    def __str__(self, ):
        return f'Thresh: {self.request_threshold} => {self.price_coefficient}'

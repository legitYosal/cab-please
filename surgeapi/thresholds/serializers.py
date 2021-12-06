from django.db.models import fields
from .models import Threshold
from rest_framework import serializers

class ThresholdSerializer(serializers.ModelSerializer):
    class Meta:
        model = Threshold
        fields = ('__all__')
        read_only_fields = ('id',)
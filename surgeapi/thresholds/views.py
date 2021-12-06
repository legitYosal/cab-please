from django.shortcuts import render
from rest_framework.generics import GenericAPIView
from rest_framework.mixins import (
    ListModelMixin, CreateModelMixin,
    DestroyModelMixin, UpdateModelMixin
)
from .permissions import IsAdmin
from .models import Threshold
from .serializers import ThresholdSerializer
from .utils.redis import set_surge_config

class BaseThresholdAPI(GenericAPIView):

    queryset = Threshold.objects.all()
    serializer_class = ThresholdSerializer
    lookup_field = 'id'

    permission_classes = [IsAdmin, ]

    @classmethod
    def update_on_redis(cls, ):
        """
            we just get all the thresholds and cache them
            into the redis
        """
        data = [[
            threshold.request_threshold,
            threshold.price_coefficient,
        ] for threshold in Threshold.objects.order_by('request_threshold')]
        set_surge_config(data)

class UpdateDeleteThresholdsAPI(BaseThresholdAPI, DestroyModelMixin,
                    UpdateModelMixin):

    def perform_update(self, serializer):
        super().perform_update(serializer)
        self.update_on_redis()

    def perform_destroy(self, instance):
        super().perform_destroy(instance)
        self.update_on_redis()

    def delete(self, request, *args, **kwargs):
        return self.destroy(request, *args, **kwargs)

    def put(self, request, *args, **kwargs):
        return self.partial_update(request, *args, **kwargs)

class ListCreateThresholdsAPI(BaseThresholdAPI, ListModelMixin,
                    CreateModelMixin):

    def perform_create(self, serializer):
        super().perform_create(serializer)
        self.update_on_redis()


    def get(self, request, *args, **kwargs):
        return self.list(request, *args, **kwargs)
    
    def post(self, request, *args, **kwargs):
        return self.create(request, *args, **kwargs)

from django.urls import path
from .views import ListCreateThresholdsAPI, UpdateDeleteThresholdsAPI

urlpatterns = [
    path('', ListCreateThresholdsAPI.as_view(), name='list-create-thresholds'),
    path('<int:id>/', UpdateDeleteThresholdsAPI.as_view(), name='update-delete-thresholds'),
]
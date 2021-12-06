from rest_framework.permissions import BasePermission

class IsAdmin(BasePermission):
    def has_permission(self, request, view):
        if request.user and request.user.is_anonymous:
            return False
        return request.user and request.user.is_admin
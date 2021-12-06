from rest_framework_simplejwt.models import TokenUser
from django.utils.functional import cached_property

class CustomizedTokenUser(TokenUser):
    @cached_property
    def is_admin(self):
        return self.token.get('is_admin', False)

    def __init__(self, token_payload):
        super().__init__(token_payload)
from .models import Product
from django.conf import settings
from rest_framework import serializers

import logging
logger = logging.getLogger(__name__)

import grpc
import discount_pb2
import discount_pb2_grpc

class ProductSerializer(serializers.ModelSerializer):
    discount = serializers.SerializerMethodField('get_discounts')
    def get_discounts(self, obj):
        userID = self.context.get('userID')
        logger.info('Request Discount user id %s product %s' % (userID, obj.title))
        try:
          channel = grpc.insecure_channel(settings.GRPC_URL)
          stub = discount_pb2_grpc.DiscountServiceStub(channel)
          req = discount_pb2.DiscountResquest(userID=int(userID), productID=obj.id)
          resp = stub.Get(req, timeout=0.5)
          return {
            "value_in_cents": resp.value_in_cents,
            "pct": resp.pct
          }
        except Exception as e:
          logger.debug(e)
          return {
            "value_in_cents": obj.price_in_cents,
            "pct": 0
          }

    class Meta:
        model = Product
        fields = [
            'id',
            'price_in_cents',
            'title',
            'description',
            'discount'
        ]

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from .models import Product
from .serializer import ProductSerializer

import logging
logger = logging.getLogger(__name__)

class ListProducts(APIView):
    def get(self, request):
        userID = request.GET.get('X-USER-ID', '')
        if not userID:
          return Response(status=status.HTTP_400_BAD_REQUEST)
        products = Product.objects.all()
        serializer = ProductSerializer(products, many=True, context={'userID': userID})
        return Response(serializer.data)
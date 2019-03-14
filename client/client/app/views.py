from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from .models import Product, User
from .serializer import ProductSerializer

import logging
logger = logging.getLogger(__name__)

class ListProducts(APIView):
    def get(self, request):
        userID = request.GET.get('X-USER-ID', None)
        if not userID:
          return Response(status=status.HTTP_400_BAD_REQUEST)

        try:
            user =  User.objects.get(id=userID)
        except User.DoesNotExist:
            return Response(status=status.HTTP_404_NOT_FOUND)
        products = Product.objects.all()
        serializer = ProductSerializer(products, many=True, context={'userID': userID})
        return Response(serializer.data)

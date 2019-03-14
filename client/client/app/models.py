from django.db import models

# Create your models here.
class User(models.Model):
    first_name = models.CharField(max_length=30)
    last_name = models.CharField(max_length=30)
    date_of_birth = models.DateField()

class Product(models.Model):
    price_in_cents = models.IntegerField()
    title = models.CharField(max_length=100)
    description = models.TextField()
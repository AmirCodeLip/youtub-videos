from django.db import models
from django import forms



class Status(models.TextChoices):
    DRAFT = 'DF', 'Draft'
    PUBLISHED = 'PB', 'Published'
    ARCHIVED = 'AR', 'Archived'

class Article(models.Model):
    status = models.CharField(
        max_length=2,
        choices=Status.choices,
        default=Status.DRAFT
    )
from django import forms
from django.db import models
from .models import Article


class Status(models.TextChoices):
    DRAFT = "DF", "Draft"
    PUBLISHED = "PB", "Published"

    @classmethod
    def published_states(cls):
        return [cls.PUBLISHED]


class ContactForm(forms.Form):
    name = forms.CharField(widget=forms.TextInput(attrs={"class": "special"}))
    email = forms.EmailField(widget=forms.EmailInput(attrs={"size": "40"}))
    message = forms.CharField(widget=forms.Textarea(attrs={"rows": 4}))
    status = forms.ChoiceField(widget=forms.RadioSelect, choices=Status.choices)


class ArticleForm(forms.ModelForm):
    class Meta:
        model = Article
        fields = ['status']
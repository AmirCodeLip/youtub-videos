from django.views import View
from django.shortcuts import render
from .forms import ContactForm, ArticleForm
from django.db import models


class CreateTask(View):

    def get(self, request):

        # form = TaskForm()
        form = ContactForm(initial={"name": "John Doe"})

        # print(
        #     form.fields["status"].choices
        # )  # Should output [('PEN', 'Pending'), ('APP', 'Approved'), ('REJ', 'Rejected')]

        return render(request, "CreateTask.html", {"form": form})


class Article(View):

    def get(self, request):
        
        form = ArticleForm()
        return render(request, "CreateTask.html", {"form": form})

    def post(self, request):

        form = ArticleForm(request.POST)
        return render(request, "CreateTask.html", {"form": form})

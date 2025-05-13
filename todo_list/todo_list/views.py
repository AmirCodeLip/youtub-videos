from django.views import View
from django.shortcuts import render


class CreateTodo(View):

    def get(self, request):

        return render(request, "create_todo.html")

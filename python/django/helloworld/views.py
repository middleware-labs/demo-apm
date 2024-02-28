import logging
from django.http import HttpResponse

def index(request):
    return HttpResponse("Hello, world!")

def ok(request):
    return HttpResponse("Hello, ok", status=404)

def run(request):
    logging.warning("run: warning log sample")
    return HttpResponse("Hello, run!")

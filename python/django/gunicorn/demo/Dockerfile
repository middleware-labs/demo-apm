FROM python:3.11-slim-buster
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN pip install -r requirements.txt
RUN middleware-bootstrap -a install
ENV DJANGO_SETTINGS_MODULE=demo.settings

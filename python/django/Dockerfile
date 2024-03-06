# Copyright The OpenTelemetry Authors
# SPDX-License-Identifier: Apache-2.0

FROM python:3.11-slim-bullseye as base

#
# Fetch requirements
#
FROM base as builder
RUN apt-get -qq update \
    && apt-get install -y --no-install-recommends g++ \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app/
COPY ./requirements.txt ./

RUN pip install --upgrade pip
RUN pip install --prefix="/reqs" -r requirements.txt
#
# Runtime
#
FROM base as runtime
WORKDIR /usr/src/app/
COPY --from=builder /reqs /usr/local
COPY ./ ./

ENV OTEL_PYTHON_TRACER_PROVIDER=sdk_tracer_provider
RUN MIDDLEWARE_CONFIG_FILE=./middleware.ini python3 manage.py migrate 
RUN MIDDLEWARE_CONFIG_FILE=./middleware.ini python3 manage.py createsuperuser --username admin --email admin@mail.com
RUN middleware-bootstrap -a install 
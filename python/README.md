# Python APM Guide
You can follow our [documentation](https://docs.middleware.io/docs/apm-configuration/python/python-apm-setup) to setup APM for your Python application.

[![PyPI - Version](https://img.shields.io/pypi/v/middleware-apm)](https://pypi.org/project/middleware-apm/)


| Traces | Metrics | Profiling | Logs (App/Custom) |
|--------|---------|-----------|-------------------|
|   Yes  |  Yes    |    Yes    | Yes/Yes           |

## Prerequisites
Ensure that you have the Middleware Host Agent installed to view Python demo data on your dashboard.

---------------------

## Installing the Package
Run the following commands in your terminal:
### Step 1: Install Middleware APM package
```
pip install middleware-apm
```
Step 2: Auto-install Required Packages
```
middleware-bootstrap -a install
```
The first command installs the Middleware APM package, and the second command automatically installs the instrumentation libraries for the packages in your active site-packages folder, if applicable.

## Your Sample Code
By using all the MW's APM functionalities like: Distributed-tracing, Logs, Metrics and Profiling, your code will look like this:
```python
import logging

from mw_tracker import MwTracker
tracker=MwTracker(
    access_token="{YOUR-ACCESS_TOKEN}"
)

tracker.collect_metrics()
tracker.collect_logs()
tracker.collect_profiling()

logging.info("Hello World!", extra={'key': 'value'})
```

## Run Your Application
To run your application, use the following command:
```
export PROTOCOL_BUFFERS_PYTHON_IMPLEMENTATION=python
middleware-instrument \
--exporter_otlp_endpoint http://localhost:9319 \
--resource_attributes=project.name={APM-PROJECT-NAME},mw.app.lang=python,runtime.metrics.python=true \
--service_name {APM-SERVICE-NAME} \
python3 app.py
```
#### Note: You need to replace <strong>\{APM-PROJECT-NAME\}</strong> and <strong>\{APM-SERVICE-NAME\}</strong> with your project name and service name respectively.

---------------------------------

## Troubleshooting Demo
If you face any protoc specific errors, Try setting ...
```
PROTOCOL_BUFFERS_PYTHON_IMPLEMENTATION=python
```

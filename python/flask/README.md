# Python APM Guide
You can follow our [documentation](https://docs.middleware.io/docs/apm-configuration/python/python-apm-setup) to setup APM for your Python application.

[![PyPI - Version](https://img.shields.io/pypi/v/middleware-apm)](https://pypi.org/project/middleware-apm/)


|  Traces  |  Metrics  |  Profiling  |  Logs (App/Custom)  |
|:--------:|:---------:|:-----------:|:-------------------:|
|   Yes    |    Yes    |     Yes     |       Yes/Yes       |

## Prerequisites
Ensure that you have the Middleware Host Agent installed to view Python demo data on your dashboard.

---------------------

## Installing the Package
Run the following commands in your terminal:
### Step 1: Install Middleware APM package
```shell
pip install middleware-apm
```

## Your Sample Code
By using all the MW's APM functionalities like: Distributed-tracing, Logs, Metrics and Profiling, your code will look like this:
```python
import logging

from middleware import MwTracker
tracker=MwTracker()

logging.warning("Sample Warning Log.")
logging.error("Sample Error Log.", extra={'tester': 'Alex'})
```
## Setup middleware.ini File
Setup your .ini file, based on below features that you want to observe in your project.
```ini
# ---------------------------------------------------------------------------
# This file contains settings for the Middleware Python-APM Agent.
# Here are the settings that are common to all environments.
# ---------------------------------------------------------------------------

[middleware.common]

# The name of your application as service-name, as it will appear in the UI to filter out your data.
service_name = Python-APM-Service

# This Token binds the Python Agent's data and profiling data to your account.
access_token = {YOUR-ACCESS-TOKEN}

# The service name, where Middleware Agent is running, in case of K8s.
;mw_agent_service = mw-service.mw-agent-ns.svc.cluster.local

# Toggle to enable/disable distributed traces for your application.
collect_traces = true

# Toggle to enable/disable the collection of metrics for your application.
collect_metrics = true

# Toggle to enable/disable the collection of logs for your application.
collect_logs = true

# Toggle to enable/disable the collection of profiling data for your application.
collect_profiling = true

# ---------------------------------------------------------------------------
```
#### Note: You need to replace <strong>\{YOUR-ACCESS-TOKEN\}</strong> with your APM's Access Token.

## Run Your Application
To run your application, use the following command:
```shell
MIDDLEWARE_CONFIG_FILE=./middleware.ini middleware-apm run python app.py
```

---------------------------------

## Troubleshooting Demo
If you face any protoc specific errors, Try setting ...
```
export PROTOCOL_BUFFERS_PYTHON_IMPLEMENTATION=python
```
---------------------------------
## Run on Docker
1. Build: `docker build -t demo-python .`
2. Run: `docker run demo-python`
3. Debug: `docker run -it demo-python sh`
# Python APM Setup
You can follow our [documentation](https://docs.middleware.io/docs/apm-configuration/python/python-apm-setup) to setup APM for your Python application.

## Prerequisites

* To monitor APM data on dashboard, Middleware Host agent needs to be installed.
* You can refer [this demo project](https://github.com/middleware-labs/demo-apm/tree/master/python) to refer use cases
  of APM.
* This demo requires the Flask library, and you can install it from requirements.txt using `pip install -r requirements.txt` or by running `pip install flask`.

## Getting Started
### Step 1: Install MW's APM Package
Run the following command in your terminal:
```
pip install middleware-apm
```

### Step 2: Auto-install Required Packages
The `middleware-bootstrap -a install` command reads through the list of packages installed in your active site-packages folder, and installs the corresponding instrumentation libraries for these packages, if applicable.
```
middleware-bootstrap -a install
```

### Step 3: Import the Tracker
Add the following lines at the very start of your project:
```python
from mw_tracker import MwTracker
tracker=MwTracker(
    access_token="{YOUR-ACCESS_TOKEN}"
)
```

## Enable MW's APM functionalities

### 1. Enable Distributed Tracing
Distributed-tracing is a default functionality of MW's APM. This will be auto-enabled when you execute the command mentioned in **Run Your Application** section.

### 2. Enable Metrics
You can enable metrics by writing the following line just after the tracker initialization.
```python
tracker.collect_metrics()
```

### 3. Enable Logs
Similar to metrics, you can enable logs by writing the following line.
```python
tracker.collect_logs()
```

### 4. Enable Profiling
For the Profiling feature, you need to write the following line.
```python
tracker.collect_profiling()
```

## Your Code will look like this
After adding all the above steps, your code will look like this:
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
To run your application, you need to execute the following command:
```
middleware-instrument \
--exporter_otlp_endpoint http://localhost:9319 \
--resource_attributes=project.name={APM-PROJECT-NAME},mw.app.lang=python,runtime.metrics.python=true \
--service_name {APM-SERVICE-NAME} \
python3 app.py
```
#### Note: You need to replace <strong>\{APM-PROJECT-NAME\}</strong> and <strong>\{APM-SERVICE-NAME\}</strong> with your project name and service name respectively.

## For APM inside Kubernetes
If you are using APM in a Kubernetes cluster make sure to follow these two steps:

### Step 1: Find your Middleware Service namespace
For older setup, your "mw-service" can be inside `mw-agent-ns-{FIRST-5-LETTERS-OF-API-KEY}` namespace
For newer setup, we simplified the namespace name to `mw-agent-ns`

### Step 2: Set this ENV variable in your application deployment YAML
```python
MW_AGENT_SERVICE=mw-service.NAMESPACE.svc.cluster.local
```
#### Note: Please replace "NAMESPACE" with the correct value that you found from Step 1.

## Troubleshoot:
If you receive a warning similar to the following:
```
WARNING: The scripts middleware-bootstrap and middleware-instrument are installed in '/home/.../.local/bin' which is not on PATH.

Consider adding this directory to PATH or, if you prefer to suppress this warning, use --no-warn-script-location
```

You can add the binary to your path
In linux, you can add this with
```
export PATH=$PATH:/home/.../.local/bin
```

## Error Handling :
If you want to record exception in traces then, you can use `tracker.record_error(e)` method.

```python
randomList = ['a', 0, 2]

for entry in randomList:
    try:
         r = 1/int(entry)
         break
    except Exception as e:
         tracker.record_error(e)
```

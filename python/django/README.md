# Python APM Guide
You can follow our [documentation](https://docs.middleware.io/docs/apm-configuration/python/python-apm-setup) to setup APM for your Django application.

[![PyPI - Version](https://img.shields.io/pypi/v/middleware-apm)](https://pypi.org/project/middleware-apm/)


|  Traces  |  Metrics  |  Profiling  |  Logs (App/Custom)  |
|:--------:|:---------:|:-----------:|:-------------------:|
|   Yes    |    Yes    |     Yes     |       Yes/Yes       |

## Prerequisites
Ensure that you have the Middleware Host Agent installed to view Python demo data on your dashboard.

---------------------
## Setup `middleware.ini` File
Setup your `middleware.ini` file, based on below features that you want to observe in your project. Place the file at the root of your project.
```ini
# ---------------------------------------------------------------------------
# This file contains settings for the Middleware Python-APM Agent.
# Here are the settings that are common to all environments.
# ---------------------------------------------------------------------------

[middleware.common]

# The name of your application as service-name, as it will appear in the UI to filter out your data.
service_name = Django-Python-APM-Service

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

## Installing Pre-Requirements

You need install the pre-requirements for run this Hello World example.

Update repositories of available packages to install, with
the following command:

```shell
sudo apt update
```

Install necessary minimum dependencies, with the following command:

```shell
sudo apt install python3-dev python3-pip python3-virtualenv sqlitebrowser
```

For run this example need to install Django
framework execute the follow command:

```python
virtualenv django_instrumentation
source django_instrumentation/bin/activate
pip install -r requirements.txt
```

And later followed by:

```python
export MIDDLEWARE_CONFIG_FILE=./middleware.ini
python3 manage.py migrate
```


For use the Django Admin Interface, it's needed to create a superuser 
for management, with the following command:

```python
python3 manage.py createsuperuser --username admin --email admin@mail.com
```
## Run Your Application
After performing all the above steps, now you can uncomment below lines, from manage.py file:
```python
from middleware import MwTracker
tracker=MwTracker()

def main():
    # -- existing code --
    tracker.django_instrument()
    # -- existing code --
```
Now, you can go ahead to run your application, by using:
```python
DJANGO_SETTINGS_MODULE='helloworld.settings' middleware-apm run python3 manage.py runserver
```

If you are running this application in a serverless setup consider adding `MW_API_KEY` and `MW_TARGET` as mentioned in the command below :
```python
MW_API_KEY=********** MW_TARGET=https://*****.middleware.io:443 DJANGO_SETTINGS_MODULE='helloworld.settings' middleware-apm run python3 manage.py runserver
```


#### Note: If `middleware.ini` isn't in your project's root, set `MIDDLEWARE_CONFIG_FILE=./path/to/middleware.ini` along with the above run command.
Then, you can open the URL http://127.0.0.1:8000/ in your web browser.

Also you can open in your web browser the URL http://127.0.0.1:8000/admin for access to 
the *Django Admin Interface*.

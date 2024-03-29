# Node.js APM Setup

## Prerequisites

* To monitor APM data on dashboard, Middleware Host agent needs to be installed.
* You can refer [this demo project](https://github.com/middleware-labs/demo-apm/tree/master/nodejs) to refer use cases of APM.

--------------------

## Step 1 : Install NPM package

Run this in your terminal
```
npm install @middleware.io/node-apm --save
```

### Troubleshoot

If your infrastructure is missing dependencies like g++, make, etc.
OR your node-gyp build fails, try ...

```
sudo apt-get build-dep build-essential
sudo apt-get install gcc
sudo apt-get install g++
sudo apt-get install make
```


## Step 2 : Prepend APM script

Add these lines given below at the very start of your project.

```
const tracker = require('@middleware.io/node-apm');
tracker.track({
    projectName: "{APM-PROJECT-NAME}",
    serviceName: "{APM-SERVICE-NAME}",
});
```
---------------------

## Collect Node.js specific metrics

The metrics collection starts automatically as soon as you complete `Step 2`


## Distributed Tracing

All your APIs are auto-instrumented as soon as you complete `Step 2`


## Add custom logs

```
tracker.info('Info sample');
tracker.warn('Warning sample');
tracker.debug('Debugging Sample');
tracker.error('Error Sample');
```

If you want to add stack trace along with the error log, you can follow the snippet as given below

```
tracker.error(new Error('Error sample with stack trace'));
```
## Note for APM inside Kubernetes

If you are using APM in a Kubernetes cluster make sure to follow these 2 steps:

### Step 1 : Find your Middleware Service namespace
For older setup, your "mw-service" can be inside "mw-agent-ns-{FIRST-5-LETTERS-OF-API-KEY}" namespace

For newer setup, we simplified the namespace name to "mw-agent-ns"

### Step 2 : Set this ENV variable in your application deployment YAML
```
MW_AGENT_SERVICE=mw-service.NAMESPACE.svc.cluster.local
```
Please replace "NAMESPACE" with the correct value that you found from Step 1.
-------------------

## Error Handling :

If you want to record exception in traces then you can use track.errorRecord(error) method.

```
 app.get('/error', function (req, res) {
    try{
        throw new Error('oh error!');
    }catch (e) {
       track.errorRecord(e)
    }
    res.status(500).send("wrong");
});
 
```

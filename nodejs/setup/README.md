# Node.js APM Setup

## Prequisites :

* To monitor APM data on dashboard, Middleware Host agent needs to be installed.
* You can refer [this demo project](https://github.com/middleware-labs/demo-apm/tree/master/nodejs) to refer use cases of APM.

--------------------

## Step 1 : Install NPM package

Run this in your terminal
```
npm install @middleware.io/node-apm --save
```

## Step 2 : Prepend APM script

Add these lines given below at the very start of your project.

```
const tracker = require('@middleware.io/node-apm');
tracker.track({
    projectName: "your-project-name",
    serviceName: "your-service-name",
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

## Note :

If you are using APM in a Kubernetes cluster, Make sure to pass this ENV variable:

```
MW_AGENT_SERVICE=mw-service.mw-agent-ns-{FIRST-5-LETTERS-OF-API-KEY}.svc.cluster.local
```

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

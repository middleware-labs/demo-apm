# Node.js APM Demo

---------------------
[![npm](https://img.shields.io/npm/v/%40middleware.io%2Fnode-apm)](https://www.npmjs.com/package/@middleware.io/node-apm)


| Traces | Metrics | Profiling | Logs (App/Custom) |
|--------|---------|-----------|-------------------|
|   Yes  | No      |    Yes    | No/Yes            |

# On a Linux Machine

## Prequisites

* If you are expecting nodejs demo data on your dashboard, make sure you have our Host Agent installed.

## Log Collection
```
node features/log.js
```

## Distributed Tracing
```
node features/trace.js
```

## Node.js Specific Metrics
```
node features/metric.js
```

## Complete Example
```
node app.js
```

## Complete Example + Database Monitoring (MySQL)
```
node dbdemo/app.js
```

---------------------------------

## Troubleshooting Demo
* If your infrastructure is missing dependencies like g++, make, etc.

    OR your node-gyp build fails, try ...
    ```
    sudo apt-get build-dep build-essential
    sudo apt-get install gcc
    sudo apt-get install g++
    sudo apt-get install make
    ```

* If your application is running inside a container, you may need to pass an ENV variable as follows:

  `MW_AGENT_SERVICE=172.17.0.1`
  
  The ENV suggests the address of your local machine where host agent is installed, for some machines this value might be "host.docker.internal"

---------------------

# On a Kubernetes Cluster

## Notes


* If you do not have a Kubernetes cluster, you can test this by running a minikube cluster in your machine.
## Prequisites

* If you are expecting nodejs demo data on your dashboard, make sure you have our Kubernetes Agent installed.

* To run Kubernetes APM demo, you may have to keep your Middleware API Key handy.

* You will need bash & curl tools to run the scripts given below.

## Log Collection

* Copy the content of `kubedemo/log.yaml` to your local machine

* Replace this placeholders in your local copy

| Placeholder      | Sample Value | Description     |
| :---             |    :----:    |          ---: |
| TARGET_MW_AGENT_SERVICE        | mw-service.mw-agent-ns-uxerb.svc.cluster.local         | Use the sample value replacing  "uxerb" with first 5 letters of your API Key      |

* then run the command given below

```
kubectl apply -f `PATH TO log.yaml`
```

If you want to customize the example, you can build a docker image from the Dockerfile provided here. Then replace the image name in your YAML file.

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
 



# Java APM Setup

## Prequisites :

* To monitor APM data on dashboard, Middleware Host agent needs to be installed.
* You can refer [this demo project](https://github.com/middleware-labs/demo-apm/tree/master/java) to refer use cases of
  APM.

--------------------

## Distributed Tracing

For recording the traces you will need to download JAR files given below.

[middleware-javaagent.jar](https://install.middleware.io/jars/middleware-javaagent.jar)

And then run your project with command given below

```
java -javaagent:/PATH/TO/middleware-javaagent.jar \
    -Dotel.service.name=test-service \
    -Dotel.resource.attributes=project.name=test-project \
    -jar <YOUR_APP>.jar
```

## Add custom logs

Add this dependency in pom.xml

```
<dependency>
  <groupId>io.github.middleware-labs</groupId>
  <artifactId>agent-apm-java</artifactId>
  <version>0.0.7</version>
</dependency>
```

### Note:
Refer the latest APM version from the link given below

https://central.sonatype.dev/artifact/io.github.middleware-labs/agent-apm-java/0.0.7/versions


Then run ...

```
mvn install
```

Import logger package

```
import io.github.middlewarelabs.agentapmjava.Logger;
```

Use these functions for logging with different severity levels

```
Logger.info("info message");
Logger.debug("debug message");
Logger.warn("warn message");
Logger.error("error message");
```

## Note

If you are using APM in a Kubernetes cluster make sure to pass ENV variable:

```
MW_AGENT_SERVICE=mw-service.mw-agent-ns-{FIRST-5-LETTERS-OF-API-KEY}.svc.cluster.local
```

## Error Handling :

If you want to record exception in traces then you can use Logger.recordError(e) method.

```
import io.github.middlewarelabs.agentapmjava.Logger;
        
try {
       int[] myNumbers = {1, 2, 3};
       System.out.println(myNumbers[10]);
     } catch (Throwable  e) {
       Logger.recordError(e);
       System.out.println("Something went wrong.");
     }
 

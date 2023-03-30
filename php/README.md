# PHP APM Setup

## Prerequisites

* To monitor APM data on dashboard, [Middleware Host-agent](https://docs.middleware.io/docs/getting-started) needs to be installed, You can refer [this demo project](https://github.com/middleware-labs/demo-apm/tree/master/php) to refer use cases of APM.
* PHP requires at least PHP 8+ and a PHP-Extension to run this agent.

--------------------

## Guide

### Initial Setup:

Before installing this agent, you need to install PHP-Extension(named otel_instrumentation) to run this agent. You can follow below steps to install & enable it:
* Run `sudo pecl install channel://pecl.php.net/opentelemetry-1.0.0beta2`.
* Then, Add the extension to your `php.ini` file like: `extension=otel_instrumentation.so`.
* And verify that the extension is installed and enabled using: `php -m | grep  otel_instrumentation`.

#### Troubleshoot:-
If you are facing `ERROR: 'phpize' failed` while installing PHP-Extension, then you need to run follow cmd:
  ```
  sudo apt-get update
  sudo apt-get install php8.1-dev
  sudo apt-get update
  sudo pecl channel-update pecl.php.net
  ```

### Step 1: Install APM-PHP package

Run below command in your terminal to install Middleware's APM-PHP package.
```
composer require middleware/agent-apm-php
```

### Step 2: Prepend APM script

Add these lines given below at the very start of your project.

```
require 'vendor/autoload.php';
use Middleware\AgentApmPhp\MwApmCollector;
```

### Step 3: Use APM Collector & Start the Tracing-scope

By using the APM Collector, you will start tracing-scope before your code, Also you need to register your hooks along with initial declaration. 

In each hook, you need to define your Classes & Functions name, so whenever they run, agent will track them auto.

```
$mwCollector = new MwApmCollector('<PROJECT-NAME>', '<SERVICE-NAME>');
$mwCollector->preTracing();

$mwCollector->registerHook('<CLASS-NAME-1>', '<FUNCTION-NAME-1>', [
    'custom.attr1' => 'value1',
    'custom.attr2' => 'value2',
]);
$mwCollector->registerHook('<CLASS-NAME-2>', '<FUNCTION-NAME-2>');

```

### Step 4 : End the Tracing-scope

After your code-flow, you need to end the tracing scope, so that agent can send the data to Middleware's APM dashboard.

```
$mwCollector->postTracing();
```

### Final code snippet will be:

```
<?php
require 'vendor/autoload.php';

use Middleware\AgentApmPhp\MwApmCollector;

$mwCollector = new MwApmCollector('<PROJECT-NAME>', '<SERVICE-NAME>');
$mwCollector->preTracing();
$mwCollector->registerHook('<CLASS-NAME-1>', '<FUNCTION-NAME-1>', [
    'custom.attr1' => 'value1',
    'custom.attr2' => 'value2',
]);
$mwCollector->registerHook('<CLASS-NAME-2>', '<FUNCTION-NAME-2>');

// Your code goes here

$mwCollector->postTracing();
```

### Sample Code:
```
<?php
require 'vendor/autoload.php';

use Middleware\AgentApmPhp\MwApmCollector;

$mwCollector = new MwApmCollector('DemoProject', 'PrintService');
$mwCollector->preTracing();
$mwCollector->registerHook('DemoClass', 'runCode', [
    'code.column' => '12',
    'net.host.name' => 'localhost',
    'db.name' => 'users',
    'custom.attr1' => 'value1',
]);
$mwCollector->registerHook('DoThings', 'printString');

class DoThings {
    public static function printString($str): void {
        // sleep(1);
        echo $str . PHP_EOL;
    }
}

class DemoClass {
    public static function runCode(): void {
        DoThings::printString('Hello World!');
    }
}

DemoClass::runCode();

$mwCollector->postTracing();
```

---------------------

## Note:

If you are using APM in a Kubernetes cluster, Make sure to pass this ENV variable:

```
MW_AGENT_SERVICE=mw-service.mw-agent-ns-{FIRST-5-LETTERS-OF-API-KEY}.svc.cluster.local
```
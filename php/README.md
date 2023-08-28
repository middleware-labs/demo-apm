# PHP APM Setup

## Prerequisites

* To monitor APM data on dashboard, Middleware Host agent needs to be installed, You can refer [this demo project](https://github.com/middleware-labs/demo-apm/tree/master/php) to refer use cases of APM.
* PHP requires at least PHP 8+ and a PHP-Extension to run this agent.

--------------------

## Guide

### Initial Setup:

Before installing this agent, you need to install PHP-Extension(named otel_instrumentation) to run this agent. You can follow below steps to install & enable it:
* Run `sudo pecl install channel://pecl.php.net/opentelemetry-1.0.0beta3`.
* Then, Add the extension to your `php.ini` file like: `extension=opentelemetry.so`.
* And verify that the extension is installed and enabled using: `php -m | grep  opentelemetry`.

#### Troubleshoot:-
While installing PHP-Extension:
* If you are facing `pecl:command not found`, then you need to run follow cmd:
  ```
  sudo apt-get update
  apt-get install php-pear php8.1-dev
  ```
* If you are facing any kind of broken dependencies issues like: `libpcre2-dev : Depends: libpcre2-8-0 / libpcre2-16-0 / libpcre2-32-0`, then you need to run follow cmd:
  ```
  sudo apt --fix-broken install
  ```
* If you are facing `ERROR: 'phpize' failed`, then you need to run follow cmd:
  ```
  sudo apt-get update
  sudo apt-get install php8.1-dev
  ```
  ```
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
use Middleware\AgentApmPhp\MwTracker;
```

### Step 3: Use APM Collector & Start the Tracing-scope

By using the APM Collector, you will start tracing-scope before your code, Also you need to register your hooks along with initial declaration. 

In each hook, you need to define your Classes & Functions name, so whenever they run, agent will track them auto.

```
$tracker = new MwTracker('<PROJECT-NAME>', '<SERVICE-NAME>');
$tracker->preTrack();

$tracker->registerHook('<CLASS-NAME-1>', '<FUNCTION-NAME-1>', [
    'custom.attr1' => 'value1',
    'custom.attr2' => 'value2',
]);
$tracker->registerHook('<CLASS-NAME-2>', '<FUNCTION-NAME-2>');

```

### Step 4 : End the Tracing-scope

After your code-flow, you need to end the tracing scope, so that agent can send the data to Middleware's APM dashboard.

```
$tracker->postTrack();
```

### Step 5 : To enable Logging feature

If you want to enable Logging feature along with tracing in your project, then you can use below code snippet.

  ```
   $tracker->warn("this is warning log.");
   $tracker->error("this is error log.");
   $tracker->info("this is info log.");
   $tracker->debug("this is debug log.");
   ```

### Final code snippet will be:

```
<?php
require 'vendor/autoload.php';
use Middleware\AgentApmPhp\MwTracker;

$tracker = new MwTracker('<PROJECT-NAME>', '<SERVICE-NAME>');
$tracker->preTrack();
$tracker->registerHook('<CLASS-NAME-1>', '<FUNCTION-NAME-1>', [
    'custom.attr1' => 'value1',
    'custom.attr2' => 'value2',
]);
$tracker->registerHook('<CLASS-NAME-2>', '<FUNCTION-NAME-2>');

$tracker->info("this is info log.");

// ----
// Your code goes here.
// ----

$tracker->postTrack();
```

### Sample Code:
```
<?php
require 'vendor/autoload.php';
use Middleware\AgentApmPhp\MwTracker;

$tracker = new MwTracker('DemoProject', 'PrintService');
$tracker->preTrack();
$tracker->registerHook('DemoClass', 'runCode', [
    'code.column' => '12',
    'net.host.name' => 'localhost',
    'db.name' => 'users',
    'custom.attr1' => 'value1',
]);
$tracker->registerHook('DoThings', 'printString');

$tracker->info("this is info log.");

class DoThings {
    public static function printString($str): void {
        // sleep(1);
        global $tracker;
        $tracker->warn("this is warning log, but from inner function.");
        
        echo $str . PHP_EOL;
    }
}

class DemoClass {
    public static function runCode(): void {
        DoThings::printString('Hello World!');
    }
}

DemoClass::runCode();

$tracker->postTrack();
```

---------------------

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

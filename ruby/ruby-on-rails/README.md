# README

This README would normally document whatever steps are necessary to get the
application up and running.

Things you may want to cover:

* Check Ruby version >= 3.0.2
```bash
ruby --version
```

* Prerequisites
    - To monitor APM data on dashboard, [Middleware Host agent](https://app.middleware.io/installation#infrastructures/ubuntu) needs to be installed.

* Dependencies 
    ```ruby
    gem 'opentelemetry-sdk'
    gem 'opentelemetry-exporter-otlp'
    gem 'opentelemetry-instrumentation-all'
    gem 'pyroscope'
    gem 'middleware_apm', '~> 1.0.0'
    ```

* Add this code at the initialization of your application
    ```ruby
    require 'middleware/ruby_gem'
    Middleware::RubyGem.init  
    ```
* Start Rails Server
    ```bash
    bin/rails server
    ```

Other Commands:

* Database migration
    ```bash
    bin/rails db:migrate
    ```
* List of available tasks
    ```bash
    bin/rails --tasks
    ```

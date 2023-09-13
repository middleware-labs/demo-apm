# Ruby on Rails

This is the sample ruby application to get the information about the inform the

### Setup Steps

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

### Other Commands

* Database migration
  ```bash
  bin/rails db:migrate
  ```
* List of available tasks
  ```bash
  bin/rails --tasks
  ```
* Run docker image with docker compose
  * **Note:** Update the environment variables in `docker-compose.yaml`
  ```bash
  docker compose up
  ```

* Run in Kubernetes
  * **Note:** Update the environment variables and value of image attribute in `ruby-on-rails.yaml` 
  ```bash
  kubectl apply -f ruby-on-rails.yaml
  ```

---

### Required environment variables

| Environment variable        | Description        | Sample values                                                                                                                                        | Is required |
|-----------------------------|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| OTEL_EXPORTER_OTLP_ENDPOINT | Exporter endpoint  | For Local: http://localhost:9320 <br/> For Docker: http://172.17.0.1:9320 <br/> For Kubernetes: http://mw-service.mw-agent-ns.svc.cluster.local:9320 | Yes         |
| OTEL_SERVICE_NAME           | Your Service Name  | sample service name                                                                                                                                  | No          |
| OTEL_RESOURCE_ATTRIBUTES    | Your Project Name  | project.name="sample project name"                                                                                                                   | No          |
| MW_API_KEY                  | Middleware API Key | {Your API Key}                                                                                                                                       | Yes         |

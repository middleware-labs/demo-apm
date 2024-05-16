# Dotnet Demo Project

This demo project is the sample configured example for dotnet application to monitor.

| Traces | Metrics | Profiling | Logs (App/Custom) |
|--------|---------|-----------|-------------------|
| Yes    | Yes     | Yes       | No/Yes            |

## Setup

- Follow the instruction for setup [.net](https://app.middleware.io/installation#apm/.net) application.

### Docker 

- Run docker image with docker compose
  - **Note:** Update the environment variables in `docker-compose.yaml` file.
  ```shell
  docker compose up
  ```

### Kubernetes
- Run docker image in kubernetes
  - **Note:** Update the environment variables and value of image attribute in `dotnet-app.yaml` file.
  ```shell
  kubectl apply -f dotnet-app.yaml
  ```
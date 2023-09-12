# Spring Boot REST Example

This demo project is for test java application with database integration.
This also contents the Dockerfile, docker-compose.yaml and yaml files to run in the kubernetes environment.

## Requirements

- Database Mysql
- java 17 +

## Description

- Set Environment variables to connect with database.

  | Environment Variable | Description             | Example value                        |
  |----------------------|-------------------------|--------------------------------------|
  | MYSQL_ROOT           | Root username for mysql | root                                 |
  | MYSQL_ROOT_PASSWORD  | Root password           | root                                 |
  | MYSQL_DATABASE       | Database name           | employeedb                           |
  | MYSQL_URL            | Database connection URL | jdbc:mysql://mysqldb:3306/employeedb |

- Update the environment variables and configuration of yaml files of docker and kubernetes.
  - Update the value of MW_API_KEY as per your account
  - Sample values of MW_AGENT_SERVICE for k8s yaml and docker file

    | Environment Variable | Docker     | Kubernetes                               |
    |----------------------|------------|------------------------------------------|
    | MW_AGENT_SERVICE     | 172.17.0.1 | mw-service.mw-agent-ns.svc.cluster.local |

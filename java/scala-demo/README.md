# Scala Demo Application

Simple demo application in Scala with instrumentation.
```env
Scala Version: 2.11.6
```

## Steps to set up

- Download the latest jar of [Middleware java agent](https://github.com/middleware-labs/opentelemetry-java-instrumentation/releases).
- Update the settings in `build.sbt` file as mentioned below and update the service name.
  ```sbt
  lazy val root = (project in file("."))
    .settings(
      ...
      (run / javaOptions) ++= Seq(
        "-javaagent:middleware-javaagent-{VERSION}.jar",
        "-Dotel.service.name={SCALA-APP-NAME}"
      )
      ...
    )
  ```
- Run the following command to start the application.
  ```shell
  MW_API_KEY = "wepjqgwgvwwtmwjanppcatwxmegaqardaxmf" sbt run
  ```
import Dependencies.*

name := "scala-demo"

scalaVersion     := "2.11.6"
version          := "0.1.0-SNAPSHOT"
organization     := "com.example"
organizationName := "example"

libraryDependencies ++= Seq(
  "com.typesafe.akka" %% "akka-http" % "10.1.15",
  "com.typesafe.akka" %% "akka-stream" % "2.5.32"
)

lazy val root = (project in file("."))
  .settings(
    name := "scala-demo",
    libraryDependencies += munit % Test,
    outputStrategy := Some(StdoutOutput),
    connectInput := true,
    run / fork := true,
    (run / javaOptions) ++= Seq(
      "-javaagent:middleware-javaagent-1.3.0.jar",
      "-Dotel.service.name=scala-app"
    )
  )

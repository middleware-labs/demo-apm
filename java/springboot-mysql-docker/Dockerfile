FROM eclipse-temurin:17-jdk-alpine

RUN apk --no-cache add libstdc++

WORKDIR /app

COPY .mvn .mvn/
COPY src src/
COPY mvnw .
COPY mvnw.cmd .
COPY pom.xml .
COPY middleware-javaagent-1.3.0.jar .


RUN chmod +x mvnw
RUN ./mvnw clean package -DskipTests

EXPOSE 8080

FROM maven:3.9-eclipse-temurin-21-alpine AS deps

WORKDIR /app
COPY pom.xml .

RUN mvn de.qaware.maven:go-offline-maven-plugin:resolve-dependencies -B

FROM maven:3.9-eclipse-temurin-21 AS  build
WORKDIR /app
COPY --from=deps /root/.m2/repository /root/.m2/repository
COPY src/ /app/src
COPY pom.xml .
RUN mvn clean package -DskipTests
FROM maven:3.9-eclipse-temurin-21 AS dev

WORKDIR /app

COPY --from=deps /root/.m2/repository /root/.m2/repository

COPY ./docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

RUN apt-get update && apt-get install -y curl inotify-tools

ENV SPRING_DEVTOOLS_RESTART_ENABLED=true
ENV SPRING_DEVTOOLS_LIVERELOAD_ENABLED=true

ENTRYPOINT ["/docker-entrypoint.sh"]


FROM openjdk:21-ea-21-jdk-slim AS release
RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 1001 ubuntu
USER ubuntu
WORKDIR /home/ubuntu
COPY --from=build /app/target/*.jar app.jar

ENV SPRING_DEVTOOLS_RESTART_ENABLED=false
ENV SPRING_DEVTOOLS_LIVERELOAD_ENABLED=false

ENTRYPOINT [ "java", "-jar", "app.jar" ]

FROM maven:3.9-eclipse-temurin-21 AS test
WORKDIR /app
COPY pom.xml .
COPY --from=deps /root/.m2/repository /root/.m2/repository

ENV SPRING_DEVTOOLS_RESTART_ENABLED=false
ENV SPRING_DEVTOOLS_LIVERELOAD_ENABLED=false

# RUN mvn dependency:go-offline
RUN mkdir -p /app/target && chmod -R 777 /app/target
# Prepare Maven dependencies
CMD ["sh", "-c", "mvn test -Dspring.profiles.active=test -Dsurefire.useFile=false"]
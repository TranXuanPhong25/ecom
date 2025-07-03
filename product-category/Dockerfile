FROM maven:3.9-eclipse-temurin-21-alpine AS deps

WORKDIR /app

RUN mvn go-offline:resolve-dependencies

FROM maven:3.9-eclipse-temurin-21 AS  build
WORKDIR /app
COPY --from=deps /root/.m2/repository /root/.m2/repository

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
COPY --from=build /app/target/*.jar app.jar

ENV SPRING_DEVTOOLS_RESTART_ENABLED=false
ENV SPRING_DEVTOOLS_LIVERELOAD_ENABLED=false

RUN addgroup --system app && adduser -S -s /bin/false -G app app
RUN chown -R app:app /
EXPOSE 8097
USER app
ENTRYPOINT [ "java", "-jar", "app.jar" ]
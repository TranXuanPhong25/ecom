FROM maven:latest AS build
COPY pom.xml .
COPY .mvn .mvn
COPY mvnw .
RUN mvn dependency:go-offline -B
COPY src ./src
RUN mvn clean package -DskipTests

FROM openjdk:21-ea-21-jdk-slim
COPY --from=build /target/*.jar app.jar
EXPOSE 8097
ENTRYPOINT [ "java", "-jar", "app.jar" ]
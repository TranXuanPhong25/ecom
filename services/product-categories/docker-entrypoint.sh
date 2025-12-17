#!/bin/sh

# Start background watcher to detect file changes
while inotifywait -r -e modify,create,delete,move ./src/main/;
do
  echo "Changes detected, recompiling..."
  mvn compile -o -DskipTests
done > /dev/null 2>&1 &

# Run Spring Boot with DevTools enabled for hot-reload
mvn spring-boot:run -Dspring-boot.run.jvmArguments="-Dspring.devtools.restart.enabled=true -Dspring.devtools.livereload.enabled=true"
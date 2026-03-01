FROM quay.io/strimzi/kafka:0.49.1-kafka-4.1.1

USER root

RUN mkdir -p /opt/kafka/plugins/debezium-postgres \
   && curl -fsSL https://repo1.maven.org/maven2/io/debezium/debezium-connector-postgres/3.4.0.Final/debezium-connector-postgres-3.4.0.Final-plugin.tar.gz \
   | tar -xz -C /opt/kafka/plugins/debezium-postgres

RUN chown -R 1001:0 /opt/kafka/plugins

USER 1001

# ping-kafka

## environment
 - HOST=127.0.0.1
 - PORT=9092
 - TIMEOUT=1000

## compose

```yml
# sample:ping kafka
# only container(host is not) can access kakfa
services:  
  kafka-server:
    container_name: test-kafka
    environment:
      JMX_PORT: 9097
      KAFKA_ADVERTISED_HOST_NAME: test-kafka
      KAFKA_ADVERTISED_PORT: 9092
      KAFKA_BROKER_ID: 1
      KAFKA_DELETE_TOPIC_ENABLE: "true"
      KAFKA_HEAP_OPTS: -Xmx1G
      KAFKA_JMX_OPTS: -Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.authenticate=false  -Dcom.sun.management.jmxremote.ssl=false
        -Dcom.sun.management.jmxremote.authenticate=false
        -Djava.rmi.server.hostname=test-kafka
      KAFKA_JVM_PERFORMANCE_OPTS: -XX:+UseG1GC -XX:MaxGCPauseMillis=20 -XX:InitiatingHeapOccupancyPercent=35
        -XX:+DisableExplicitGC -Djava.awt.headless=true
      KAFKA_LOG_CLEANER_ENABLE: "true"
      KAFKA_LOG_CLEANUP_POLICY: delete
      KAFKA_LOG_DIRS: /logs/kafka-logs-24bf1bde016a
      KAFKA_LOG_RETENTION_HOURS: 120
      KAFKA_ZOOKEEPER_CONNECT: test-zookeeper:2181
      KAFKA_ZOOKEEPER_CONNECTion_timeout_ms: 60000
    # extra_hosts:
    # - test-kafka:10.202.101.43
    image: pangpanglabs/kafka
    ports:
    - 9092
    - 9097
  zookeeper-server:
    container_name: test-zookeeper
    environment:
      ZOO_MY_ID: 1
      ZOO_SERVERS: server.1=test-zookeeper:2888:3888
    # extra_hosts:
    # - test-kafka:10.202.101.43
    image: zookeeper:3.4.9
    ports:
    - 2181
    - 2888
    - 3888
  ping-kafka-server:
    container_name: test-ping-kafka
    command: sh -c 'echo "wait kafka..." && /go/bin/wait-for.sh test-kafka:9092 -t 36000 -- ./ping-kafka'
    depends_on:
    - kafka-server
    environment:
    - KAFKA_HOST=test-kafka
    - PORT=9092
    - TIMEOUT=1000
    image: relaxed/ping-kafka
    volumes:
      - ./wait-for.sh:/go/bin/wait-for.sh
    ports:
    - 8080
version: "3"

```

[wait-for.sh](https://github.com/eficode/wait-for.git)
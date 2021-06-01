# Kafka consumer and producer with Go

Simple example that produces data to Kafka topic and consumer that consumes that data.

## Requirements

- Apache Kafka (https://kafka.apache.org/downloads)

## Setup

After unpacking Apache Kafka, cd into the directory in order to execute the following executables

1. Start Zookeeper
    ```bash
    bin/zookeeper-server-start.sh config/zookeeper.properties
    ```
2. Create separate configuration in order to start 3 brokers for our cluster:
   ```bash
   cp config/server.properties config/server0.properties
   cp config/server.properties config/server1.properties
   cp config/server.properties config/server2.properties
   
   # Edit each of those .properties files and change the following:
   # 
   # broker.id = 0 # increment for the next config
   # listeners=PLAINTEXT://localhost:9092 # increment port for the next config
   # log.dir=/tmp/kafka-logs-1 # increment number for the next config
   ```
3. Run each broker in separate terminal tab
   ```bash
   bin/kafka-server-start.sh config/server0.properties
   bin/kafka-server-start.sh config/server1.properties
   bin/kafka-server-start.sh config/server2.properties
   ```
4. Create a topic

   ```bash
   bin/kafka-topics.sh --zookeeper localhost:2181 --create \
     --topic golang-topic \
     --replication-factor 2 \
     --partitions 2
   ```

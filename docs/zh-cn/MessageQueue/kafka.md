# Kafka


### 1. QuickStart


### 2. 概念

- Topic
- Partition
- Producer
- Consumer
- Broker


### 3. How to start

- docker

```
kafka
kafka-manager
zookeeper
```

- docker-compose

```
docker-compose up -d
docker-compose stop
docker-compose.yaml

```

```
version: '2'
services:
  ui:
    image: sheepkiller/kafka-manager:latest
    depends_on:
    - zookeeper
    ports:
    - 9000:9000
    environment:
      ZK_HOSTS: zookeeper:2181
      APPLICATION_SECRET: letmein
    restart: 'no'
  zookeeper:
    image: wurstmeister/zookeeper:latest
    ports:
    - 2181:2181
  server:
    image: wurstmeister/kafka:latest
    depends_on:
    - zookeeper
    ports:
    - 9092:9092
    environment:
      KAFKA_OFFSETS_TOPIC_REPLIATION_FACTOR: 1
      KAFKA_ADVERTISED_HOST_NAME: 10.1.10.11
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
    volumes:
    - /var/run/docker.sock:/var/run/docker.sock
    - /Users/xiewei/docker/src/kafka-logs:/kafka
```

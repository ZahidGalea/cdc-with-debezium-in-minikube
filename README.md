# Currently Working on this :)

---

# Change Data Capture Demo

## What is my plan?

1) Generate a DB with a Logistic Model and an API to generate Test traffic into. :heavy_check_mark:
2) Crete a replication of an Oracle DB to Kafka in near real-time using Debezium. :heavy_check_mark:
3) Subscribe to the kafka topics a series of applications like:
    * Realtime Dashboards
    * Apache Beam processing for real time analytics
    * ML Model implementation

---

## 0 - Start Config

```
# In my case I start and use the docker daemon trought CLI
sudo dockerd

# Start with the minikube configuration:

minikube config set memory 5000
minikube start

-----
# Creating our demo namespaces:

kubectl apply -f namespaces/
kubectl config set-context --current --namespace=application-ns

```

## 1 - Oracle Database Setup

Let's set a simple Database :sweat_smile:

```
# First of all.. you will need to build the oracledb image using the following guide:
# https://github.com/oracle/docker-images/tree/main/OracleDatabase/SingleInstance
# And as result you will have an oracle image ready to use locally
# as example:
# cd OracleDatabase/SingleInstance/dockerfiles/19.3.0
# You will have to download the linux x64 binaries from the oracle DB then:
# docker build -t oracle/database:19.3.0-ee --build-arg DB_EDITION=ee .
# Or add the image to minikube with:
# minikube image load oracle/database:19.3.0-ee

# To use or not use the docker daemon of the minikube:
eval $(minikube docker-env)
eval $(minikube docker-env -u)


# Sets the context to avoid using of parameter namespace
kubectl config set-context --current --namespace=application-ns

# Lets create some secrets first:
kubectl apply -f secrets/*

# Configmap creation for startup SQL
# if exists: kubectl delete configmap filler-app-db-creation
kubectl create configmap filler-app-db-creation --from-file=database/startup-scripts/filler-app-db-creation.sql -n application-ns
# This configmap will be used for the logminer configuration
# if exists: kubectl delete configmap log-miner-config
kubectl create configmap log-miner-config --from-file=database/startup-scripts/setup-logminer.sh -n application-ns

# Applying the database and the service nodeport resources
kubectl apply -f database/db.yml

# To test the database  Forward the minikube port
kubectl port-forward service/oracle18xe-svc 1521:1521
```

## 2 - Filler APP Build

```

# Lets dockerize our java jar filler app and save it
# (BTW, it comes from an app that I created before)
# https://github.com/ZahidGalea/logistics-spring-boot-app
# Just package it and use it if u want! 
docker build -t=logistic-app:latest filler-application/Dockerfile

# Also an script that makes request to this app
docker build -t=simulation-logistic-app:latest filler-application/simulation_app/

# The following ENV Variables must be set in the deployment:
# ORACLE_DB_HOST: localhost or an IP
# ORACLE_DB_PORT: 1521 mostly for all oracle db 
# ORACLE_DB_NAME: DB where the app will create the tables and fill with data
# ORACLE_DB_USERNAME: ---
# ORACLE_DB_PASSWORD: ---

# Lets startup our application:
kubectl apply -f filler-application/filler-application.yml

# And a simple script with a simulation of transactions:
kubectl apply -f filler-application/transaction-simulation.yml

# Open a new terminal and, stream your logs...
kubectl logs -f deploy/filler-app

# Watch that the application is logging correctly....

```

## 3 - Zookepeker, Kafka & Kafka connect

```

# Before everything for this demo:
# In order to work with oracle we will have to mount a directory into minikube first:
# And keep it running btw! 
# This is required for the kafka connect Oracle Connector
minikube mount ${PWD}/debezium-connector-oracle:/oracle_data/

# Kafka folder will create zookper, kafka with 3 replicas, manager, schema registry and kafka connect
# The containers will fail because doesn't exist defined dependencies between them, just wait some minutes.
kubectl apply -f kafka/*

# How to open kafka manager?
# minikube service -n application-ns kafka-manager --url

```

## 4 - Streaming DB Changes with debezium

```

# Now, set the debezium connector into the kafka connect using a simple curl. but first we will need to expose
# the port 8083 to be able to curl it from our computer.
kubectl port-forward service/kafka-connect-svc 8083:8083

# Test it with:
# curl -H "Accept:application/json" localhost:8083/
# It should return you something like:
# {"version":"7.0.1-ccs","commit":"b7e52413e7cb3e8b","kafka_cluster_id":"287OqZkVRgi3TI80fjxJCg"}

# https://debezium.io/documentation/reference/stable/connectors/oracle.html#required-debezium-oracle-connector-configuration-properties
# Now register the debezium connector with something like:
POST localhost:8083/connectors/
{
  "name": "fillerapplication-connector",
  "config": {
    "connector.class": "io.debezium.connector.oracle.OracleConnector",
    "database.hostname": "oracle18xe-svc",
    "database.port": "1521",
    "database.user": "c##dbzuser",
    "database.password": "dbz",
    "database.dbname": "ORCLCDB",
    "database.pdb.name": "ORCLPDB1",
    "database.server.name": "filler_application",
    "database.connection.adapter": "logminer",
    "table.include.list": "FILLERAPPLICATION.ESTADO_ENVIO,FILLERAPPLICATION.ENVIO",
    "event.processing.failure.handling.mode": "warn",
    "poll.interval.ms": "2000",
    "tasks.max" : "1",
    "database.history.kafka.bootstrap.servers": "kafka-svc:9092",
    "database.history.kafka.topic": "schema-changes.fillerapplication",
    "snapshot.mode": "initial"
  }
}
# With curl:
curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" \
localhost:8083/connectors/ --data '{"name":"fillerapplication-connector","config":{"connector.class":"io.debezium.connector.oracle.OracleConnector","database.hostname":"oracle18xe-svc","database.port":"1521","database.user":"c##dbzuser","database.password":"dbz","database.dbname":"ORCLCDB","database.pdb.name":"ORCLPDB1","database.server.name":"filler_application","database.connection.adapter":"logminer","table.include.list":"FILLERAPPLICATION.ESTADO_ENVIO,FILLERAPPLICATION.ENVIO","event.processing.failure.handling.mode":"warn","poll.interval.ms":"2000","tasks.max":"1","database.history.kafka.bootstrap.servers":"kafka-svc:9092","database.history.kafka.topic":"schema-changes.fillerapplication","snapshot.mode":"initial"}}'

# If u want to list the connectors:
# curl -H "Accept:application/json" localhost:8083/connectors/

# If u want to get the status of a connector:
# curl -H "Accept:application/json" localhost:8083/connectors/?expand=status

# If u want to restart it:
# curl -H "Accept:application/json" -X POST localhost:8083/connectors/inventory-connector/restart

# If u want to delete it:
# curl -X DELETE http://localhost:8083/connectors/fillerapplication-connector

```


## N - End it

```
minikube docker-env --unset
minikube stop
minikube delete
```

## Notes

[Notes Readme](NOTES.md)

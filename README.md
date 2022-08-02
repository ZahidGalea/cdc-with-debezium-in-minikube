# Currently Working on this :)

---

# Change Data Capture Demo

## What is my plan?

1) Generate a DB with a Logistic Model and an API to generate Test traffic into. :heavy_check_mark:
2) Crete a replication of PostgresSQL Database to Kafka in near real-time using Debezium.
3) Subscribe to the kafka topics a series of applications like:
    * Realtime Dashboards
    * Apache Beam processing for real time analytics
    * ML Model implementation

---

## 0 - Start Config

```bash
# Start with the minikube configuration:
minikube start --memory 11000 --cpus 6

# Creating our demo namespaces:
kubectl apply -f namespaces/
kubectl config set-context --current --namespace=kafka
```

## 1 - Postgres Database Setup

Let's set a simple Database :sweat_smile: In this case a PostgreSQL Database
https://bitnami.com/stack/postgresql/helm

```bash
# To use or not use the docker daemon of the minikube:
eval $(minikube docker-env)
#eval $(minikube docker-env -u)


#kubectl create configmap log-miner-config --from-file=database/startup-scripts/setup-logminer.sh -n kafka
kubectl create configmap filler-app-db-creation --from-file=database/startup-scripts/db.creation.sh -n kafka

# Create the database (It will apply the startup scripts folder with the above configmaps)
kubectl apply -f database/db.yml

# Test the connection and the database if u want with:
# kubectl port-forward service/postgres-svc 5432:5432

```

## 2 - Filler APP Build

```bash
# Lets dockerize our java jar filler app and save it
# (BTW, it comes from an app that I created before)
# https://github.com/ZahidGalea/logistics-spring-boot-app
# Just package it and use it if u want! 
# Be aware that the image must be available to the minikube daemon
# To use or not use the docker daemon of the minikube:
eval $(minikube docker-env)
#eval $(minikube docker-env -u)

docker build -t=logistic-app:latest filler-application/

# Also an script that makes request to this app
docker build -t=simulation-logistic-app:latest filler-application/simulation_app/

# Lets create some secrets first with the application:
kubectl apply -f secrets/

# Lets startup our application:
kubectl apply -f filler-application/filler-application.yml

# And a simple script with a simulation of transactions:
kubectl apply -f filler-application/transaction-simulation.yml

# Watch that the application is logging correctly....
# kubectl logs -f deploy/filler-app


```

## 3 - Zookepeker, Kafka & Kafka connect

```bash

# Kafka folder will create zookper, kafka with 3 replicas, manager, schema registry and kafka connect
# The containers will fail because doesn't exist defined dependencies between them, just wait some minutes.
kubectl apply -f kafka
# How to open kafka manager?
# minikube service -n cdc-demo kafka-manager --url

```

## 4 - Streaming DB Changes with debezium

```bash
# Now, set the debezium connector into the kafka connect using a simple curl. but first we will need to expose
# the port 8083 to be able to curl it from our computer.
kubectl port-forward service/kafka-connect-svc 8083:8083
```
```bash
# Test it with:
# curl -H "Accept:application/json" localhost:8083/
# It should return you something like:
# {"version":"7.0.1-ccs","commit":"b7e52413e7cb3e8b","kafka_cluster_id":"287OqZkVRgi3TI80fjxJCg"}

# https://debezium.io/documentation/reference/stable/connectors/oracle.html#required-debezium-oracle-connector-configuration-properties
# Now register the debezium connector with something like:
#POST localhost:8083/connectors/
{
  "name": "logisticapp",
  "config": {
    "connector.class": "io.debezium.connector.postgresql.PostgresConnector",
    "plugin.name": "pgoutput",
    "database.hostname": "postgres-svc",
    "database.port": "5432",
    "database.user": "dbz_rpl_user",
    "database.password": "vn53nag",
    "database.dbname": "logisticapp",
    "database.server.name": "postgres"
    "table.include.list" : "public.envio"
  }
}

```

```bash

# With curl:
curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" \
localhost:8083/connectors/ --data '{"name":"logisticapp","config":{"connector.class":"io.debezium.connector.postgresql.PostgresConnector","plugin.name":"pgoutput","database.hostname":"postgres-svc","database.port":"5432","database.user":"dbz_rpl_user","database.password":"vn53nag","database.dbname":"logisticapp","database.server.name":"postgres","table.include.list":"public.envio"}}'

# If u want to delete a connector:
# curl -X DELETE localhost:8083/connectors/logisticapp
```

```bash
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

```bash
minikube docker-env --unset
minikube stop
minikube delete
```

## Notes

[Notes Readme](NOTES.md)

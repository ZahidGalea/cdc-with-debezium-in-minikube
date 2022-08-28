
# CDC with debezium to a PostgreSQL Database

## What is my plan?

1) Generate a DB with a Logistic Model and an API to generate Test traffic into. :heavy_check_mark:
2) Crete a replication of PostgresSQL Database to Kafka in near real-time using Debezium. :heavy_check_mark:

---

## 0 - Start Config

Requeriments:
* docker
* minikube 
* envsubst command


```bash
# Start with the minikube configuration:
minikube start --memory 11000 --cpus 6  --insecure-registry "10.0.0.0/24"
minikube addons enable registry

kubectl create ns debezium-example
kubectl config set-context --current --namespace=debezium-example
```

## 1 - Postgres Database Setup

Let's set a simple Database :sweat_smile: In this case a PostgreSQL Database
https://bitnami.com/stack/postgresql/helm

```bash
# To use or not use the docker daemon of the minikube:
eval $(minikube docker-env)
#eval $(minikube docker-env -u)

kubectl create configmap filler-app-db-creation --from-file=database/startup-scripts/db.creation.sh

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

## 3 - Zookepeker, Kafka & Kafka connect (With strimzi)

* Strimzi Operator
```bash
curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.20.0/install.sh | bash -s v0.20.0
kubectl create -f https://operatorhub.io/install/strimzi-kafka-operator.yaml
```
* Kafka cluster

```bash
kubectl create -n debezium-example -f kafka/kafka.yml
```

* Kafka connect

```bash
export MNK_REGISTRY_IP=$(kubectl -n kube-system get svc registry -o jsonpath='{.spec.clusterIP}')
envsubst < kafka/kafka-connect.yml | kubectl apply -f -
```

## 4 - Streaming DB Changes with debezium

Create the DBZ connector with the following file:
```bash
kubectl apply -f kafka/dbz-connector.yml
```

If u want to list strimzi resources, and delete wrong connectors just use some of those:
```bash
kubectl get strimzi
kubectl describe kafkaconnector.kafka.strimzi.io/debezium-connector-postgresql
kubectl delete kafkaconnector.kafka.strimzi.io/debezium-connector-mysql
```

To check the message in the topics you can use something like this:
```bash
kubectl run -n debezium-example -it --rm --image=quay.io/debezium/tooling:1.2  \
--restart=Never watcher -- kcat -b debezium-cluster-kafka-bootstrap:9092 -C -o beginning \
-t postgres.public.envio
```

## N - End it

```bash
minikube docker-env --unset
minikube stop
minikube delete
```

# Currently Working on this :)

---

# Change Data Capture Demo

## What is my plan?

1) Generate a DB with a Logisitc Model and an API to generate Test traffic into. :heavy_check_mark:
2) Crete a replication of the DB using Debezium + Zookeper +Kafka
3) Subscribe to the kafka topics a serie of applications like:
    * Realtime Dashboards
    * Apache Beam processing for real time analytics
    * ML Model implementation
4) Automatic deployment and integration trought a CI/CD Pipeline using Gitlab runners and Github as source repo.

---

## 0 - Start Config

```
# Start with the minikube configuration:

minikube config set memory 6900
minikube start
minikube addons enable metrics-server

# This able us to use localmachine images
eval $(minikube docker-env)

-----
# Creating our demo namespaces:

kubectl apply -f namespaces/

```

## 1 - Oracle Database Setup

Lets set a simple Database :sweat_smile:

```
# Sets the context to avoid using of parameter namespace
kubectl config set-context --current --namespace=oracle-ns

# Configmap creation for startup SQL, also .sh can be added to the same path
kubectl create configmap filler-app-db-creation --from-env-file=database/startup-scripts/filler-app-db-creation.sql -n oracle-ns

# Dockere login could be required to pull the oracle image
docker login... 

# Downloading a custom image to avoid official oracle download ()
docker pull pvargacl/oracle-xe-18.4.0:latest

# Applying the database and the service nodeport resources
kubectl apply -f database/node-port-service.yml
kubectl apply -f database/db.yml

## Important! 
# It exposes the service to make it work if your application if outside
# minikube service -n oracle-ns oracle18xe --url
# Or you can expose a localhost port to the pod using: as example
# kubectl port-forward svc/oracle18xe-svc 30007:1521

# DB Credentials as System:
# usuario: system
# pw: oracle

# Tip to get into the running pod 
# kubectl exec --stdin --tty POD -- /bin/sh

```

[Database Readme](database/README.md)

## 2 - Kafka & Zookeper

```
# Creates namespace for kafka
kubectl apply -f namespaces/kafka.yml
kubectl config set-context --current --namespace=kafka-ns

------
# Creates zookeper
kubectl apply -f zookeper/zookeper.yml

# Creates kafka cluster with 3 replicas
kubectl apply -f kafka/kafka.yml

# Creates the kafka manager
kubectl apply -f kafka/kafka-manager.yml

# How to open kafka manager?
# minikube service -n kafka-ns kafka-manager --url
# Open the resultant url from the output

```

## 3 - Filler APP Build

```
# Lets dockerize our java jar filler app and save it
docker build -t=logistic-app:latest filler-app/

```

## N - End it

```
minikube stop
minikube delete
```

## Thanks to:

##### Zookeper & kafka deployment

https://github.com/d1egoaz/minikube-kafka-cluster

##### Oracle setup un minikube:

https://ronekins.com/2020/03/14/running-oracle-12c-in-kubernetes-with-minikube-and-virtualbox/

## Notes

[Notes Readme](NOTES.md)

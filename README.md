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
eval $(minikube docker-env)

-----
# Creating our demo namespaces:

# Database namespace
kubectl apply -f namespaces/oracle-namespace.yml

```

## 1 - Oracle Database Setup

Lets set a simple Database :sweat_smile:

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
minikube service -n kafka-ns kafka-manager --url


```

## X - End it

```
minikube stop
minikube delete
```

## Resources copied or replicated:

##### Zookeper & kafka deployment

https://github.com/d1egoaz/minikube-kafka-cluster

## Notes

[Notes Readme](NOTES.md)

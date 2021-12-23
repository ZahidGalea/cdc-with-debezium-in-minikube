# Currently Working on this :)

## Change Data Capture Demo

---

### What is my plan? 

1) Generate a DB with a Logisitc Model and an API to generate Test traffic into. :heavy_check_mark:
2) Crete a replication of the DB  using Debezium + Kafka
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
minikube addons enable metrics-server
minikube start
eval $(minikube docker-env)

-----
# Creating our demo namespaces:

# Database namespace
kubectl apply -f namespaces/oracle-namespace.yml

```


## 1 - Oracle Database Setup

Lets set a simple Database :sweat_smile:

[Database Readme](database/README.md)


## X - End it

```
minikube stop
minikube delete
```
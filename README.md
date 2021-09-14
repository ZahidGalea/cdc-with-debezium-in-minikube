# Change Data Capture Demo

---

## 0 - Start Config

```
# Start with the minikube configuration:

minikube config set memory 10240
minikube addons enable metrics-server
minikube start
eval $(minikube docker-env)

-----
# Creating our demo namespaces:

# Database namespace
kubectl apply -f namespaces/oracle-namespace.yml

```


## 1 - Oracle Database Setup

[Database Readme](database/README.md)
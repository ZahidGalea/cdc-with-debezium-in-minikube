# Oracle Setup

Guide: https://ronekins.com/2020/03/14/running-oracle-12c-in-kubernetes-with-minikube-and-virtualbox/

Apuntes:
```
# Sets the context to avoid using of parameter namespace
kubectl config set-context --current --namespace=oracle-namespace

# It creates a configmap file into the namespace for reusability
kubectl create configmap oradb --from-env-file=oracle.properties -n oracle-namespace

# Dockere login to pull the requireed image
docker login... 

# Downloading a custom image to avoid official  oracle download
docker pull pvargacl/oracle-xe-18.4.0:latest

# Applying the database and the service nodeport resources
kubectl apply -f node-port-service.yml
kubectl apply -f db.yml

## Important! 
# It exposes the service to make it work
minikube service -n oracle-namespace oracle18xe --url

# DB Credentials:
usuario: system
pw: oracle

```

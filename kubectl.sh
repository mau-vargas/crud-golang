#!/usr/bin/env bash
kubectl create -f crudgolangapi-service.yaml,crudgolangapi-deployment.yaml,crudgolangapi-claim0-persistentvolumeclaim.yaml,crudgolang-service.yaml,crudgolang-deployment.yaml
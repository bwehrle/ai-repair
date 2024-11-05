#!/bin/zsh

if (( # == 0 )); then
   print >&2 "Usage: $0 (create|delete)"
   exit 1
elif [[ $1 == "create" || $1 == "delete" || $1 == "render" ]]; then
else
echo >&2 "Usage: $0 (create|delete) [--upsert]"
exit 1
fi

if [[ $1 == "create" ]]; then
kubectl apply -f namespace.yaml
helm repo add bwehrle-ms https://bwehrle.github.io/helm-microservice/
helm repo update
helm install httpbin-staging manifest --values manifest/staging/values.yaml --namespace httpbin-helm --version 0.6.2

elif [[ $1 == "render" ]]; then
helm template httpbin-staging manifest --values manifest/staging/values.yaml --namespace httpbin-helm --version 0.6.2

elif [[ $1 == "delete" ]]; then
helm uninstall httpbin-staging --namespace httpbin-helm
kubectl delete namespace httpbin
fi
#!/bin/bash

set -euo pipefail

# Local registery
docker run -d -p 5000:5000 --restart=always --name registry registry:2

# Install Istio
curl -L https://github.com/knative/serving/releases/download/v0.4.0/istio.yaml \
  | kubectl apply --filename -

# Label the default namespace with istio-injection=enabled.
kubectl label namespace default istio-injection=enabled

# Install Knative Serving
curl -L https://github.com/knative/serving/releases/download/v0.4.0/serving.yaml \
  | kubectl apply --filename -

# Install Knative Build
curl -L https://github.com/knative/build/releases/download/v0.4.0/build.yaml \
  | kubectl apply --filename -

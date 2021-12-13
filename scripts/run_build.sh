#!/usr/bin/env sh

_registry="$1"
_tag="$2"
_image="$3"
_platform="linux/amd64,linux/arm64,linux/386"

if [ -z "$_registry" ] || [ -z "$_tag" ] || [ -z "$_image" ]; then
  echo "Please specify image repository and tag img and "
  exit 0;
fi

# prepare dir
mkdir -p ./bin
# build demo app
CGO_ENABLED=0 GOOS=linux GOARCH= go build -o ./bin/PrometheusAlert main.go

docker push "$_repository/$_image:latest"

# clean dir bin
rm -rf ./PrometheusAlert
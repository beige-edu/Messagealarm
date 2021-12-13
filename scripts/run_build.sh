#!/usr/bin/env sh

_registry="$1"
_tag="$2"
_img="$3"
_platform="linux/amd64,linux/arm64,linux/386"

if [ -z "$_registry" ] || [ -z "$_tag" ] || [ -z "$_img" ]; then
  echo "Please specify image repository and tag img and "
  exit 0;
fi

# create and use builder
docker buildx inspect builder >/dev/null 2>&1
if [ "$?" != "0" ]; then
  docker buildx create --use --name builder
fi

# prepare dir
mkdir -p ./bin
# build demo app
CGO_ENABLED=0 go build -o ./bin/PrometheusAlert main.go

# docker image
docker buildx build --platform "$_platform" \
  -f "build/Dockerfile" \
  -t "docker.webtest.51beige.com/crm-warning-backend:latest" \
  --push .

# clean dir bin
rm -rf ./PrometheusAlert
#!/usr/bin/env sh

_registry="$1"
_tag="$2"
_image="$3"
_platform="linux/arm64"

if [ -z "$_registry" ] || [ -z "$_tag" ] || [ -z "$_image" ]; then
  echo "Please specify image repository and tag img and "
  exit 0;
fi

# prepare dir
mkdir -p ./bin
# build demo app
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/PrometheusAlert main.go


# build and push docker image
docker build -t "$_image:latest" -f "build/Dockerfile" .

# docker push "$_repository/$_image:latest"

# clean dir bin
# rm -rf ./PrometheusAlert
#!/bin/bash

if [ $# -lt 2 ]; then
  echo "err: serve or env not found. env in [prod, test]. eg: ./build.sh rpc prod"
  exit 1
fi

SERVE_NAME=$1
ENV=$2

help() {
    echo "Usage:"
    echo "    ./build.sh [服务] [环境(test|prod)]"
    echo "    ./build.sh rpc prod        [*] 编译 rpc 服务"
    echo "    ./build.sh api prod        [*] 编译 api 服务"
}

case $1 in
rpc)
  DOCKER_BUILDKIT=0 docker build -t "serve.ucenter-"${SERVE_NAME} --build-arg CONF_ENV="${ENV}" -f ./rpc/Dockerfile .
  docker image prune -f
  ;;
api)
  DOCKER_BUILDKIT=0 docker build -t "serve.ucenter-"${SERVE_NAME} --build-arg CONF_ENV="${ENV}" -f ./api/Dockerfile .
  docker image prune -f
  ;;
*)
  help
  ;;
esac
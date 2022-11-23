#!/bin/bash

if [ $# -lt 2 ]; then
  echo "err: serve or env not found. env in [prod, test]. eg: ./build.sh user-rpc prod"
  exit 1
fi

SERVE_NAME=$1
ENV=$2

help() {
    echo "Usage:"
    echo "    ./build.sh [服务] [环境(test|prod)]"
    echo "    ./build.sh user-rpc prod        [*] 编译 用户 rpc 服务"
    echo "    ./build.sh user-api prod        [*] 编译 用户 api 服务"
}

case $1 in
user-rpc)
  DOCKER_BUILDKIT=0 docker build -t "serve."${SERVE_NAME} --build-arg CONF_ENV="${ENV}" -f ./app/ucenter/rpc/Dockerfile .
  docker image prune -f
  ;;
user-api)
  DOCKER_BUILDKIT=0 docker build -t "serve."${SERVE_NAME} --build-arg CONF_ENV="${ENV}" -f ./app/ucenter/api/Dockerfile .
  docker image prune -f
  ;;
*)
  help
  ;;
esac
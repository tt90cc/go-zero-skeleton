#!/bin/bash

# eg:
# ./build.sh api:8214 test

serve_name="tt90.ucenter"
log_path="/tmp"

src=`echo $1 | sed -r 's/([^:_]+)?:([^:_]+)?/\1/'`
port=`echo $1 | sed -r 's/([^:_]+)?:([^:_]+)?/\2/'`
cluster=$2
project_path=$(cd `dirname $0`; pwd)
src_path=$project_path"/"$src
full_serve_name="go."$serve_name"-"$src

if [ ! -d $src_path ]; then
  printf "main path err.\n"
  exit 100
fi

if [ ! -d $src_path"/etc/"$cluster ]; then
    printf "not found cluster path.\n"
    exit 101
fi

log_path=$log_path"/"$full_serve_name
if [ ! -d $log_path ]; then
  mkdir -p $log_path
  if [ $? != 0 ]; then
    printf "mkdir log_path failed.\n"
    exit 103
  fi
fi

# shellcheck disable=SC2120
printEnv(){
  printf "\n"
  printf "Print Env \n"
  printf "============================================\n"
  printf "Commond Params        | src:%s port:%s env:%s\n" $src $port $cluster
  printf "Full Serve Name       | %s\n" $full_serve_name
  printf "Project Path          | %s\n" $project_path
  printf "Src Path              | %s\n" $src_path
  printf "Log Path              | %s\n" $log_path
  printf "============================================\n\n\n"
}

# copy to pull script
#>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
#cicd_dir=/tmp/cicd
#
#pull(){
#  printf "Pull from git \n"
#  printf "============================================\n"
#
#  rm -rf $cicd_dir
#  mkdir -p $cicd_dir
#  cd $cicd_dir
#
#  git clone https://github.com/tt90cc/go-zero-skeleton.git && cd go-zero-skeleton
#
#  printf "============================================\n\n\n"
#}
#
#pull
#
#cd $cicd_dir"/go-zero-skeleton" && ./build.sh api:8214 test
#>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

build(){
  printf "Build Docker \n"
  printf "============================================\n"
  DOCKER_BUILDKIT=0 docker build -t $full_serve_name --build-arg CONF_ENV="${cluster}" -f $src_path"/Dockerfile" .
  docker image prune -f
  if [ $? != 0 ]; then
      printf "docker build failed.\n"
      exit 105
  fi

  printf "============================================\n\n\n"
}

clear(){
  printf "Clear Images and Container \n"
  printf "============================================\n"
  if docker ps | grep ${full_serve_name}; then
      docker stop ${full_serve_name}
  fi

  if docker ps -a | grep ${full_serve_name}; then
      docker rm ${full_serve_name}
  fi

  printf "============================================\n\n\n"
}

run(){
  printf "Run Docker \n"
  printf "============================================\n"
  docker run -d --restart=always --name $full_serve_name -p $port:$port -v $log_path:/app/logs $full_serve_name
  if [ $? != 0 ]; then
      printf "run docker failed.\n"
      exit 106
  fi

  printf "============================================\n\n\n"
}

printEnv
build
clear
run
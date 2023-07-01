#!/bin/bash

# eg: ./build.sh test
# eg: ./build.sh pro

exe="tt90.server"
remote="root@127.0.0.1"
workspace="/www/wwwroot/${exe}"

cluster=$1
branch="dev"

if [ $cluster == "pro" ]; then
  confirm_def="n"
  read -t 300 -p "confirm deploy pro. (y/n) [$confirm_def] " confirm
  confirm="${confirm:-$confirm_def}"
  if [ $confirm != "y" ]; then
      exit 0
  fi
  branch="master"
  remote="root@127.0.0.1"
fi

echo "==========================="
echo "拉取 ${branch} 代码"
echo "==========================="
echo ""

rm -rf /tmp/pullwk
mkdir -p /tmp/pullwk
cd /tmp/pullwk && \
git clone -b dev git@github.com:tt90cc/go-zero-skeleton.git && \
cd go-zero-skeleton

echo ""
echo "==========================="
echo "编译"
echo "==========================="
echo ""

go mod download && GOOS=linux GOARCH=amd64 go build -o ${exe} main.go

if [ $? -ne 0 ]; then
   echo "go build failed"
   exit 1
fi

echo "==========================="
echo "部署"
echo "==========================="
echo ""

builder_path="/tmp/publish"

rm -rf ${builder_path}
mkdir -p ${builder_path}/logs ${builder_path}/etc

mv ${exe} ${builder_path}
cp -r ${server}/etc/${cluster}/* ${builder_path}/etc
cp -r ${server}/template ${builder_path}

cd ${builder_path} && zip -q -r ${exe}.zip *

ssh ${remote} "mkdir -p ${workspace}"
scp ${exe}.zip ${remote}:${workspace}
ssh ${remote} "cd ${workspace} && unzip -o ${exe}.zip && chmod +x ${exe}"

echo "==========================="
echo "启动 ${exe} 服务"
echo "==========================="
echo ""

ssh $remote "/www/server/panel/pyenv/bin/supervisorctl restart ${exe}:${exe}_00"
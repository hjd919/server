#!/bin/bash

# push代码
# 上env服务器拉取代码
# 执行服务器上脚本 传service和env

service=$1 #服务名称
env=$2     #环境

# 定义
# 服务器根路径
serverDir=/opt/game_micro
# 服务器host
if [ "$env" == "pro" ]; then
    serverHost=101.200.41.141
else
    serverHost=39.96.187.72
fi

git add -A
git commit -m "推送部署"
git pull
git push

ssh root@${serverHost} "
source /etc/profile
cd ${serverDir}/deploy && \
git pull && \
/bin/bash ./compile.sh ${service} ${env}"

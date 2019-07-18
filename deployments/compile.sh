#!/bin/bash
service=$1 #服务名称
env=$2     #环境

# 拉取代码
# 判断并创建dist
# 复制配置
# 替换配置参数
# 停止，替换，启动

serverdir=/opt/jdcj_micro
serverDir=/opt/game_micro
runDir=/opt/game_micro_run
basedir=$serverDir
projectname=jdcj
rsyncServer=jdcj

echo "拉取代码"
git pull

if [ "$env" == "pro" ]; then
    serverhost=172.17.101.202
    dbname="dev_wxgame"
    redisPort="2"
else
    serverhost=172.17.101.205
    dbname="wxgame_go"
    redisPort="8"
fi

# 引入公共方法
. ./common.sh

compileApp() {
    local service=$1 #服务名称
    echo "---部署${service}---"
    cd $serverDir/src/$service

    echo "判断并创建dist"
    distDir=./dist/logs
    if [ ! -d "$distDir" ]; then
        mkdir -p $distDir
    fi

    echo "复制配置"
    copyToDist

    echo "改配置"
    replaceConfig $env

    echo "编译到dist"

    # 需要服务器安装所有依赖包。。包管理工具govendor|godep|makefile 1.挂载到docker的go环境中 2.挂载依赖包文件
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/server_upload --tags "consul" main.go

    echo "复制dist到运行目录"
    # mv dist $runDir/$service
    rsync -a $serverDir/src/$service/dist/* $runDir/$service

    echo "完成编译"
}
copyGlibs(){
    echo "部署glibs"
    rsync -a $serverDir/src/glibs $runDir
}

# *拉代码下来，并编译
case "$service" in
"glibs")
    copyGlibs
    ;;
"all")
    # copyGlibs

    # . ./all_config.sh
    # for ser in ${services[@]}; do
    #     compileApp $ser $env
    # done
    ;;
*)
    if [ $env ]; then
        # 其他微服务
        compileApp $service $env
    fi
    ;;
esac

# *进入部署服务器 运行rsync命令同步整个项目代码
ssh root@${serverhost} "rsync -avz root@${rsyncServer}::game_micro ${runDir}"

# *进入服务目录 重启服务
restartApp() {
    service=$1
    ssh root@${serverhost} "cd ${runDir}/${service} && \
        docker-compose -p ${projectname} stop && \
        mv server_upload ${service} && \
        docker-compose -p ${projectname} up -d && \
        docker ps|grep ${service} && \
        docker logs ${projectname}_${service}_1"
}

case "$service" in
"all")
    # . ./all_config.sh
    for ser in ${services[@]}; do
        restartApp $ser $env
    done
    ;;
*)
    if [ $env ]; then
        # 其他微服务
        restartApp $service $env
    fi
    ;;
esac

echo '部署完成'

#!/bin/bash

# 部署
# use: ./deploy.sh git|glibs|gateway|微服务目录名 提交备注
# example: ./deploy.sh gateway dev|pro
# example: ./deploy.sh personality dev|pro
# example: ./deploy.sh all dev|pro

# 定义输入
service=$1 #服务名称
env=$2     #环境 test/pro

# 定义服务器
if [ "$env" == "pro" ]; then
    serverhost=101.200.41.141
else
    serverhost=39.96.187.72
fi

serverdir=/opt/jdcj_micro

# 定义项目名
projectname=jdcj

# 定义本地路径
basedir=$(
    cd $(dirname $0)
    pwd -P
)

. ./deploy/common.sh

deployApp() {
    local deploy=$1
    local env=$2

    echo "正在部署服务...${deploy}"

    # 进入服务目录
    targetDir=${basedir}/src/${deploy}
    if [ ! -d "$targetDir" ]; then
        echo "不存在服务---${deploy}"
        exit
    fi
    cd $targetDir

    # 删除原来dist目录
    rm -rf dist

    # 重建dist目录
    distDir=${basedir}/src/${deploy}/dist
    if [ ! -d "$distDir" ]; then
        mkdir -p $distDir
    fi

    # 拷贝文件进dist目录
    copyToDist

    # 拷贝dist目录中配置文件参数（服务器ip、文件路径等)
    replaceConfigMac ${env}

    # 编译go
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/server_upload --tags "consul" main.go

    cd dist
    tar -czf dist.tar.gz *

    # 创建服务端文件夹
    ssh root@${serverhost} "mkdir -p ${serverdir}/${deploy}/logs"

    scp -r dist.tar.gz root@${serverhost}:${serverdir}/${deploy}
    rm dist.tar.gz
    rm server_upload
    cd ../

    ssh root@${serverhost} "
cd ${serverdir}/${deploy} && \
tar -xzf dist.tar.gz && \
rm -f dist.tar.gz && \
mkdir -p logs && \
docker-compose -p ${projectname} stop && \
mv server_upload ${service} && \
docker-compose -p ${projectname} up -d && \
docker ps|grep ${deploy} && \
docker logs jdcj_${deploy}_1 "
    echo '部署完成'
}

deployglibs() {
    cd ${basedir}/src/
    tar -czf glibs.tar.gz glibs
    scp -r glibs.tar.gz root@${serverhost}:${serverdir}
    rm glibs.tar.gz
    ssh root@${serverhost} "
cd ${serverdir} && \
tar -xzf glibs.tar.gz && \
rm -f glibs.tar.gz"
    echo '部署完成'
}

deployStatic() {
    cd ${basedir}/src/gateway/static
    tar -czf glibs.tar.gz *
    scp -r glibs.tar.gz root@${serverhost}:${serverdir}/gateway/static
    rm glibs.tar.gz
    ssh root@${serverhost} "
cd ${serverdir}/gateway/static && \
tar -xzf glibs.tar.gz && \
rm -f glibs.tar.gz"
    echo '部署完成'
}

case "$service" in
"glibs")
    deployglibs
    ;;
"static")
    deployStatic
    ;;
"all")
    . ./deploy/all_config.sh
    for service in ${services[@]}; do
        deployApp $service $env
    done
    ;;
*)
    if [ $env ]; then
        # 其他微服务
        deployApp $service $env
    fi

    ;;
esac

# deployGit() {
#     if [ ! $1 ]; then
#         commitMessage="默认提交备注"
#     else
#         commitMessage=$1
#     fi
#     git add -A
#     git commit -m $commitMessage -a
#     git pull
#     git push
# }

#!/bin/bash
type=$1
env=$2

serverDir=/opt/server
projectname=kratos
service=api

case "$type" in
"git")
    git add -A
    git commit -m '提交代码' -a
    git pull
    git push
    ;;
"deploy")
    git pull

    go build -o cmd/main cmd/main.go
    chmod +x cmd/main
    ./cmd/main -conf ./configs
    ;;
"docker")
    cd $serverDir

    git pull

    #发布目录
    mkdir -p dist

    #复制
    # 配置目录
    cp -r configs dist/
    #docker-compose
    cp -r deployments/docker-compose.yml dist/

    #修改配置

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/server_upload cmd/main.go
    chmod +x dist/server_upload
    
    cd dist

    docker-compose -p ${projectname} stop
    mv server_upload ${service}
    docker-compose -p ${projectname} up -d
    docker ps | grep ${service}
    docker logs ${projectname}_${service}_1

    ;;
"run")
    go run cmd/main.go -conf configs
    ;;
*)
    echo "error"
    ;;
esac
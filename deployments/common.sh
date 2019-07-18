# 拷贝文件进dist目录
copyToDist() {
    # 拷贝配置文件目录
    cp -r configs dist/

    if [ -f "./configs/app.example.yaml" ]; then
        cp configs/app.example.yaml dist/configs/app.yaml
        cp configs/conf.example.yaml dist/configs/conf.yaml
    fi

    # 拷贝docker-compose.yml
    cp ${basedir}/docker/docker-compose.yml dist/
}

# 替换配置文件的参数
replaceConfig() {
    if [ "$1" == "pro" ]; then
        local serverhost=172.17.101.202
        # local dbname="dev_wxgame"
        local dbname="dev_wxgame"
        local redisPort="2"
    else
        local serverhost=172.17.101.205
        local dbname="wxgame_go"
        local redisPort="8"
    fi

    #conf.yaml
    sed -i "s/0\.0\.0\.0/${serverhost}/g" dist/configs/conf.yaml
    sed -i "s/127\.0\.0\.1/${serverhost}/g" dist/configs/conf.yaml
    sed -i "s/101\.200\.41\.141/${serverhost}/g" dist/configs/conf.yaml
    sed -i "s/39\.96\.187\.72/${serverhost}/g" dist/configs/conf.yaml
    sed -i "s/env: [a-z]*$/env: ${env}/g" dist/configs/conf.yaml
    sed -i "s/wxgame_go/${dbname}/g" dist/configs/conf.yaml
    sed -i "s|dbNum: 8|dbNum: ${redisPort}|g" dist/configs/conf.yaml

    # 数据库ip

    #app.yaml
    sed -i "s/..\/glibs/\/glibs/g" dist/configs/app.yaml
    sed -i "s/0\.0\.0\.0/${serverhost}/g" dist/configs/app.yaml
    sed -i "s/127\.0\.0\.1/${serverhost}/g" dist/configs/app.yaml

    # 替换app.yaml服务端口
    # servicePort=${portMap["$service"]}
    # sed -i "s/server_addr: ${serverhost}:\d+/server_addr: ${serverhost}:${servicePort}/g" dist/configs/app.yaml
    # sed -i "s/tcp@${serverhost}:\d+/tcp@${serverhost}:${servicePort}/g" dist/configs/app.yaml

    # 替换docker-compose.yml
    sed -i "s/service_name/${service}/g" dist/docker-compose.yml

    #TODO windows文件路径替换
    sed -i "s|path: D:.*|path: /glibs/config/data/|" dist/configs/app.yaml
    sed -i "s|filename: D:.*|filename: ./logs/info.log|" dist/configs/conf.yaml
    sed -i "s|8893|9888|" dist/configs/conf.yaml
}

replaceConfigMac(){
    if [ "$1" == "pro" ]; then
        local serverhost=172.17.101.202
        # local dbname="dev_wxgame"
        local dbname="dev_wxgame"
        local redisPort="2"
    else
        local serverhost=172.17.101.205
        local dbname="wxgame_go"
        local redisPort="8"
    fi

    #conf.yaml
    sed -i "" "s/0\.0\.0\.0/${serverhost}/g" dist/configs/conf.yaml
    sed -i "" "s/127\.0\.0\.1/${serverhost}/g" dist/configs/conf.yaml
    sed -i "" "s/101\.200\.41\.141/${serverhost}/g" dist/configs/conf.yaml
    sed -i "" "s/39\.96\.187\.72/${serverhost}/g" dist/configs/conf.yaml
    sed -i "" "s/env: [a-z]*$/env: ${env}/g" dist/configs/conf.yaml
    sed -i "" "s/wxgame_go/${dbname}/g" dist/configs/conf.yaml
    sed -i "" "s|dbNum: 8|dbNum: ${redisPort}|g" dist/configs/conf.yaml

    # 数据库ip

    #app.yaml
    sed -i "" "s/..\/glibs/\/glibs/g" dist/configs/app.yaml
    sed -i "" "s/0\.0\.0\.0/${serverhost}/g" dist/configs/app.yaml
    sed -i "" "s/127\.0\.0\.1/${serverhost}/g" dist/configs/app.yaml

    # 替换app.yaml服务端口
    # servicePort=${portMap["$service"]}
    # sed -i "" "s/server_addr: ${serverhost}:\d+/server_addr: ${serverhost}:${servicePort}/g" dist/configs/app.yaml
    # sed -i "" "s/tcp@${serverhost}:\d+/tcp@${serverhost}:${servicePort}/g" dist/configs/app.yaml

    # 替换docker-compose.yml
    sed -i "" "s/service_name/${service}/g" dist/docker-compose.yml
    sed -i "" "s|../../glibs|../glibs|g" dist/docker-compose.yml

    #TODO windows文件路径替换
    sed -i "" "s|path: D:.*|path: /glibs/config/data/|" dist/configs/app.yaml
    sed -i "" "s|filename: D:.*|filename: ./logs/info.log|" dist/configs/conf.yaml
    sed -i "" "s|8893|9888|" dist/configs/conf.yaml
}

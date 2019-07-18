#!/bin/bash
type=$1

case "$type" in
"git")
    git add -A
    git commit -m '提交代码' -a
    git pull
    git push
    ;;
"run")
    git pull

    go build -o cmd/main cmd/main.go
    chmod +x cmd/main
    ./cmd/main -conf ../configs
    ;;
"deploy")
    go run cmd/main.go -conf configs
    ;;
*)
    echo "error"
    ;;
esac
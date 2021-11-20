#!bin/bash

# 
# PostgreSQLのイメージをビルドする
# 

# -e:エラー時に停止, -u:未定義変数を検知時に停止
set -eu

buildContext=$(cd $(dirname $0); pwd)
imageTag="daikideal/postgres:14.0-alpine"

docker build -t $imageTag $buildContext

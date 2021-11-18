# !bin/bash

#
# データベースの(クリーンと)セットアップを行う
# NOTE: run.sh でコンテナを実行していることが前提
#

# -e:エラー時に停止, -u:未定義変数を検知時に停止
set -eu

# コンテナ設定
containerName="chapter6-db"
destPath="/tmp"

docker cp ./*.sql $containerName:$destPath

docker exec $containerName bash -c 
  "su postgres &&
   createuser -P -d gwp &&
   createdb gwp &&
   psql -U gwp -f tmp/setup.sql -d gwp"

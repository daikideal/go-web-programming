#!bin/bash

# 
# ビルドしたPostgreSQLのイメージからコンテナを実行する
# 

# -e:エラー時に停止, -u:未定義変数を検知時に停止
set -eu

# コンテナ設定
imageTag="daikideal/postgres:14.0-alpine"
containerName="chapter7-db"
volume="/var/lib/postgresql/data"

# ログイン情報
password="password"

docker run -d --rm \
      --name $containerName \
      --mount type=volume,src=postgres-data,dst=$volume \
      -e POSTGRES_PASSWORD=$password \
      -p 5432:5432 \
      $imageTag

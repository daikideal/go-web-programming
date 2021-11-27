# !bin/bash

#
# データベースの(クリーンと)セットアップを行う
#

# -e:エラー時に停止, -u:未定義変数を検知時に停止
set -eu

# コンテナ設定
containerName="chapter7-db"
destPath="/tmp"

# このスクリプトを実行したディレクトリの絶対パス
# この直下にsetup.shが存在することが期待される
workDir=$(pwd)

if [ -e $workDir/setup.sql ]; then
  # SQLファイルをコンテナにコピーする
  docker cp $workDir/*.sql $containerName:$destPath

  # SQLファイルを実行してデータベースをセットアップする
  command="su postgres &&
          createuser -P -d gwp &&
          createdb gwp &&
          psql -U gwp -f $destPath/setup.sql -d gwp"
  docker exec $containerName bash -c $command
else
  echo "abort! setup files are not found."
fi

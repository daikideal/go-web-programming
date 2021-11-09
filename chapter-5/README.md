## テンプレートファイルのディレクトリについての注意

ファイルパスに関して特に処理を行わない場合、実行したワーキングディレクトリに存在するファイルしか参照してくれない。
例えば、

```bash
% pwd
/Users/daikideal/workspace/go-web-programming/chapter-5
```

ここで`go run simple/server.go`を実行すると、
サーバは起動するが、localhost:8080/process にアクセスした時`runtime error: invalid memory address or nil pointer dereference`というエラーで落ちる。

`template.ParseFiles()`が返しているエラーを確認すればわかるが、
この時Goはtmpl.htmlを見つけられていない。

しかし、server.goとtmpl.htmlが存在するディレクトリに移動してからサーバを起動し、localhost:8080/process にアクセスすると、"Hello World!"を確認できる。

```bash
% cd simple
% go run server.go
```

この(学習用の)リポジトリのようなディレクトリ構成が特殊ゆえに起こっているだけかもしれないが、テンプレート(ファイル)を取り扱う場合には`path/filepath`などを駆使する必要があるかもしれない。

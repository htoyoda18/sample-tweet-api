# sample-tweet-api
簡易的なTwitterに近い REST APIです(未完成)

## ER図

[![Image from Gyazo](https://i.gyazo.com/c23b47b5d527088bce3e7cb265020d83.png)](https://gyazo.com/c23b47b5d527088bce3e7cb265020d83)

## Development Commands

### Setup <br>
```
make setup
```
Dockerが立ち上がって開発を始められます

また、swaggoによりswaggerが生成されます

http://localhost:8080/swagger/index.html にアクセスしていただくとAPIのドキュメントが見れます

### Test <br>
```
make test NAME="handlerのディレクトリ名" FN="メソッド名"
```
APIのテストを回すことができます

## Directory Structure
```
.
.
├── golang
│   └── src
│       ├── controller // HTTPリクエストを受け取り、UseCase を使って処理を行い、結果を返す 
│       ├── domain     // DDDのmodel、repositoryが入ってる
│       ├── infra      // DBの処理はここで行っている
│       ├── injector   // 依存関係を定義
│       ├── router     // ginのルーティング設定、ミドルウエアの処理がある
│       ├── service    // コード全体で汎用的に使う処理 (infra、repositoryに依存している)
│       ├── shared     // サービス全体で汎用的に使う処理 (どこのパッケージに依存していない)
│       ├── test       // テストのfixtuteのデフォルトファイル、ヘルパー処理などがある。
│       └── usecase    // システムのユースケースを満たす処理をしています。
└── mysql              // DBの設定などが書かれています
```

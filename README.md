# Go × Ginで作成したAPIのサンプルプログラム

## 使用技術について
### 言語
`golang`
### フレームワーク
`Gin`
### ORM
`GORM`
### ログライブラリ
`logrus`
### パッケージ管理ツール
`dep`
### その他
`Docker`

## 使用ポート
|ポート|用途|
|---|---|
|8080|APIアプリケーション(Docker)|
|3306|Mysql(Docker)|

## API定義
・現状、下記のAPIを作成しています。
* ユーザー取得API
* ユーザー作成API
* ユーザー更新API
* ユーザー削除API

## ドキュメント

#### API定義書
[API定義](https://github.com/inagacky/go_gin_sample/blob/master/docs/api/api_design.md)
#### DB定義書
[DB定義書](https://github.com/inagacky/go_gin_sample/blob/master/docs/db/database_design.md)

## プログラム使用方法
Dockerコンテナとして起動します。

### ビルド方法　
`/bin/sh build.sh` 

### 実行方法
`/bin/sh run.sh`

#### ユーザー作成API
##### 正常時
```
% curl -v -X POST -d "{ \"firstName\" : \"テスト\", \"lastName\" : \"Inagacky\", \"email\" : \"test+1@test.com\"}" -H "accept: application/json" -H "Content-Type: application/json" "http://localhost:8080/api/v1/users" | jq .
< HTTP/1.1 200
{
  "status": 200,
  "error": null,
  "result": {
    "user": {
      "id": 6,
      "created_at": "2019-04-24T11:15:59.8577435+09:00",
      "updated_at": "2019-04-24T11:15:59.8577435+09:00",
      "deleted_at": null,
      "firstName": "テスト",
      "lastName": "Inagacky",
      "email": "test+1@test.com",
      "status": 1
    }
  }
}
```
##### メールアドレス重複時
```
% curl -v -X POST -d "{ \"firstName\" : \"テスト\", \"lastName\" : \"Inagacky\", \"email\" : \"test+1@test.com\"}" -H "accept: application/json" -H "Content-Type: application/json" "http://localhost:8080/api/v1/users" | jq .
< HTTP/1.1 200
{
  "status": 500,
  "error": {
    "code": 200,
    "error": "指定されたメールアドレスのユーザーは既に存在します。"
  },
  "result": null
}
```

#### ユーザー取得API
```
curl -X GET "http://localhost:8080/api/v1/users/6" | jq .
< HTTP/1.1 200
{
  "status": 200,
  "error": null,
  "result": {
    "user": {
      "id": 6,
      "created_at": "2019-04-24T02:16:00Z",
      "updated_at": "2019-04-24T02:16:00Z",
      "deleted_at": null,
      "firstName": "テスト",
      "lastName": "Inagacky",
      "email": "test+1@test.com",
      "status": 1
    }
  }
}
```

#### ユーザー更新API
```
% curl -v -X PUT -d "{ \"firstName\" : \"TEST_FIRST_NAME\", \"lastName\" : \"TEST_LAST_NAME\", \"email\" : \"test+3@gmail.com\" }" -H "accept: application/json" -H "Content-Type: application/json" "http://localhost:8080/api/v1/users/6" | jq .
< HTTP/1.1 200
{
  "status": 200,
  "error": null,
  "result": {
    "user": {
      "id": 6,
      "created_at": "2019-04-24T02:16:00Z",
      "updated_at": "2019-04-24T11:20:50.7426184+09:00",
      "deleted_at": null,
      "firstName": "TEST_FIRST_NAME",
      "lastName": "TEST_LAST_NAME",
      "email": "test+3@gmail.com",
      "status": 1
    }
  }
}

```

#### ユーザー削除API
```
% curl -X DELETE "http://localhost:8080/api/v1/users/6" | jq .
{
  "status": 200,
  "error": null,
  "result": {
    "user": {
      "id": 6,
      "created_at": "2019-04-24T02:16:00Z",
      "updated_at": "2019-04-24T11:21:49.064395+09:00",
      "deleted_at": "2019-04-24T11:21:49.0638962+09:00",
      "firstName": "TEST_FIRST_NAME",
      "lastName": "TEST_LAST_NAME",
      "email": "test+3@gmail.com",
      "status": 9
    }
  }
}
```
##### 確認のため再取得
```
curl -X GET "http://localhost:8080/api/v1/users/6" | jq .
< HTTP/1.1 200
{
  "status": 200,
  "error": null,
  "result": {
    "user": null
  }
}
```

## DB
DBはmysql8.0を使用しています。  
Dockerコンテナとして起動します。
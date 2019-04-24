# API定義

## ユーザー作成API

### エンドポイント  
・ http://localhost:8080/api/v1/users
### メソッド
・POST
### 用途
・ ユーザーの作成
### リクエストパラメータ

| 項目 | 名称 | 型 | 必須 | 例 | 備考 |  
|:---|:---|:---|:---|:---|:---|  
|firstName| 苗字| String | ○ | - | - |
|lastName| 名前 | String | ○ | - | - |
|email| メールアドレス| String | ○ | sample@gmail.com | - |

```
{
  "firstName": "TEST_FIRST_NAME",
  "lastName": "TEST_LAST_NAME",
  "email": "sample@gmail.com",
}
```
### レスポンス

```
{
  "status": 200,
  "error": null,
  "result": {
    "user": {
      "id": 1,
      "created_at": "2019-04-24T02:16:00Z",
      "updated_at": "2019-04-24T11:20:50.7426184+09:00",
      "deleted_at": null,
      "firstName": "TEST_FIRST_NAME",
      "lastName": "TEST_LAST_NAME",
      "email": "sample@gmail.com",
      "status": 1
    }
  }
}

```

---

## ユーザー更新

### エンドポイント  
・ http://localhost:8080/api/v1/users/${userId}
### メソッド
・PUT
### 用途
・ユーザーの更新
### リクエストパラメータ

| 項目 | 名称 | 型 | 必須 | 例 | 備考 |  
|:---|:---|:---|:---|:---|:---|  
|userId| ユーザーID| Int | ○ | - | - |
|firstName| 苗字| String | ○ | - | - |
|lastName| 名前 | String | ○ | - | - |
|email| メールアドレス| String | ○ | sample@gmail.com | - |

```
{
  "firstName": "TEST_FIRST_NAME",
  "lastName": "TEST_LAST_NAME",
  "email": "sample@gmail.com",
}
```
### レスポンス

```
{
  "status": 200,
  "error": null,
  "result": {
    "user": {
      "id": 1,
      "created_at": "2019-04-24T02:16:00Z",
      "updated_at": "2019-04-24T11:20:50.7426184+09:00",
      "deleted_at": null,
      "firstName": "TEST_FIRST_NAME",
      "lastName": "TEST_LAST_NAME",
      "email": "sample@gmail.com",
      "status": 1
    }
  }
}

```

---

## ユーザー取得API

### エンドポイント  
・ http://localhost:8080/api/v1.0/users/${userId}
### メソッド
・GET
### 用途
・ ユーザーの取得
### リクエストパラメータ

| 項目 | 名称 | 型 | 必須 | 例 | 備考 |  
|:---|:---|:---|:---|:---|:---|  
|userId| ユーザーID| Int | ○ | - | - |


### レスポンス

```
{
  "status": 200,
  "error": null,
  "result": {
    "user": {
      "id": 1,
      "created_at": "2019-04-24T02:16:00Z",
      "updated_at": "2019-04-24T11:20:50.7426184+09:00",
      "deleted_at": null,
      "firstName": "TEST_FIRST_NAME",
      "lastName": "TEST_LAST_NAME",
      "email": "sample@gmail.com",
      "status": 1
    }
  }
}
```

---

## ユーザー削除API

### エンドポイント  
・ http://localhost:8080/api/v1.0/users/${userId}
### メソッド
・DELETE
### 用途
・ ユーザーの削除
### リクエストパラメータ

| 項目 | 名称 | 型 | 必須 | 例 | 備考 |  
|:---|:---|:---|:---|:---|:---|  
|userId| ユーザーID| Int | ○ | - | URLに含む |

### レスポンス

```
{
  "status": 200,
  "error": null,
  "result": {
    "user": {
      "id": 1,
      "created_at": "2019-04-24T02:16:00Z",
      "updated_at": "2019-04-24T11:20:50.7426184+09:00",
      "deleted_at": null,
      "firstName": "TEST_FIRST_NAME",
      "lastName": "TEST_LAST_NAME",
      "email": "sample@gmail.com",
      "status": 9
    }
  }
}
```
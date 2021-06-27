# Golang-RESTful-API

## 💡 Overview
- GolangでAPIを作成してDockerで動かす
- フレームワーク：echo（Golang）
  - [https://echo.labstack.com](https://echo.labstack.com)
  - [https://github.com/labstack/echo](https://github.com/labstack/echo)

## 🌍 Requirements
| 言語/フレームワーク | バージョン |
| :---: | :---: |
| Golang | 1.15 |
| Docker | 20.10.5 |
| docker-compose | 1.29.0 |
| MySQL | 8.0.23 |

## 🚀 Usage  
```
### 起動
$ make up

### appコンテナに入る
$ make app/api

### dbコンテナに入る
$ make app/db

### dbコンテナに入る + MySQL接続
$ make mysql

### ログ
$ make logs

### Fetch All
$ curl localhost:9090/users

### Fetch where ID
$ curl localhost:9090/users/1

### 確認
=== * 起動するDockerコンテナ * ===
$ docker ps
CONTAINER ID   IMAGE                COMMAND                  CREATED         STATUS         PORTS                               NAMES
d20bdfdcbb9e   go-restful-api_db    "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes   33060/tcp, 0.0.0.0:3307->3306/tcp   go-restful-api_db
86b3d84ae452   go-restful-api/app   "./main"                 2 minutes ago   Up 2 minutes   0.0.0.0:9090->8080/tcp              go-restful-api_app

=== * 作成されるDockerイメージ * ===
$ docker images
REPOSITORY                           TAG        IMAGE ID       CREATED         SIZE
go-restful-api/app                   latest     dfd2fc6df75e   4 minutes ago   14MB
go-restful-api_db                    latest     8c6bc13c1a11   3 days ago      546MB

=== * 作成されるDockerネットワーク * ===
$ docker network ls
NETWORK ID     NAME                  DRIVER    SCOPE
f56484cfd19a   go-restful-api_link   bridge    local
```

## 💣 Other
```
### 不要（`<none>`）imageのみを削除
$ docker rmi $(docker images -f "dangling=true" -q)

### 停止 & 除去
$ make down

### 全てのイメージ、ネットワーク、コンテナを削除
$ make down/all

### 全てのボリュームを削除
$ make down/vol
```

## 📝 Reference
- [GolangでシンプルなRESTful APIを作ってみた](https://qiita.com/yuuulf/items/b464735dfb3486d248bf)
- [Goで作ったWebアプリをDockerで動かす](https://qiita.com/yuuulf/items/b678b00972d60c7d63dd)
- [MySQL8.0新機能 (caching_sha2_password 認証プラグイン)](https://www.s-style.co.jp/blog/2018/05/1807/)


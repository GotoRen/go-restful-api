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
### キャッシュ無しビルド & 起動
$ docker-compose build --no-cache && docker-compose up -d

### 起動
$ docker-compose up -d

### 確認
=== * 起動するDockerコンテナ * ===
$ docker ps
CONTAINER ID   IMAGE                COMMAND                  CREATED              STATUS              PORTS                               NAMES
11444393d32a   go-restful-api/db    "docker-entrypoint.s…"   About a minute ago   Up About a minute   33060/tcp, 0.0.0.0:3307->3306/tcp   go-restful-api_db
ac3e965540e8   go-restful-api/app   "./main"                 About a minute ago   Up About a minute   0.0.0.0:9090->8080/tcp              go-restful-api_app

=== * 作成されるDockerイメージ * ===
$ docker images
REPOSITORY                           TAG        IMAGE ID       CREATED              SIZE
go-restful-api/app                   latest     dc41b5111da1   About a minute ago   14MB
go-restful-api/db                    latest     8c6bc13c1a11   3 days ago           546MB

=== * 作成されるDockerネットワーク * ===
$ docker network ls
NETWORK ID     NAME                    DRIVER    SCOPE
93c6a7a73d3b   go-restful-api_link     bridge    local

### ログ
$ docker-compose logs -f

### Fetch All
$ curl localhost:9090/users

### Fetch where ID
$ curl localhost:9090/users/1
```

## 💣 Other
```
### 不要（`<none>`）イメージの削除
$ docker rmi $(docker images -f "dangling=true" -q)

### 停止 & 除去
$ docker-compose down
$ docker rmi [イメージID]
```

## 📝 Reference
- [GolangでシンプルなRESTful APIを作ってみた](https://qiita.com/yuuulf/items/b464735dfb3486d248bf)
- [Goで作ったWebアプリをDockerで動かす](https://qiita.com/yuuulf/items/b678b00972d60c7d63dd)
- [MySQL8.0新機能 (caching_sha2_password 認証プラグイン)](https://www.s-style.co.jp/blog/2018/05/1807/)


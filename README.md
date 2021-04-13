# Golang-RESTful-API

## 💡 Overview
- GolangでAPIを作成してDockerで動かす
- フレームワーク：echo（Golang）
  - [https://echo.labstack.com](https://echo.labstack.com)
  - [https://github.com/labstack/echo](https://github.com/labstack/echo)

## 🌍 Environment
- 環境

| 言語/フレームワーク | バージョン |
| :---: | :---: |
| go | go version go1.15 darwin/amd64 |
| MySQL | mysql  Ver 8.0.23 for Linux on x86_64 (MySQL Community Server - GPL)|

- __OS__：macOS Catalina ver. 10.15.7
  - __Golang__：go version go1.15.4 darwin/amd6
  - __MySQL for Docker__：Server version: 8.0.22 MySQL Community Server - GPL
- __OS__：Linux 20.04.1 LTS (Focal Fossa) ⬅︎ 仮想マシン
  - __Golang__：go version go1.15.6 linux/amd6
  - __MySQL for Docker__：Server version: 8.0.22 MySQL Community Server - GPL
- __OS__：Linux 20.04.1 LTS (Focal Fossa)
  - __Golang__：go version go1.13.8 linux/amd6
  - __MySQL for Docker__：Server version: 8.0.22 MySQL Community Server - GPL

## 🚀 Usage  
- 必要な権限と構成情報を設定した後、実行する  
```
### コンテナ起動 & 実行
$ docker-compose up -d

### 確認
=== * 起動するDockerコンテナ * ===
$ docker ps
CONTAINER ID   IMAGE                COMMAND                  CREATED         STATUS         PORTS                               NAMES
f95985909759   go-restful-api/db    "docker-entrypoint.s…"   3 minutes ago   Up 2 minutes   33060/tcp, 0.0.0.0:3307->3306/tcp   go-restful-api_db
e2c1a557162e   go-restful-api/app   "./main"                 3 minutes ago   Up 2 minutes   0.0.0.0:9090->8080/tcp              go-restful-api_app

=== * 作成されるDockerイメージ * ===
$ docker images
REPOSITORY                           TAG        IMAGE ID       CREATED         SIZE
go-restful-api/app                   latest     9b0353e1fadb   3 minutes ago   14MB
go-restful-api/db                    latest     8c6bc13c1a11   3 days ago      546MB

=== * 作成されるDockerネットワーク * ===
$ docker network ls
NETWORK ID     NAME                    DRIVER    SCOPE
116d3317efc5   go-restful-api_link     bridge    local

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


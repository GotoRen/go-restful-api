# Golang-RESTful-API

## 💡 Overview
- Golangでシンプルな RESTful API を作成してDockerで動かす
- フレームワーク：[echo（Golang）](https://echo.labstack.com/)

## 🌍 Environment
- 動作させた環境
  - __OS__：macOS Catalina ver. 10.15.7
    -  __Golang__：go version go1.15.4 darwin/amd6
    - __MySQL for Docker__：Server version: 8.0.22 MySQL Community Server - GPL
  - __OS__：Linux 20.04.1 LTS (Focal Fossa) ⬅︎ 仮想マシン
    -  __Golang__：go version go1.15.6 linux/amd6
    - __MySQL for Docker__：Server version: 8.0.22 MySQL Community Server - GPL
  - __OS__：Linux 20.04.1 LTS (Focal Fossa)
    -  __Golang__：go version go1.13.8 linux/amd6
    - __MySQL for Docker__：Server version: 8.0.22 MySQL Community Server - GPL

## ⚡ Configure
- 権限
  - `$ sudo chmod -R 775 go-restful-api`
- MySQL8系をMySQL5系クライアントと合わせる場合
  - `docker-compose.yml`に以下を追加
    ```yml
    db:
      image: mysql
  
      ## ここに追加する
      command: --default-authentication-plugin=caching_sha2_password
  
      restart: always
    ```
  - `command`：MySQL8系から認証系の中身が変更になったため、5系のクライアントと合わせるためのオプション
  - デフォルトに設定されている認証プラグインを確認
    - `mysql> show variables like 'default_authentication_plugin';`

## 🚀 Usage  
- 必要な権限と構成情報を設定した後、実行する  
```
### コンテナ起動 & 実行
$ docker-compose up -d --build

### 確認
=== * 起動するDockerコンテナ * ===
$ docker ps
CONTAINER ID   IMAGE                COMMAND                  CREATED       STATUS       PORTS                               NAMES
bfb04d890070   mysql                "docker-entrypoint.s…"   2 hours ago   Up 2 hours   33060/tcp, 0.0.0.0:3307->3306/tcp   go-restful-api_db_1
6b75c3927f60   go-restful-api_api   "/bin/hello"             2 hours ago   Up 2 hours   0.0.0.0:1323->1323/tcp              go-restful-api_api_1

=== * 作成されるDockerイメージ * ===
$ docker images
go-restful-api_api   latest    264c29164612   About a minute ago   11.3MB
<none>               <none>    1b0971fd8812   About a minute ago   421MB

=== * 作成されるDockerネットワーク * ===
$ docker network ls
b6f1e3ca4819   go-restful-api_default   bridge    local

### ログ
$ docker-compose logs -f

### Fetch All
$ curl localhost:1323/users

### Fetch where ID
$ curl localhost:1323/users/1
```

## 💣 Other
- 不要（`<none>`）イメージの削除
  - `$ docker rmi $(docker images -f "dangling=true" -q)`
- ※ コンテナを削除
  - `$ docker-compose down`
  - `$ docker rmi [イメージID]`
  - `$ docker network rm [ネットワークID]`


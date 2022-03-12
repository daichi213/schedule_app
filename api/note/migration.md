# Go製マイグレーションツールまとめ

## golang-migration

go製のマイグレーションツールで、ginなどでwebアプリやAPI開発時にSQLサーバーに対してtableを作成することができ、テーブル作成のマイグレーションをファイルとして残すことができる開発にとても便利なツール。

### インストール

dockerを使用して構築する場合は、Dockerfileに以下の記載を追加する。主として参考にしたサイトは[こちら](https://dev.classmethod.jp/articles/db-migrate-with-golang-migrate/)

```Dockerfile
# tarコマンドまでを実行した時点でmigrate.linux-amd64の実行ファイルがインストールされる。その実行ファイルをmvでbinディレクトリへ移動させる
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 /usr/bin/migrate
# 上記コマンドでインストールした実行ファイルのパスを設定する。この設定によってコンテナのグローバル環境でmigrateコマンドが実行可能になる。
ENV PATH $PATH:/usr/bin/migrate
```

### How to

使用方法を参考にしたサイトはインストール時と[同じサイト](https://dev.classmethod.jp/articles/db-migrate-with-golang-migrate/)

```bash
# DSN(Data Source Name):
# username:password@protocol(address)/dbname?param=value
$ export MYSQL_URL="mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}"
# テーブルに適用しているマイグレーションファイルのバージョンを表示する
$ migrate -database ${MYSQL_URL} -path ./migrations version
# マイグレーションファイルの作成
$ migrate create -ext sql -dir migrations/example1 -seq create_users_table
$ migrate create -ext sql -dir migrations/example1 -seq add_mood_to_users

# マイグレーションを元にした実行
$ migrate -database ${MYSQL_URL} -path migrations/example1 up
1/u create_users_table (29.5482ms)
# 切り戻し
$ migrate -database ${MYSQL_URL} -path migrations/example1 down
```

以下のUPファイルにテーブルを作成する際のマイグレーションファイルの記載例を示す。UPファイルにはテーブルに加えたい変更内容を記載し、DOWNファイルにはUPファイルに記載した処理内容の切り戻し処理を記載する。なお、ファイルにはSQLを用いて処理を記載する。

```sql
-- UPファイル
CREATE TABLE IF NOT EXISTS todo(
    ID serial PRIMARY KEY,
    Title VARCHAR (50) NOT NULL,
    Content VARCHAR (250) NOT NULL,
    Status BIT NOT NULL
    -- email VARCHAR (300) UNIQUE NOT NULL
);

-- DOWNファイル
DROP TABLE IF EXISTS todo;
```

>●つまづきポイント
APIコンテナからDBコンテナへ接続する際にpingが通らない現象が発生した。
>>原因はdocker-compose.ymlで設定したdb-network内へapiコンテナを所属させていなかったことが原因だった。そのため、APIコンテナとDBコンテナはネットワークとしては隔離された状態だった。composeに定義したネットワークが異なる場合はGatewayのアドレスに接続する必要があった。

## sql-migration

golang-migrationと異なり、設定ファイルで接続先情報をまとめられたり、実行コマンドがシンプルだったりと使いやすいマイグレーションツール。

### インストール

1. 以下コマンドを実行する。これだけで、コンテナにパスも通る。

```bash
$ go get github.com/rubenv/sql-migrate/...
```

2. 設定ファイルを作成する。

```yml
test:
  dialect: mysql
  datasource: mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_TEST_DATABASE}
  dir: migrations/test
development:
  dialect: mysql
  datasource: mysql://root:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}
  dir: migrations/development
```

### How to

```bash
# dbconfのdir項目を参照してファイルを生成してくれるので、先に手動でディレクトリを作成しておく
$ sql-migrate new -env="test" create_user_table
# マイグレーションの実行
$ sql-migrate up
# マイグレーションの実行状況の確認
$ sql-migrate status
# 直前の操作を再実行する
$ sql-migrate redo

# HELP
# 全般
$ sql-migrate --help
# upに関するHELP
$ sql-migrate up --help
```
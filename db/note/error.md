# MySQLエラー集

## 日本語対応

### API側から見たエラー

API側からDBへ日本語のレコードを保存すると以下のように、Ginが例外を起こしてコンテナが落ちてしまう。

```log
api_1    | 2021/12/18 04:25:06 title:test title content:これはテストです？ええ、テストです！
api_1    | 2021/12/18 04:25:06 status:
api_1    | 2021/12/18 04:25:06 input is read.
api_1    | 
api_1    | 2021/12/18 04:25:06 /go/src/api/model.todo.go:39 Error 1366: Incorrect string value: '\xE3\x81\x93\xE3\x82\x8C...' for column 'content' at row 1
api_1    | [2.383ms] [rows:0] INSERT INTO `todos` (`created_at`,`updated_at`,`deleted_at`,`title`,`content`,`status`,`user_id`) VALUES ('2021-12-18 04:25:06.764','2021-12-18 04:25:06.764',NULL,'test title','これはテストです？ええ、テストです！',0,0)
api_1    | 2021/12/18 04:25:06 Could not create: Error 1366: Incorrect string value: '\xE3\x81\x93\xE3\x82\x8C...' for column 'content' at row 1
api_1    | exit status 1
todoapp_api_1 exited with code 
```


### DB側から見た挙動

/var/lib/mysql/General.logでログを確認すると以下のようなログが吐かれていた。

```log
2021-12-18T04:25:06.761154Z         3 Connect   root@192.168.200.1 on todo_test using TCP/IP
2021-12-18T04:25:06.761807Z         3 Query     SET NAMES utf8
2021-12-18T04:25:06.762343Z         3 Query     SELECT VERSION()
2021-12-18T04:25:06.763281Z         3 Query     START TRANSACTION
2021-12-18T04:25:06.765409Z         3 Prepare   INSERT INTO `todos` (`created_at`,`updated_at`,`deleted_at`,`title`,`content`,`status`,`user_id`) VALUES (?,?,?,?,?,?,?)
2021-12-18T04:25:06.765671Z         3 Execute   INSERT INTO `todos` (`created_at`,`updated_at`,`deleted_at`,`title`,`content`,`status`,`user_id`) VALUES ('2021-12-18 04:25:06.764','2021-12-18 04:25:06.764',NULL,'test title','これはテストです？ええ、テストです！',0,0)
2021-12-18T04:25:06.766274Z         3 Close stmt
2021-12-18T04:25:06.766367Z         3 Query     ROLLBACK
```

### DBの日本語対応

#### OSレベルの日本語対応

上記のようなエラーの原因はDBのOSレベルから日本語対応していないことが原因のようだった。[こちらのページ](https://www.naokilog.com/2017/11/03/docker-%E3%81%AE%E6%97%A5%E6%9C%AC%E8%AA%9E%E8%A8%AD%E5%AE%9A/)に従ってDockerfileにてイメージをセットアップすることで、シェル上に日本語を入力できるようになった。

```Dockerfile
RUN apt-get update \
    && apt-get install -y locales \
    && locale-gen ja_JP.UTF-8 \
    && localedef -f UTF-8 -i ja_JP ja_JP
```

#### DBレベルの日本語対応

以下のクエリを実行することで、DBごとにCHARSETを設定することができる。

```sql
CREATE DATABASE IF NOT EXISTS todo_test;
ALTER DATABASE todo_test CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

### 原因

今回テーブル作成はGinのsql-migrationを使用して行っていたが、テーブルを以下のCHARSETをlatin1に設定していたことが原因だった。mysql全体の設定やdatabase自体の設定をいじっても解決しなかったが、テーブルをピンポイントで設定してしまっていた・・・
この部分を「latin1」から「utf8mb4」へ変更して問題なく解決した。

```sql
CREATE TABLE `todos` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext,
  `content` longtext,
  `status` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_todos_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;
```

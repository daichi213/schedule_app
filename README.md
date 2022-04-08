# 環境構築

## Build手順まとめ

```sh
$ docker-compose run --rm front npx create-next-app --ts next_app 
$ docker-compose build
$ docker-compose exec front sh
$ tsc --init
```


## frontについて

### 初期設定

1. frontディレクトリを作成後にfront用のDockerfileを配置

2. 以下コマンドを順番に実行する

```sh
$ docker-compose run --rm front npx create-next-app todo_app 
$ docker-compose build
# front
$ docker-compose run front sh
$ tsc --init
$ npm install --save-dev jest jest-dom @types/jest ts-jest @testing-library/dom @testing-library/jest-dom @testing-library/react babel-jest identity-obj-proxy react-test-renderer
# api
```

3. tsconfig.jsonの設定を以下とする
{ 
    "compilerOptions": {
    "jsx": "react-jsx",
    }
}

### jest, React Testing Libraryの導入

jestとRTLの導入に関しては[Nextjsの公式ページ](https://nextjs.org/docs/testing)を参考にした。

1. nextjsのプロジェクトルートディレクトリにjest.config.jsとjest.setup.jsを配置する。

```js
// jest.config.js

module.exports = {
    setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
    collectCoverageFrom: [
      '**/*.{js,jsx,ts,tsx}',
      '!**/*.d.ts',
      '!**/node_modules/**',
    ],
    moduleNameMapper: {
      /* Handle CSS imports (with CSS modules)
      https://jestjs.io/docs/webpack#mocking-css-modules */
      '^.+\\.module\\.(css|sass|scss)$': 'identity-obj-proxy',
  
      // Handle CSS imports (without CSS modules)
      '^.+\\.(css|sass|scss)$': '<rootDir>/__mocks__/styleMock.js',
  
      /* Handle image imports
      https://jestjs.io/docs/webpack#handling-static-assets */
      '^.+\\.(jpg|jpeg|png|gif|webp|svg)$': '<rootDir>/__mocks__/fileMock.js',
    },
    testPathIgnorePatterns: ['<rootDir>/node_modules/', '<rootDir>/.next/'],
    testEnvironment: 'jsdom',
    transform: {
      /* Use babel-jest to transpile tests with the next/babel preset
      https://jestjs.io/docs/configuration#transform-objectstring-pathtotransformer--pathtotransformer-object */
      '^.+\\.(js|jsx|ts|tsx)$': ['babel-jest', { presets: ['next/babel'] }],
    },
    transformIgnorePatterns: [
      '/node_modules/',
      '^.+\\.module\\.(css|sass|scss)$',
    ],
  }
```

以下のファイルにテストを走らせる際に毎回読み込ませたいライブラリを記載することで、各テストスィートの実行前に毎回準備してくれる。

```js
// jest.setup.js

import '@testing-library/jest-dom/extend-expect'
```

2. __mocks__, __tests__ディレクトリnode_modulesと同階層のディレクトリに作成する

3. __mocks__ディレクトリに以下ファイルと設定を追加する

```js
// __mocks__/fileMock.js

(module.exports = "test-file-stub")
```

```js
// __mocks__/styleMock.js

module.exports = {};
```

4. 以下コマンドを実行する

 ```bash
$ npm install --save-dev jest jest-dom @types/jest ts-jest @testing-library/dom @testing-library/jest-dom @testing-library/react babel-jest identity-obj-proxy react-test-renderer
 ```

5. package.jsonにtestコマンドを追加する

```json
"scripts": {
  "dev": "next dev",
  "build": "next build",
  "start": "next start",
  "test": "jest --watchAll"
}
```

### Material UIの導入

基本的に[こちらのサイト](https://maku.blog/p/s6djqw3/)を参考にした。

1. コンテナ内で以下コマンドを実行する

```bash
$ npm install @material-ui/core @material-ui/icons
```

2. NextjsでMaterialUIを使用する際は、SSRとの兼ね合いからスタイルの処理順序を制御する必要があるとのこと。そのため、以下の設定ファイルを準備する。

```

```

### 本番環境で必要なnpmモジュール

本番環境でも必要になるモジュールを以下コマンドでインストールする。node_modulesが存在するディレクトリでnpm installして導入する。

```bash
$ cd todo_app
$ npm install axios
```

### ライブラリを追加する場合

- Dcoerfileに追加するライブラリを記載してbuildする

#### 注意点



## api(gin)について

### 初期設定
    
1. apiディレクトリを作成後にfront用のDockerfileを配置

2. golangのベースイメージを元にコンテナを起動し、作業ディレクトリにて以下コマンドを実行してgo.modファイルを作成する。golangのwebフレームワークであるginを使用する場合は以下urlを追加する

```sh
# コンテナの起動
$ docker-compose run api sh
# modファイルの作成
$ go mod init ginApp
# modファイルへ必要なライブラリの追加
$go get -u github.com/gin-gonic/gin
$ touch main.go
$ go run main.go
```

以下コマンドをコンテナーへ入った後にまとめて貼り付けてまとめて実行する
```sh
go get golang.org/x/tools/cmd/godoc
go get github.com/lib/pq
go get gorm.io/driver/postgres 
go get github.com/rubenv/sql-migrate/...
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get github.com/go-delve/delve/cmd/dlv@latest
go get github.com/stretchr/testify
go get github.com/DATA-DOG/go-sqlmock
go get github.com/google/wire
go get github.com/gin-contrib/sessions
go get github.com/koron/go-dproxy
go install github.com/x-motemen/gore/cmd/gore@latest
```

3. 作成されたgo.mod, go.sumをイメージ内へ転写して `go mod download`を実行することでイメージにライブラリをインストールすることができる。

### 環境構築手順

```sh
$ docker-compose exec api sh
# Table作成
$ sql-migrate up -env="development"
# Tableが作成されているか確認する
$ sql-migrate status -env="development"
# saltを生成するためのシェル
$ sh /go/src/api/generatingSalt.sh
# saltが生成されてるか確認する
$ cat /go/src/api/salt.env
SALT="ivYD6S8xRZ0SfacCmhmcfROlBKz7VyjWmIUV"/go/src/api # 
```

### 田桐さん

```bash
curl http://localhost:8080/login --data "{ \"UserName\": \"testUser\", \"Email\": \"testUser@gin.org\",  \"Password\": \"password\"}" -H 'Content-type: application/json' -X POST -v

http -f GET localhost:8080/auth/hello "Authorization:Bearer xxxxxxxxx"  "Content-Type: application/json"
```

## 備考

### chacheの削除

mysqlのイメージを使用時にcomposeファイルでenvironmentを設定した時、一度コンテナを立てるとchacheに設定が保存されている。そのため、このようなchacheに設定が保存されている場合はchacheを削除する必要がある。chacheの削除は以下コマンドで実行することができる。

```bash
$ docker builder prune
```

## 既知の脆弱性への対策

### 認証系

#### ブルートフォース攻撃

#### レインボーテーブル攻撃


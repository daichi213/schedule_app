# ScheduleAppの開発中メモ

## PASSWORDのハッシュ化について

### Golangでのパスワードハッシュ化

Gin自体にパスワードのハッシュ化関数などはないようなので、goの標準のBcryptというパッケージを使用した。また、パスワードのハッシュ化に際しても、[レインボーテーブル攻撃などの手法が存在するようで、SALTなどを使用してパスワードの管理強度を高める必要がある](https://christina04.hatenablog.com/entry/password-hash-function)。


### SALT

SALTをハードコードするのもセキュリティの観点からよくないと思われるため、APPをbuildした際に毎回ランダムな32byte文字列でSALTを生成するようにした。

```sh
# このコマンドで36Byteのランダムな文字列を生成し、先頭に環境変数を定義するenvファイルを生成できる
echo \"`cat /dev/urandom | tr -dc 'a-z
A-Z0-9' | fold -w 36 | head -n 1 | sort | uniq`\" | 
sed -e '1iSALT=' | tr -d '\n' > salt.env 
```

## godoc

godocコマンドを使用することで、go.modに存在しているパッケージのドキュメントをブラウザからオフラインで参照することができる。
必要に応じて、実行ファイルに実行権限を付与する。

```sh
$ chmod +x /go/bin/godoc
$ godoc -http ":8080"
```
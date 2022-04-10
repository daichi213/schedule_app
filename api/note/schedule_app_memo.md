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

## レコードの挿入

レコード自体VALUES句の後ろに指定するが、この時、レコードはダブルクォーテーションではなく、シングルクォーテーションで囲むようにする。ダブルクォーテーションは列名として認識されるよう。

```sql
INSERT INTO users (created_at, username, email, password, admin_flag) VALUES (current_timestamp, 'testUser', 'test@gin.org', decode('password','escape'), 1);
```

## jwtを使用した認証

認証時にtokenをクライアント側へ発行し、その発行したtokenを使用して認可を行う方式のこと。
今回、SPAでバックエンド側をginで行うにあたり、"github.com/appleboy/gin-jwt/v2"のライブラリを使用した。

### tokenの保管先

tokenの保管には十分に注意する必要があり、[特にセキュリティを意識せずクライアントのlocalstrageに保存してしまうと、JSから簡単にtokenを盗めてしまうため、注意が必要になる。](https://tech.hicustomer.jp/posts/modern-authentication-in-hosting-spa/)

今回の保管方法については以下を候補にした。

- [Cookieを使用した認証](https://korattablog.com/2020/07/20/gin%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%9Fgo-api%E9%96%8B%E7%99%BA%E3%81%AE%E5%88%9D%E6%AD%A9%EF%BC%88cookie%E7%B7%A8%EF%BC%89/)
- sessionStrageを使用した認証
    - この方法では、クライアント側にtokenがあるため別のサービスへの認証をシームレスに行う際に役に立つ
    →田桐さんから教えてもらったメリット的にはこの方法でないと意味がないか・・・
    →[セキュアにするなら、AWSのcognitoなどを組み合わせて使用するのが良い？](https://tech.hicustomer.jp/posts/modern-authentication-in-hosting-spa/)
    - jsから簡単にアクセスできてしまうためセキュリティの担保が難しい

今回はサービスとして運用する予定はないため、セキュリティは必要最低限にし、認証についての勉強を兼ねてSessionStrageを使用したtokenの保管方法を採用する

# Party-7 2020-09-13
## 読んだところ
- 4章 Go での並行処理パターン
    - 4.12 context パッケージ
    - 4.13 まとめ
- 5章 大規模開発での並行処理
    - 5.1 エラー伝搬
    - 5.2 タイムアウトとキャンセル処理
    
## Note
いつも忘れる:
```
<-chan が読み込み専用権限
chanまたはchan<- が書き込み権限
```

- コンテキストパッケージ:
    - https://pkg.go.dev/context?tab=doc

```
// DeadlineExceeded is the error returned by Context.Err when the context's
// deadline passes.
var DeadlineExceeded error = deadlineExceededError{}

type deadlineExceededError struct{}

func (deadlineExceededError) Error() string   { return "context deadline exceeded" }
func (deadlineExceededError) Timeout() bool   { return true }
func (deadlineExceededError) Temporary() bool { return true }
```
- とても理解が進むありがたいスライド:
    - https://www.slideshare.net/JamesKirk67/using-contextcontext-in-context-166532460
    - コールグラフ (関数間の呼び出し) の各枝: Context の木構造
    - 上記のスライドの8ページに木があるぞ。
    
- locale: 情報を受け取る場所みたいな意味

- TODO: 次の文賞の部分は、実コードがあると身に入りやすそう。

> 先の例でHandleResponseがresponseパッケージに存在していたとしましょう。そしてProcessRequestがprocessという名前のパッケージに存在していたとしましょう。processパッケージ はHandleResponseを呼び出すためにresponseパッケージをインポートしなければなりませんが、 HandleResponseは processパッケージ内で定義されたアクセサ関数にアクセスする方法がありませ ん。なぜなら、process パッケージのインポートは循環参照になるかもしれないからです。Context 内 にキーを保管するために使った型は process パッケージ内のプライベートな型であるため、response パッケージがこのデータを受け取る方法はありません。

- P146: 何を保管するのに適しているかについての最も普及しているガイドは context パッケージにあるこの 少し曖昧なコメントです。
    - > コンテキスト値はプロセスや API の境界を通過するリクエストスコープでのデータに絞って使いましょう。 関数にオプションのパラメーターを渡すために使うべきではありません。
        - https://pkg.go.dev/context?tab=doc
```
Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.
```

- https://github.com/pkg/errors を読んでおくと良い！

## Playground

### Next
- 5章 5.3 ハートビート から

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

```
Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.
```

- https://github.com/pkg/errors を読んでおくと良い！

## Playground

### Next
- 5章 5.3 ハートビート から

# Party-2 2020-07-05
## 読んだところ
- 1章 並行処理入門
    - ライブロック
    - リソース枯渇
    - 1.2.5 並行処理の安全性を見極める
    - 1.3 複雑さを前にした簡潔さ
- 2章 並行性をどうモデル化するか：CSP とはなにか
    
## Note
- ケイデンス:
    - https://tsuhon.jp/column/6374
- ライブロックについて説明があるブログ:
    - https://zkro.hateblo.jp/entry/toward-livelovk
 ```
 スレッドAで、A2 ~ A5 でlock2がロックされている場合、lock1 のロックを解除して「もう一方のスレッドに譲る（一旦ロックを解除し(A3)、少し待ってから(A4) 再ロックを行う (A5) ）」行為を行っています。しかしながらスレッドBも同じことをしているため、最悪の場合お互い譲り合いをしあう状態が続き、whileループを抜け出せなくなってしまいます。
 ```
- クリティカルセクション:
    - 単一の計算資源 (リソース) に対して、複数の処理が同時期に実行されると、破綻をきたす部分のこと。
- Pi:
    - https://golang.org/pkg/math/
- マルチプレキシング:
    - 複数のソースからのデータが、同じメディア (テープ) に同時に書き込まれるプロセスのこと。
- https://qiita.com/gold-kou/items/4431d3dd41606d41732b
- Go の並行処理について:
    - https://blog.golang.org/concurrency-is-not-parallelism
    - https://hub.packtpub.com/concurrency-and-parallelism-in-golang-tutorial/
- エクスプロイト:
    - コンピュータやスマートフォンの OS、ソフトウェアなどの脆弱性を悪用して攻撃を行うプログラム
のこと。
- CSP についての説明:
    - https://ja.wikipedia.org/wiki/Communicating_Sequential_Processes
    - https://www.amazon.co.jp/Communicating-Sequential-Processes-International-Computing/dp/0131532715
    - https://en.wikipedia.org/wiki/Calculus_of_communicating_systems
- ダイクストラ氏:
    - https://ja.wikipedia.org/wiki/%E3%82%A8%E3%83%89%E3%82%AC%E3%83%BC%E3%83%BB%E3%83%80%E3%82%A4%E3%82%AF%E3%82%B9%E3%83%88%E3%83%A9
    - "構造化プログラミング" の提唱者。
    - グラフ理論の "ダイクストラ法"
- メッセージパッシング ←→ メモリアクセス同期・メモリ共有同期
- ランタイムって何？
    - プログラムの実行に必要ないろいろなもののこと。
    - Java の例:
        - Windows で Java のプログラムを利用する場合、アプリだけダウンロードしても動かすことができません。ランタイムと呼ばれる、アプリケーションを動かすための実行環境をインストールする必要があります。
    - https://eng-entrance.com/java-basic-runtime
    - プログラムを実行したときに発生するエラーをruntimeエラーと呼ぶ。

```
通信によってメモリを共有し、メモリの共有によって通信してはいけない
```

- Go 言語の初心者が見ると幸せになれる場所:
    - https://qiita.com/tenntenn/items/0e33a4959250d1a55045
- FAQ の並行処理:
    - https://golang.org/doc/faq#Concurrency
- メッセージパッシング:
    - CSP 形式に代表される
- メモリ共有同期:
    - ロック機構など。例: sync.Mutex

### Next
- 2020-07-DD 10:00 - 12:00
- 3章 Go における並行処理の構成要素 から


## その他
### 自分で並行処理を書いて理解したい

- Youtube を見つつ、触ってみる:
    - https://github.com/adityamenon/Google-IO_2012_Go-Concurrency-Patterns
- WaiGroup Tutorial:
    - https://tutorialedge.net/golang/go-waitgroup-tutorial/
- Go ハンズオンの並行処理:
    - トレーシングをしながら並行処理を学ぶ。
    - https://t.co/b7qKE9C0lU?amp=1

### 並行処理の次のテーマ候補
- システムプログラミング
    - https://www.lambdanote.com/products/go
- インタプリタを作ってみる？
    - https://www.oreilly.co.jp/books/9784873118222/

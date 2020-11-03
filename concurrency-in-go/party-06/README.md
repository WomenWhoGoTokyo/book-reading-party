# Party-06 2020-08-22
## 読んだところ
- 4章 Go での並行処理パターン
    - 4.5 エラーハンドリング
    - 4.6 パイプライン
    - 4.7 ファンアウト、ファンイン
    - 4.8 or-done チャネル
    - 4.9 tee チャネル
    - 4.10 bridge チャネル
    - 4.11 キュー
    
## Note
いつも忘れる:
```
<-chan が読み込み専用権限
chanまたはchan<- が書き込み権限
```

- パイプライン、シェルのパイプラインのようなイメージ。
- メモリフットプリント:
    - コンピュータ上でプログラムが動作する際に使用されるメインメモリの容量 (メモリ使用量) のことである。 一般的に、大規模なプログラムほど大きなフットプリントを必要とする。
- 文章の誤りかも？
    - Miki さんが Issue 立てた。
    - https://github.com/oreilly-japan/concurrency-in-go-support/issues/43
- 型アサーション:
    - interface型から具体的な型に変換するやつ
    - https://go-tour-jp.appspot.com/methods/15
    - interface 型の基の型に見当がついてたり特定の型に変換したい時は型アサーションで、型が分かってない時は型スイッチ使います。
- NumCPU:
    - https://pkg.go.dev/runtime?tab=doc#NumCPU
- bufio:
    - バッファアイオー
- エラトステネスの篩:
    - ttps://note.com/mathchannel/n/n274845e028b6https://note.com/mathchannel/n/n274845e028b6
- Tee:
    - https://ja.wikipedia.org/wiki/Tee_(UNIX)
- リトルの法則:
    - https://ja.wikipedia.org/wiki/リトルの法則

## Playground
今回、コード多めだった。次からは事前に playground いくつか用意しておこう！と思います。
- add + multiply:
    - https://play.golang.org/p/9CTn_sxyU_O
    - https://play.golang.org/p/gGHAo898UPn
- パイプライン:
    - https://play.golang.org/p/bdBjweD-foB
- toString:
    - https://play.golang.org/p/sPD1qBtZsIZ

### Next
- 4章 4.12 context パッケージ から

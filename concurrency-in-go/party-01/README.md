# Party-01 2020-06-27
## 読んだところ
- 訳者まえがき
- 序文
- 1章 並行処理入門
    - 1.2.4 デッドロック、ライブロック、リソース枯渇
        - P14 ライブロックの手前まで
    
## Note
- https://go.dev/
- シュタゲのようである。
- ムーアの法則: 半導体の集積率は18ヶ月で2倍になる。
- 驚異的並列: Amazing juxtaposition
    - 並列処理の演習に出てくる。
    - http://www.ie.u-ryukyu.ac.jp/~e085739/_downloads/chap31.pdf
- GUI ベースのプログラムと pi 計算の違い:
    - 人間による相互作用 (直列) vs 驚異的並列が使える
- Web スケール:
    - https://i.dell.com/sites/csdocuments/Shared-Content_data-Sheets_Documents/ja/jp/APJ-DEIESRC67979-Nutanix-Solutions-brochure-JA.pdf
- `time.Sleep()` する前のコードは、理論としては、次のことがあるのですよね。
    - 0 がプリントされることも
    - 1 がプリントされることも
    - 何もプリントされないことも
    - 何回か実行しても 0 しか出力されないのは確率の問題なのかしら...
        - 処理速度とか、中の処理の重さによって違うので、今回は `data++` だからかなぁ
- https://ja.wikipedia.org/wiki/クリティカルセクション
- https://qiita.com/taizo/items/d73e54111bd470c1f1b3
- Sync パッケージ:
    - https://pkg.go.dev/sync?tab=doc
    - https://pkg.go.dev/golang.org/x/sync/errgroup?tab=doc

```
パッケージ同期は、相互排他ロックなどの基本的な同期プリミティブを提供します。Once 型と WaitGroup 型以外は、ほとんどが低レベルのライブラリルーチンでの使用を意図しています。より高レベルの同期化は、チャンネルや通信を介して行うのが良いでしょう。
このパッケージで定義されている型を含む値はコピーしてはいけません。
```

- https://github.com/kat-co/concurrency-in-go-src/blob/master/an-introduction-to-concurrency/why-is-concurrency-hard/memory-access-synchronization/fig-basic-memory-access-sync.go
- デッドロックの条件:
    - コフマン条件: 相互排他/条件待ち/横取り不可/循環待ち
    - 全ての条件を満たした場合デッドロックが発生する

### Next
- 2020-07-05 10:00 - 12:00
- P14 ライブロックから


## Code
- https://play.golang.org/p/HjnSW6L7E2F
- https://play.golang.org/p/ldbqdkpgFBB
- https://play.golang.org/p/P3x6izUiGsR

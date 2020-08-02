
# Party-5 2020-08-01
## 読んだところ
- 3章 Go における並行処理の構成要素
    - 3.4 select 文
    - 3.5 GOMAXPROCS レバー
    - 3.6 まとめ
- 4章 Go での並行処理パターン
    - 4.1 拘束
    - 4.2 for-select ループ
    - 4.3 ゴルーチンリークを避ける
    - 4.4 or チャネル
    
## Note
- チャネルおさらい
```
var receiveChan <-chan interface{}
var sendChan chan<- interface{}
```
- <-chan が読み込み専用権限
- chanまたはchan<- が書き込み権限

- この書き方、どういうものなんだろう。
```
c1 := make(chan interface{}); close(c1)
c2 := make(chan interface{}); close(c2)
```

- 問題空間とは...？？
    - 文脈。どんな問題を抱えていてこの処理をしようとしているか、とかそういうこと。

- `func GOMAXPROCS(n int) int` については以下を読んでおくと良さそう:
    - https://golang.org/pkg/runtime/
    - https://deeeet.com/writing/2014/07/30/golang-parallel-by-cpu/

- Immutable = 作成後にその状態が変わらないオブジェクト
- ポリモーフィズムの話:
    - https://ja.wikipedia.org/wiki/ポリモーフィズム
    - https://en.wikibooks.org/wiki/Introduction_to_Programming_Languages/Ad_Hoc_Polymorphism
- for...select でループを抜けたい:
    - https://qiita.com/cia_rana/items/a2c3e1609bd25a5c5596  
- `close` するのは送り手のチャネルだけ。
- 再帰 (さいき):
    - あるものについて記述する際に、記述しているものそれ自身への参照が、その記述中にあらわれることをいう。
- or パターン:
    - http://t10471.hatenablog.com/entry/2019/09/10/054128
    - https://dasshshsd.hatenablog.com/entry/2019/10/22/005702
- Go の並行処理パターンを見てみよう:
    - https://github.com/adityamenon/Google-IO_2012_Go-Concurrency-Patterns
    - https://www.ymotongpoo.com/works/goblog-ja/

## Playground
- https://play.golang.org/p/ni-VVk4GEOQ
- 複数のチャネルが読み込んでいる場合:
    - https://play.golang.org/p/PBz4-22tlgT
- 1つも読み込みがされない例:
    - https://play.golang.org/p/ezkZ00nyfZp
- チャネルがすべてブロックされているのに、でもなにかしたい:
    - https://play.golang.org/p/M2Q1r8TPUcl
- ループがなにかをしてて、止めるべきか確認している:
    - https://play.golang.org/p/EO63Fe9c7V9
- アドホック拘束の例:
    - https://play.golang.org/p/_Pw7WbZgy0E
- レキシカル拘束の例:
    - https://play.golang.org/p/ri1K9iID6TD
- 並行安全でないデータ構造を使った拘束:
    - https://play.golang.org/p/7XQdGJdoKCs
- ゴルーチンリークの例 [1]:
    - https://play.golang.org/p/bcxnvESO-N1
- ゴルーチンリークの例 [2]:
    - https://play.golang.org/p/I30zgFIVYpA
- defer された Println が実行されない例:
    - https://play.golang.org/p/sEcazV80Bpj
- https://play.golang.org/p/GN7mBKB9dwD
- or パターン:
    - https://play.golang.org/p/zl_qefvq9LI

### Next
- 2020-08-xx 10:00 - 12:00
- 3章 4.5 エラーハンドリング (P99) から

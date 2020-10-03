# Party-8 2020-09-20
## 読んだところ
- 5章 大規模開発での並行処理
    - 5.3 ハートビート
    - 5.4 複製されたリクエスト
    - 5.5 流量制限
    
## Note
いつも忘れる:
```
<-chan が読み込み専用権限
chanまたはchan<- が書き込み権限
```

- Heartbeats について:
    - https://syossan.hateblo.jp/entry/2018/05/26/175835
- https://ja.wikipedia.org/wiki//dev/null#:~:text=%2Fdev%2Fnull（nullデバイス,ない（EOFを返す）。
- https://ja.wikipedia.org/wiki/定足数
- https://docs.oracle.com/cd/E19957-01/806-5348/6jec7p68e/index.html

> イギリスで治安判事(justices of the peace)を任命するとき、多くの判事の中から（任命状の「その中で(of whom)」にあたるラテン語がquorumだった）執行に必要不可欠な判事を指名し、これらの判事がjustices of quorumと呼ばれたことから。「定足数」の意味で使われるようになったのは1616年から

- https://en.wikipedia.org/wiki/Quorum_(distributed_computing)
- https://ja.wikipedia.org/wiki/トークンバケット
- RateLimit: https://github.com/juju/ratelimit
- http://e-words.jp/w/バーストトラフィック.html

## Playground

### Next
- 5章 5.6 不健全なゴルーチンを直す から

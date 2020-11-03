# Party-04 2020-07-24
## 読んだところ
- 3章 Go における並行処理の構成要素
    - 3.2.5 Pool
    - 3.3 チャネル (channel)
    
## Note
- Object Pool パターン:
    - デザインパターンのひとつ。
    - https://thinkit.co.jp/article/934/1
- `暖気` とは:
    - https://dev.classmethod.jp/articles/elastic-load-balancing-pre-warming/
    - 「予めロードバランサーを増やしておくことを暖機と言います。」    
- playground で Test や Benchmark したい:
    - https://qiita.com/cia_rana/items/b6bd6dd9f5af3f95ea2a
    - https://play.golang.org/p/36UPL27Y7f
    - https://kaneshin.hateblo.jp/entry/2018/12/08/231045
- [疑問] 「プール内のオブジェクトはおよそ均質なものであるべき」というのは、どういうことなのか。

## Playground
読みながら書いていったコード集です。typo 多めな...
- Sync.Pool の例:
    - https://play.golang.org/p/28conS78l7k
- ベンチマーク (今は playground で動かないのであとで何とかする) の例:
    - https://play.golang.org/p/DXsqDhNly__7
- Pool を利用した例:
    - https://play.golang.org/p/pSQLW8GIDSm
- chan の例:
    - https://play.golang.org/p/cBWav9oGaKn
- 書き込みと読み込みがチグハグな例:
    - https://play.golang.org/p/5oR_LCvKUBB
- チャネルの戻り値:
    - https://play.golang.org/p/GpSvRy86lAT
    - false, 指定した型のゼロ値
- 繰り返し:
    - https://play.golang.org/p/MM2NkxQLRt0
- バッファ:
    - https://play.golang.org/p/K6wS8JOA4-m

### Next
- 2020-08-02 10:00 - 12:00
- 3章 3.4 select から
- はじめの部分で、チャネル (channel) のふりかえりをします！ (大事だから)
    - 復習用: https://tour.golang.org/concurrency/2

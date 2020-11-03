# Party-03 2020-07-18
## 読んだところ
- 3章 Go における並行処理の構成要素
    - 3.1 ゴルーチン (goroutine)
    - 3.2 sync パッケージ
    - 3.2.1 WaitGroup
    - 3.2.2 Mutex と RWMutex
    - 3.2.3 Cond
    - 3.2.4 Once
    
 ## Note
- プリエンプティブ (preemptive) な OS とは:
    - 実行するタスクを OS が一定時間ごとに切り替える OS を言います。
- コルーチン:
    - サブルーチンがエントリーからリターンまでを一つの処理単位とするのに対し、コルーチンはいったん処理を中断した後、続きから処理を再開できる。
- https://postd.cc/scheduler/
- 駄目な例:
    - https://play.golang.org/p/LQEp0kaKTZO
- クロージャの例:
    - https://play.golang.org/p/3HvnsVtCB6a
- 変数を引数で goroutine に渡す:
    - https://play.golang.org/p/qvD1iwBh6Ht
- 漸近 (ぜんきん):
    - 徐々に近づいていくさま。特に数値などについて言う。
- 大数の法則:
    - https://ja.wikipedia.org/wiki/大数の法則
- panic: sync: negative WaitGroup counter
    - https://play.golang.org/p/l2vc0ovlYoZ
- fatal error: all goroutines are asleep - deadlock!
    - https://play.golang.org/p/Une2CxavgOS
- ちゃんと動くやつ:
    - https://play.golang.org/p/zecH1tQ_ZYD
- RWMutex:
    - https://pkg.go.dev/sync?tab=doc#RWMutex
- time.Since ってなに？:
    - https://golang.org/pkg/time/#Since
    - https://qiita.com/taizo/items/acbee530bd33c803dab4
```
func Since
func Since(t Time) Duration
Since returns the time elapsed since t. It is shorthand for time.Now().Sub(t).
```

- tabwriter ってなんだっけ？:
    - https://pkg.go.dev/text/tabwriter?tab=doc
- Locker interface:
```
type Locker interface {
    Lock()
    Unlock()
}
```
- FIFO:
    - First in First Out 先入先出法
- sync.Once:
    - https://play.golang.org/p/cgQv-OvARzi
    - どこからでも呼ばれる可能性があるけど一回しか実行したくない時。init したい時とかに使う。
    - https://www.slideshare.net/takuyaueda967/gaegosync
        - P13
- Broadcast を使った例:
  - とあるメッセージを全てのユーザに対して一気に送信する。みたいな時に使うこともある。  

### Next
- 3章 3.2.5 Pool から

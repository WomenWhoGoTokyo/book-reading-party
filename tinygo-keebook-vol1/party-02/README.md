# Party-02 2025-08-17


## 参加者

- micchie
- marie
- momo

## 読んだところ
- はじめに
- 第1章 TinyGo への貢献
- 第2章 2025 年のTinyGo のnet package
- 第3章 TinyGo でも通信がしたい！ ～ W5500-EVB-Pico-PoE でTCPIP 通信～

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/725156029033218080/1406432769684541470)

```
$ tinygo flash --target wioterminal --size short examples/blinky1
```

```
tinygo flash --target wioterminal --size short examples/serial
```

たくさん USB をつなげているとき、ports を調べて wioterminal がどこにつながっているか確認する。

```
tinygo ports 
```

wiotermiaal のポートを指定して monitor する。

```
tinygo monitor --port /dev/cu.usbmodem1101
```

- https://github.com/sago35/tinygo-examples

```
tinygo flash --target wioterminal --size short ./wioterminal/microphone/
```

勉強会に持っていって自慢しよう！

```
tinygo flash --target wioterminal --size short ./wioterminal/wiobadge/
```

## 次回

- 2025-08-30 (土) 10:00
- P51 第4章 ゲーム作成ガイド / 4.3 ゲームの設定

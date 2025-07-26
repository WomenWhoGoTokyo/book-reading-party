# Party-01 2025-07-26


## 参加者

- micchie
- momo
- marie

## 読んだところ
- はじめに
- 第1章 TinyGo への貢献 〜 1.4 まで

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/1398459069429317672)

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

- 2025-08-03 (日) 10:00
- P15 第1章 TinyGo への貢献 / 1.5 速度測定 〜

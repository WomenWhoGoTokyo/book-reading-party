# Party-03 2025-09-20


## 参加者

- micchie
- marie
- momo

## 読んだところ
- 第4章 ゲーム作成ガイド / 4.3 ゲームの設定
- 第5章 TinyGo で作るMIDI キーボード

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/1411135777219412078)

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

- 2025-10-05 (日) 10:00
- P77 第5章 TinyGo で作るMIDI キーボード 5.5.5 ハードウェア制御 〜

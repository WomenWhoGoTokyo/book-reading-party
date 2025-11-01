# Party-04 2025-10-05


## 参加者

- micchie
- marie
- momo
- Misato

## 読んだところ
- 第5章 TinyGo で作るMIDI キーボード

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/725156029033218080/1424190050257735742)

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

- 2025-10-18 (土) 10:00
- 第5章 TinyGo で作るMIDI キーボード 実践

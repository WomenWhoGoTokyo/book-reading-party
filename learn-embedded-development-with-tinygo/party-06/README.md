# Party-06 2025-02-24

https://github.com/sago35/tinygobook

## 参加者

- micchie
- momo
- marie
- miki
- sasami
- sago

## 読んだところ
- Chapter5 各ペリフェラルの使い方

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/725156029033218080/1343368806537494599)

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

- 2025-02-24 (月) 10:00
- Chapter5 各ペリフェラルの使い方 P125

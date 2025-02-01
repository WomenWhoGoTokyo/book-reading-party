# Party-04 2025-02-01

https://github.com/sago35/tinygobook

## 参加者

- micchie
- marie
- momo
- sago

## 読んだところ
- Chapter3 Goの基本
  - 08 インターフェース
  - 09 unsafe
  - 10 cgo

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/725156029033218080/1335048305926864916)

```
$ tinygo build -o hello.uf2 -target wioterminal -size short examples/blinky1
```

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

- 2025-02-11 (火) 10:00
- Chapter4 TinyGO Internals P103

# Party-05 2025-02-11

https://github.com/sago35/tinygobook

## 参加者

- micchie
- misato
- marie
- sago

## 読んだところ
- Chapter4 TinyGO Internals
  - 01 GOROOT と TINYGOROOT
  - 02 ビルドタグ
  - 03 TinyGo の標準 package
  - 04 TinyGo のビルドの流れ
  - 05 TinyGo の実行

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/725156029033218080/1338660342330556459)

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

- 2025-02-24 (月) 10:00
- Chapter5 各ペリフェラルの使い方 P125

# Party-03 2025-01-25

https://github.com/sago35/tinygobook

## 参加者

- micchie
- misato
- marie
- sago
- momo

## 読んだところ
- Chapter3 Goの基本
  - 04 制御フロー
  - 05 関数
  - 06 メソッド
  - 07 goroutineとchannelとselect

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/725156029033218080/1332507443900452934)

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

## 次回

- 2025-02-01 (土) 10:00
- 08 インターフェース P82

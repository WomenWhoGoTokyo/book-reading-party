# Party-02 2025-01-11

https://github.com/sago35/tinygobook

## 参加者

- micchie
- momo
- misato
- marie

## 読んだところ
- はじめに
- 3. Goの基本

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/725156029033218080/1324898602476245179)

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

- 2025-01-25 (土) 10:00
- 3. Goの基本 04 制御フロー から

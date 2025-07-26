# Party-15 2025-07-05

https://github.com/sago35/tinygobook

## 参加者

- micchie
- momo
- marie
- sago

## 読んだところ
- chapter 7
  - 09 HTTP/HTTPS

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/725156029033218080/1390849929332789359)

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

- 2025-07-21 (月) 10:00
- Chapter8 アプリケーション作成

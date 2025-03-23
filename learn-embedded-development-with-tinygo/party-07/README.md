# Party-07 2025-03-20

https://github.com/sago35/tinygobook

## 参加者

- micchie
- momo
- misato
- sago

## 読んだところ

## Note

[Discord Thread](https://discord.com/channels/689414179752247409/1352085367452598363)

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

- 2025-03-23 (日) 10:00

package main

import (
	"fmt"
	"time"
)

const timeFormat = "15:04:05.000"

func main() {
	timer1(3)
	timer2(3)
}

// timer1は、rangeによりチャネルに値が入るたびにループが回る性質を利用したタイマー
// time.Afterにより一定時間経過後に１回ループが回って処理が終了するフローになる
func timer1(sec int) {
	fmt.Println("START ", getTimeStr(time.Now()))

	for a := range time.After(time.Duration(sec) * time.Second) {
		fmt.Println("FINISH", getTimeStr(a))
		return
	}
}

// timer2は、チャネルからデータが受信できるまで処理をブロックする性質を利用したタイマー
func timer2(sec int) {
	fmt.Println("START ", getTimeStr(time.Now()))
	<-time.After(time.Duration(sec) * time.Second)
	fmt.Println("FINISH", getTimeStr(time.Now()))

}

// getTimerStrは、フォーマットした時刻文字列を返却する
func getTimeStr(t time.Time) string {
	return t.Local().Format(timeFormat)
}

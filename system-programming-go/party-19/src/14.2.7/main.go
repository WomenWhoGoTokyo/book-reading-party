package main

import (
	"fmt"
	"runtime"
	"sync"
)

func calc(id, price int, interestRate float64, year int) {
	months := year * 12
	interest := 0
	for i := 0; i < months; i++ {
		balance := price * (months - i) / months
		interest += int(float64(balance) * interestRate / 12)
	}
	fmt.Printf("year=%d total=%d interest=%d id=%d\n",
		year, price+interest, interest, id)
}

// ワーカー
func worker(id, price int, interestRate float64, years chan int,
	wg *sync.WaitGroup) {
	// タスクがなくなってタスクのチャネルがclose されるまで無限ループ
	for year := range years {
		calc(id, price, interestRate, year)
		wg.Done()
	}
}
func main() {
	// 借入額
	price := 40000000
	// 利子1.1% 固定
	interestRate := 0.011
	// タスクはchan に格納
	years := make(chan int, 35)
	for i := 1; i < 36; i++ {
		years <- i
	}
	var wg sync.WaitGroup
	wg.Add(35)
	// CPU コア数分のgoroutine 起動
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, price, interestRate, years, &wg)
	}
	// すべてのワーカーが終了する
	close(years)
	wg.Wait()
}

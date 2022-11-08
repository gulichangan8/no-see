package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	wg.Add(3)
	go func() {
		for {
			if t.Hour() == 02 {
				fmt.Println("我还能再战4小时！")
			}
		}
	}()
	go func() {
		for {
			if t.Hour() == 06 {
				fmt.Println("我要去图书馆开卷！")
			}
		}
	}()
	ticker := time.NewTicker(time.Second * 30)
	go func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("芜湖！起飞！")
			}
		}
	}(ticker)
	wg.Wait()
}

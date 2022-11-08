package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

func main() {
	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	var define int
	var s time.Duration
	var i time.Duration
	var S time.Duration
	var Define bool
	w.Add(3)
	go contain(t, 02, "我还能再战4小时！")
	go contain(t, 06, "我要去图书馆开卷！")
	go repeat(30 * time.Second)
	fmt.Println("您是否想自定义一个计时器或闹钟（1.闹钟 2.定时器 3.取消自定义）")
	fmt.Scan(&define)
	if define == 1 {
		var hour int
		fmt.Println("您想将闹钟定在几点：")
		fmt.Scan(&hour)
		w.Add(1)
		go contain(t, hour, "起床了")
	}
	if define == 2 {
		var ok int
		fmt.Println("您是否需要此定时器循环计时（1.需要 2.不需要）")
		fmt.Scan(&ok)
		if ok == 1 {
			fmt.Println("您想将定时器定为多少秒一次：")
			fmt.Scan(&s)
			go repeat(s * time.Second)
		}
		if ok == 2 {
			fmt.Println("您想将定时器定为多少秒一次：")
			fmt.Scan(&s)
			w.Add(1)
			go notRepeat(s * time.Second)
		}
	}
	fmt.Println("您是否想删除一个计时器（想请输入true,不想则输入false）")
	fmt.Scan(&Define)
	if Define {
		fmt.Println("您想删除循环定时器还是不循环定时器（1.循环 2.不循环）")
		fmt.Scan(&i)
		if i == 1 {
			fmt.Println("您想删除时长为多少秒的定时器：")
			fmt.Scan(&S)
			notRepeat(S * time.Second).Stop()
		}
		if i == 2 {
			fmt.Println("您想删除时长为多少秒的定时器：")
			fmt.Scan(&S)
			repeat(S * time.Second).Stop()
		}
	}
	fmt.Println("您是否想重置一个定时器（想请输入true,不想则输入false）")
	fmt.Scan(&Define)
	if Define {
		fmt.Println("您想重置循环定时器还是不循环定时器（1.循环 2.不循环）")
		fmt.Scan(&i)
		if i == 1 {
			fmt.Println("您想重置时长为多少秒的定时器：")
			fmt.Scan(&S)
			notRepeat(S * time.Second).Reset(i * time.Second)
		}
		if i == 2 {
			fmt.Println("您想删除时长为多少秒的定时器：")
			fmt.Scan(&S)
			repeat(S * time.Second).Reset(S * time.Second)
		}
	}
	w.Wait()
}
func contain(t time.Time, hour int, str string) {
	for {
		if t.Hour() == hour {
			fmt.Println(str)
		}
	}
}
func repeat(s time.Duration) *time.Ticker {
	ticker := time.NewTicker(s)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("芜湖！起飞！")
		}
		return ticker
	}
}
func notRepeat(s time.Duration) *time.Ticker {
	ticker := time.NewTicker(s)
	for {
		<-ticker.C
		fmt.Printf("芜湖！起飞！")
		return ticker
	}
}

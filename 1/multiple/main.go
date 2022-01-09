package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var fmeng = false

func producer(id int, wg *sync.WaitGroup, ch chan string) {
	count := 0
	for !fmeng {
		time.Sleep(time.Second * 1)
		count++
		data := strconv.Itoa(id) + "===" + strconv.Itoa(count)
		fmt.Printf("product 发送 ,%s\n", data)
		ch <- data
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(time.Second * 1)
		fmt.Printf("comsumer 接收了,%s \n", data)
	}
	wg.Done()
}
func main() {
	chanSteam := make(chan string, 10)
	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wgPd.Add(1)
		go producer(i, wgPd, chanSteam)
	}
	for j := 0; j < 2; j++ {
		wgCs.Add(1)
		go consumer(wgCs, chanSteam)
	}
	go func() {
		time.Sleep(time.Second * 10)
		fmeng = true
	}()
	wgPd.Wait()
	close(chanSteam)
}

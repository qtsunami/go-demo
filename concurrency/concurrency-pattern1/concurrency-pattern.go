// main goroutine 通过 spawn 函数返回的channel 与新的子goroutine 建立联系
// 这个 channel 的用途就是在两个 goroutine 之间建立退出事件的"信号"通信机制。
// main goroutine 在创建完新 goroutine 后便在该 channel 上阻塞等待，直到新 goroutine 退出前向该 channel 发送了一个信号

package main

import "time"

func worker(args ...interface{}) {
	if len(args) == 0 {
		return
	}

	interval, ok := args[0].(int)
	if !ok {
		return
	}

	println("worker doing ...")
	time.Sleep(time.Second * (time.Duration(interval)))
}

func spawn(f func(args ...interface{}), args ...interface{}) chan struct{} {
	c := make(chan struct{})

	go func() {
		f(args...)
		c <- struct{}{}
	}()

	println("spawn over")
	return c
}

func main() {
	done := spawn(worker, 5)
	println("spawn a worker goroutine")

	<-done
	println("worker done")
}

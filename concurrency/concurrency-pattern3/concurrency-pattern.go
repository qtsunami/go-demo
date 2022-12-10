// main goroutine 通过 spawn 函数返回的channel 与新的子goroutine 建立联系
// 这个 channel 的用途就是在两个 goroutine 之间建立退出事件的"信号"通信机制。
// main goroutine 在创建完新 goroutine 后便在该 channel 上阻塞等待，直到新 goroutine 退出前向该 channel 发送了一个信号

// 通过 Go语言提供的 sync.WaitGroup 实现等待多个 goroutine 退出模式

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(args ...interface{}) {
	if len(args) == 0 {
		return
	}

	interval, ok := args[0].(int)
	if !ok {
		return
	}

	time.Sleep(time.Second * (time.Duration(interval)))
}

func spawnGroup(n int, f func(args ...interface{}), args ...interface{}) chan struct{} {
	c := make(chan struct{})
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			name := fmt.Sprintf("worker-%d:", i)
			f(args...)
			println(name, "done")
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		c <- struct{}{}
	}()

	return c
}

func main() {
	done := spawnGroup(5, worker, 3)
	println("spawn a group of workers")
	<-done
	println("group workers done")
}

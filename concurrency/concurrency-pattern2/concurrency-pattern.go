// main goroutine 通过 spawn 函数返回的channel 与新的子goroutine 建立联系
// 这个 channel 的用途就是在两个 goroutine 之间建立退出事件的"信号"通信机制。
// main goroutine 在创建完新 goroutine 后便在该 channel 上阻塞等待，直到新 goroutine 退出前向该 channel 发送了一个信号
// 获取 goroutine 的退出状态

package main

import (
	"errors"
	"fmt"
	"time"
)

var OK = errors.New("ok")

func worker(args ...interface{}) error {
	if len(args) == 0 {
		return errors.New("invalid args")
	}

	interval, ok := args[0].(int)
	if !ok {
		return errors.New("invalid interval arg")
	}

	time.Sleep(time.Second * (time.Duration(interval)))

	return OK
}

func spawn(f func(args ...interface{}) error, args ...interface{}) chan error {
	c := make(chan error)

	go func() {
		c <- f(args...)
	}()

	return c
}

func main() {
	done := spawn(worker, 5)
	println("spawn worker1")
	err := <-done
	fmt.Println("worker1 done:", err)

	done = spawn(worker)
	println("spawn worker2")
	err = <-done
	fmt.Println("worker2 done: ", err)
}

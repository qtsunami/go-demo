package main

import (
	"time"
)

// 并发版

const (
	idCheckTmCost    = 60  // 身份检查耗费时间
	bodyCheckTmConst = 120 // 人身检查耗费时间
	xRayCheckTmCost  = 180 // X光机检查耗费时间
)

func idCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	print("\tgoroutine-", id, ":idCheck ok\n")
	return idCheckTmCost
}

func bodyCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmConst))
	print("\tgoroutine-", id, ":bodyCheck ok\n")
	return bodyCheckTmConst
}

func xRayCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	print("\tgoroutine-", id, ":xRayCheck ok\n")
	return xRayCheckTmCost
}

func start(id string, f func(string) int, next chan<- struct{}) (chan<- struct{}, chan<- struct{}, <-chan int) {

	queue := make(chan struct{}, 10)
	quit := make(chan struct{})
	result := make(chan int)

	go func() {
		total := 0
		for {
			select {
			case <-quit:
				result <- total
				return
			case v := <-queue:
				total += f(id)
				if next != nil {
					next <- v
				}
			}
		}
	}()
	return queue, quit, result
}

func newAirportSecurityCheckChannel(id string, queue <-chan struct{}) {
	go func(id string) {
		print("goroutine-", id, ": airportSecurityCheckChannel is ready ...\n")

		// 启动 X 光检查
		queue3, quit3, result3 := start(id, xRayCheck, nil)
		// 启动人身检查
		queue2, quit2, result2 := start(id, bodyCheck, queue3)
		// 启动身份检查
		queue1, quit1, result1 := start(id, idCheck, queue2)

		for {
			select {
			case v, ok := <-queue:
				if !ok {
					close(quit1)
					close(quit2)
					close(quit3)
					total := max(<-result1, <-result2, <-result3)
					print("goroutine-", id, ": airportSecurityCheckChannel time cost:", total, "\n")
					print("goroutine-", id, ": airportSecurityCheckChannel closed\n")
					return
				}

				queue1 <- v
			}
		}

	}(id)
}

func max(args ...int) int {
	n := 0
	for _, v := range args {
		if v > n {
			n = v
		}
	}

	return n
}

func main() {
	total := 0
	passengers := 30

	queue := make(chan struct{}, 30)
	newAirportSecurityCheckChannel("channel1", queue)
	newAirportSecurityCheckChannel("channel2", queue)
	newAirportSecurityCheckChannel("channel3", queue)

	time.Sleep(5 * time.Second)

	for i := 0; i < passengers; i++ {
		queue <- struct{}{}
	}

	time.Sleep(5 * time.Second)
	close(queue)
	time.Sleep(1000 * time.Second)

	println("total time const:", total)
}

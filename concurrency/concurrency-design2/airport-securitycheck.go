package main

import "time"

const (
	idCheckTmCost    = 60  // 身份检查耗费时间
	bodyCheckTmConst = 120 // 人身检查耗费时间
	xRayCheckTmCost  = 180 // X光机检查耗费时间
)

func idCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	print("\tgoroutine-", id, ":idCheck ok\n")
	return idCheckTmCost
}

func bodyCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmConst))
	print("\tgoroutine-", id, ":bodyCheck ok\n")
	return bodyCheckTmConst
}

func xRayCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	print("\tgoroutine-", id, ":xRayCheck ok\n")
	return xRayCheckTmCost
}

func airportSecurityCheck(id int) int {
	print("\tgoroutine-", id, ":airportSecurityCheck ...\n")
	total := 0

	total += idCheck(id)
	total += bodyCheck(id)
	total += xRayCheck(id)

	print("\tgoroutine-", id, ":airportSecurityCheck ok\n")

	return total
}

func start(id int, f func(int) int, queue <-chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		total := 0
		for {
			_, ok := <-queue
			if !ok {
				c <- total
				return
			}
			total += f(id)
		}
	}()

	return c
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

	c := make(chan struct{})

	c1 := start(1, airportSecurityCheck, c)
	c2 := start(2, airportSecurityCheck, c)
	c3 := start(3, airportSecurityCheck, c)

	for i := 0; i < passengers; i++ {
		c <- struct{}{}
	}

	close(c)

	total = max(<-c1, <-c2, <-c3)

	println("total time const:", total)
}

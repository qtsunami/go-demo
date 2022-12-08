package main

import "time"

const (
	idCheckTmCost    = 60  // 身份检查耗费时间
	bodyCheckTmConst = 120 // 人身检查耗费时间
	xRayCheckTmCost  = 180 // X光机检查耗费时间
)

func idCheck() int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tidCheck ok")
	return idCheckTmCost
}

func bodyCheck() int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmConst))
	println("\tbodyCheck ok")
	return bodyCheckTmConst
}

func xRayCheck() int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	println("\txRayCheck ok")
	return xRayCheckTmCost
}

func airportSecurityCheck() int {
	println("airportSecurityCheck ...")
	total := 0

	total += idCheck()
	total += bodyCheck()
	total += xRayCheck()

	println("airportSecurityCheck ok")

	return total
}

func main() {
	total := 0
	passengers := 30

	for i := 0; i < passengers; i++ {
		total += airportSecurityCheck()
	}

	println("total time const:", total)
}

package main

import (
	"fmt"

	"./CacheLogic"
	"./ServiceLogic"
)

func main() {
	mockRedis := CacheLogic.MockRedis{InstanceID: "xx"}
	x := ServiceLogic.MockServiceHandler{InstanceID: "sdf", CacheHandler: mockRedis}
	handleService(x)
	fmt.Println("Test")

}

func handleService(s ServiceLogic.ServiceHandler) {
	s.ServiceMethod()
}

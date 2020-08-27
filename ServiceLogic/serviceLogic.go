package ServiceLogic

import (
	"fmt"

	"../CacheLogic"
)

type MockServiceHandler struct {
	InstanceID   string
	CacheHandler CacheLogic.CacheHandler
}

func (s MockServiceHandler) ServiceMethod() {
	fmt.Println(s.InstanceID)
	fmt.Println(s.CacheHandler.Get("gds"))
}

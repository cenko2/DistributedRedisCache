package ServiceLogic

import (
	"fmt"
	"net/http"

	"../CacheLogic"
	"github.com/gorilla/mux"
)

type MockServiceHandler struct {
	ParamName           string
	DefaultTTLInMinutes int
	CacheHandler        CacheLogic.CacheHandler
}

func (s MockServiceHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := (vars[s.ParamName])

	if !s.CacheHandler.KeyExists(key) {
		fmt.Println("Key " + key + " not found")
	} else {
		fmt.Println("Cache hit!")
	}

	fmt.Println(s.CacheHandler.Get("gds"))
}

func (s MockServiceHandler) HandlePost(w http.ResponseWriter, r *http.Request) {

	fmt.Println(s.CacheHandler.Get("ParamName"))
}

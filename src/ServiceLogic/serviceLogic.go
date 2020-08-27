package ServiceLogic

import (
	"fmt"
	"io"
	"net/http"
    "bytes"

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
		http.Error(w, "Key not found", http.StatusNotFound)
	} else {
		fmt.Println("Cache hit!")
		value := s.CacheHandler.Get(key)
		io.WriteString(w, value)
	}

	fmt.Println(s.CacheHandler.Get("gds"))
}

func (s MockServiceHandler) HandlePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := (vars[s.ParamName])

	buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
    value := buf.String()
	fmt.Println( "Update called with value " + value + " key : "+key)
	s.CacheHandler.Insert(key,value,s.DefaultTTLInMinutes)
}

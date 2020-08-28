package ServiceLogic

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"../CacheLogic"
	"github.com/gorilla/mux"
)

type ServiceHandler struct {
	ParamName           string
	DefaultTTLInMinutes int
	CacheHandler        CacheLogic.CacheHandlerInterface
}

func (s ServiceHandler) HandleGet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := (vars[s.ParamName])
	now := time.Now()
	defer func() {
		log.Printf("Get\tkey:%q\ttime:%v", key, time.Since(now))
	}()

	if !s.CacheHandler.KeyExists(key) {
		fmt.Println("Key " + key + " not found")
		http.Error(w, "Key not found", http.StatusNotFound)
	} else {
		fmt.Println("Cache hit!")
		value := s.CacheHandler.Get(key)
		io.WriteString(w, value)
	}
}

func (s ServiceHandler) HandlePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := (vars[s.ParamName])
	now := time.Now()
	defer func() {
		log.Printf("Post\tkey:%q\ttime:%v", key, time.Since(now))
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	value := buf.String()
	fmt.Println("Update called with value " + value + " key : " + key)
	s.CacheHandler.Insert(key, value, s.DefaultTTLInMinutes)
}

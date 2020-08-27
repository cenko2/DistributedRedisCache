package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"os"

	"./CacheLogic"
	"./ServiceLogic"
	"github.com/gorilla/mux"
)

const (
	defaultAddr = ":http"
)

func main() {

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = defaultAddr
	}

	mockRedis := CacheLogic.MockRedis{InstanceID: "xx"}
	x := ServiceLogic.MockServiceHandler{ParamName: "key", CacheHandler: mockRedis, DefaultTTLInMinutes: 30}

	
	r := mux.NewRouter()
	r.HandleFunc("/cache/{key}", x.HandleGet).Methods("GET")
	r.HandleFunc("/cache/{key}", x.HandlePost).Methods("POST")
	r.HandleFunc("/status", handleHearBeat).Methods("GET")
	http.Handle("/", r)
    fmt.Println("Address :" + addr)
	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Instance initialized routing complete")
	log.Fatal(srv.ListenAndServe())
}

func  handleHearBeat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Heartbeat recevied")
}
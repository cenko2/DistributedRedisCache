package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"./CacheLogic"
	"./ServiceLogic"
	"github.com/gorilla/mux"
)

func main() {
	mockRedis := CacheLogic.MockRedis{InstanceID: "xx"}
	x := ServiceLogic.MockServiceHandler{ParamName: "key", CacheHandler: mockRedis, DefaultTTLInMinutes: 30}

	fmt.Println("Test")
	r := mux.NewRouter()
	r.HandleFunc("/cache/{key}", x.HandleGet).Methods("GET")
	r.HandleFunc("/cache/{key}", x.HandleGet).Methods("POST")

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

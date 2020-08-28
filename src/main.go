package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"./CacheLogic"
	"./ServiceLogic"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

const (
	defaultAddr = ":http"
)

func main() {

	fmt.Println(os.Getenv(os.Getenv("ENV")))

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = defaultAddr
	}
	//  docker pull redis
	//  docker run -d -p 6379:6379 redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisClient := CacheLogic.RedisCache{Rdb: rdb}

	x := ServiceLogic.ServiceHandler{ParamName: "key", CacheHandler: redisClient, DefaultTTLInMinutes: 30}

	r := mux.NewRouter()
	r.HandleFunc("/cache/{key}", x.HandleGet).Methods("GET")
	r.HandleFunc("/cache/{key}", x.HandlePost).Methods("POST")
	r.HandleFunc("/status", handleHearBeat).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Address :" + addr)
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Instance initialized routing complete")
	log.Fatal(srv.ListenAndServe())
}

func handleHearBeat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Heartbeat recevied")
}

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"GoSleepTimerTicker/common"
)

func main() {
	// run pprof server
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	for {
		time.Sleep(common.TimeInterval)
	}
}

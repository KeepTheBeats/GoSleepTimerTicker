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

	var t *time.Timer
	for {
		t = time.NewTimer(common.TimeInterval)
		<-t.C
	}
}

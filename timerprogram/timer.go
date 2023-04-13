package main

import (
	"time"

	"GoSleepTimerTicker/common"
)

func main() {
	var t *time.Timer
	for {
		t = time.NewTimer(common.TimeInterval)
		<-t.C
	}
}

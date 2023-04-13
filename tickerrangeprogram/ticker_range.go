package main

import (
	"time"

	"GoSleepTimerTicker/common"
)

func main() {
	ticker := time.NewTicker(common.TimeInterval)
	defer ticker.Stop()
	for range ticker.C {
	}
}

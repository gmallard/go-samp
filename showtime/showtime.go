// time package, short demo

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "SHOWTIME ", log.Ldate|log.Lmicroseconds|log.Lshortfile)

func runTicker(quit chan bool) {
	ticker := time.NewTicker(1e9 * 3)
	q := false
	for {
		logger.Println("loop start")
		select {
			case ct := <- ticker.C:
				logger.Println(time.NanosecondsToLocalTime(ct))
			case q = <- quit:
				logger.Println("runTicker done")
				ticker.Stop()
				break
		}
		if q {
			break
		}
	}
	logger.Println("runTicker ends")
}

func main() {
	fmt.Println()
	logger.Println("Start...")

	logger.Println(time.LocalTime())

	time.Sleep(1e9 * 10) // 10 secs
	logger.Println(time.LocalTime())
	logger.Println(time.UTC())

	donechan := make(chan bool)
	go runTicker(donechan)

	logger.Println("start 30")
	time.Sleep(1e9 * 30) // 30 secs
	logger.Println("done 30")
	donechan <- true

	time.Sleep(1e9) // 1 sec

	lt := time.NanosecondsToLocalTime(time.Nanoseconds())
	ft := lt.Format(time.StampMicro)
	logger.Println(ft)
	logger.Println(time.NanosecondsToLocalTime(time.Nanoseconds()).Format(time.StampMicro))
	fmt.Println(time.NanosecondsToLocalTime(time.Nanoseconds()).Format(time.StampMicro))
	logger.Println("End...")
}

/*
	An experiment with using channels as a replacement for a formal "lock".
	Seems to work, no corruption of the updated resource.

	However, the first version of this program was racy (according to go's race
	detector).  This version is not (also according to go's race detector).

	This version eliminates hard coding each worker functionality as an
	individual function.  The maximum number of workers can be changed
	by a variable before a compile.

	It was hoped that this version whould take less total time than previous
	version(s).  To date that hope has *not* been realized.
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	donechan  = make(chan struct{})  // A global channel for signaling done
	startchan = make(chan struct{})  // A global channel to signal start
	reqlock   = make(chan chan bool) // Requests for the lock are sent on this channel
	freelock  = make(chan bool)      // A channel to signal that the lock is freed (unlocked)
	workerfin = make(chan bool)      // A channel to signal that each worker is done

	//
	wanted  = 1337 // An arbitrary number we want
	current = 0    // The shared resource. All quit when current == wanted

	// Change number of workers here
	maxworkers = 5      // max number of worker goroutines
	wcs        = []int{ // worker max wait time, each worker
		100, 300, 900,
		250, 750,
	}
)

const (
	//
	maxothersleep = 1500               // max wait time for "other" work
	waitatend     = maxothersleep + 50 // main will wait for all
)

//
func worker(r, wc int) {
	_ = <-startchan // Wait for the starting gun
	log.Printf("worker%d starts\n", r)
	var lc = 0               // Local update count
	var rp = make(chan bool) // reply channel, worker has lock
runLoop:
	for {
		select {
		case _ = <-donechan: // Check for quit
			break runLoop
		default:
		}
		log.Printf("worker%d requests lock %d\n", r, lc)

		// Lock request
		reqlock <- rp // Request the lock, send reply channel to lm
		_ = <-rp      // Worker has the lock
		// Lock obtained, critical section start
		log.Printf("worker%d has lock\n", r)
		if current == wanted {
			// doSleep("is EQ", r, wc) // wc is max sleep for this go routine
			freelock <- true // Free the lock
			log.Printf("worker%d lock is free - worker done %d\n", r, lc)
			// continue runLoop
			break runLoop
		}
		doSleep("after eqcheck", r, wc) // wc is max sleep for this go routine
		if current < wanted {
			lc++
			current++
			log.Printf("worker%d bumped %d\n", r, current)
		} else {
			log.Printf("worker%d run is done\n", r, current)
			freelock <- true
			break runLoop
		}
		freelock <- true // Free the lock
		// critical section end

		log.Printf("worker%d lock is free - A %d\n", r, lc)
		doSleep("otherwork", r, maxothersleep) // Do other things besides updating the shared resource
	}
	workerfin <- true
	log.Printf("worker%d ends %d", r, lc)
}

//
func lm() {
	log.Printf("lm starts\n")
	var lc = 0
	var rp = make(chan bool)
lmLoop:
	for {
		select {
		case _ = <-donechan:
			break lmLoop
		case rp = <-reqlock:
		}
		lc++
		log.Printf("lm received a lock request %d\n", lc)
		rp <- true
		log.Printf("lm sent locked OK %d\n", lc)
		_ = <-freelock
		log.Printf("lm lock is freed %d\n", lc)
	}
	log.Printf("lm ends %d\n", lc)
}

//
func main() {
	//
	log.Printf("%s\n", "main starts")
	go lm() // Start lock manager
	for i := 1; i <= maxworkers; i++ {
		go worker(i, wcs[i-1]) // next worker
	}
	log.Printf("%s\n", "main workers started")
	close(startchan)
	log.Printf("%s\n", "main workers notified to start")
	// time.Sleep(1000 * time.Millisecond)
	var rp = make(chan bool)
mainLoop:
	for {
		log.Printf("main loop starts %d\n", wanted)
		doSleep("mainloop", 0, -1)
		//
		reqlock <- rp // Request the lock
		_ = <-rp      // Wait for the lock

		// critical section start
		log.Printf("main loop has lock %d %d\n", current, wanted)
		if current == wanted {
			freelock <- true
			break mainLoop
		}
		log.Printf("main loop will free lock %d %d\n", current, wanted)
		freelock <- true // Free the lock
		// critical section end

	}
	log.Printf("main loop finished %d %d\n", current, wanted)
	close(donechan)
	for i := 1; i <= maxworkers; i++ {
		_ = <-workerfin
	}
	log.Printf("%s %d\n", "main current at end", current)
	log.Printf("%s\n", "main ends")
}

func doSleep(tag string, r, m int) {
	var n int
	if m < 0 {
		n = 250
	} else {
		n = rand.Intn(m)
	}
	ms := fmt.Sprintf("%d", n) + "ms"
	d, e := time.ParseDuration(ms)
	if e != nil {
		log.Fatalf("Bad Parse: %v\n", e)
	}
	log.Printf("%s doSleep runner%d sleep starts %v\n", tag, r, d)
	time.Sleep(d) // do some work with the lock held
	log.Printf("%s doSleep runner%d sleep done %v\n", tag, r, d)
}

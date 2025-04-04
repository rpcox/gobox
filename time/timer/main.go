package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"
	"time"
)

func sigHandler(sigChan chan os.Signal, done chan bool) {
	for sig := range sigChan {
		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			fmt.Println("signal: shutting down")
			close(done)
			return
		}
	}
}

func CalcTime(at string) (time.Time, time.Duration, error) {
	layout := "2006-01-02 15:04:05"

	var tp, nullTime time.Time
	tp, err := time.Parse(layout, at)
	now := time.Now()

	if err != nil {
		return nullTime, time.Duration(0), err
	}

	if tp.Before(now) {
		return nullTime, time.Duration(0), fmt.Errorf("submitted time is in the past")
	} else if tp.Equal(now) {
		return nullTime, time.Duration(0), fmt.Errorf("submitted time is equal to current time")
	}

	diff := tp.Sub(now)
	fmt.Println("first alert time :", tp)
	fmt.Printf("    current time : %v\n\n", now.UTC().Format(layout))

	return tp, diff, nil
}

func Work(work, done chan bool) {
	for {
		select {
		case <-work:
			fmt.Println("work, work, work")
		case <-done:
			fmt.Println("leaving work")
			return
		}
	}
}

func NewTimer(td time.Duration, done, alert, reset chan bool) {
	fmt.Println("* new timer")
	timer := time.NewTimer(td)

	go func() {
		select {
		case <-timer.C:
			fmt.Println("pop: time to inform the worker")
			timer.Stop()
			alert <- true
			reset <- true
			return
		case d := <-done:
			if d {
				fmt.Println("exiting timer")
				return
			}
		}
	}()
}

func CalcInterval(interval string) int {
	r, err := regexp.Compile(`^(\d+)(\w)`)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	matches := r.FindStringSubmatch(interval)
	m := 0

	switch matches[2] {
	case `s`:
		m = 1
	case `m`:
		m = 60
	case `h`:
		m = 3600
	case `d`:
		m = 86400
	}

	i, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return m * i
}

func main() {
	_at := flag.String("at", "", "time to alert. format = '2006-01-02 15:04:05'")
	_interval := flag.String("interval", "", "N[smhd] s = secs, m = mins, ...")
	flag.Parse()

	at, td, err := CalcTime(*_at)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	interval := CalcInterval(*_interval)

	alert := make(chan bool)
	done := make(chan bool)
	reset := make(chan bool)
	work := make(chan bool)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go sigHandler(sigChan, done)

	go Work(work, done)
	NewTimer(td, done, alert, reset)

	for {
		select {
		case <-done:
			fmt.Println("exiting main")
			return
		case <-alert:
			now := time.Now().Format(time.RFC3339Nano)
			fmt.Println("alarm time:", now)
			fmt.Println("pass the word to the worker")
			work <- true
		case <-reset:
			at = at.Add(time.Duration(interval) * time.Second)
			now := time.Now()
			td = at.Sub(now)
			NewTimer(td, done, alert, reset)
		}
	}

}

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
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

func ValidateTime(at string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"

	now := time.Now()
	fmt.Println("  now:", now.UTC().Format(layout))

	var tp, nullTime time.Time
	tp, err := time.Parse(layout, at)
	fmt.Println("alert:", tp)

	if err != nil {
		return nullTime, err
	}

	if tp.Before(now) {
		return nullTime, fmt.Errorf("submitted time is in the past")
	} else if tp.Equal(now) {
		return nullTime, fmt.Errorf("submitted time is equal to current time")
	}

	return tp, nil
}

func Alert(alert, done chan bool) {
	for {
		select {
		case <-alert:
			fmt.Println("alert: time to do something")
		case <-done:
			fmt.Println("exiting alert()")
			return
		}
	}
}

func NewTimer(t time.Time, done, alert, reset chan bool) {
	fmt.Println("new timer")
	timer := time.NewTimer(time.Duration(60) * time.Second)

	go func() {
		select {
		case <-timer.C:
			fmt.Println("alert: time to do something")
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
func main() {
	_at := flag.String("at", "", "time to alert. format = '2006-01-02 15:04:05'")
	//_interval := flag.String("interval", "", "N[smhd] s = secs, m = mins, ...")
	flag.Parse()

	t, err := ValidateTime(*_at)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan bool)
	alert := make(chan bool)
	reset := make(chan bool)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go sigHandler(sigChan, done)

	NewTimer(t, done, alert, reset)

	for {
		select {
		case <-done:
			fmt.Println("exiting main")
			return
		case <-alert:
			now := time.Now().Format(time.RFC3339)
			fmt.Println(now, "time to do something")
		case <-reset:
			t = t.Add(time.Duration(2) * time.Minute)
			NewTimer(t, done, alert, reset)
		}
	}

}

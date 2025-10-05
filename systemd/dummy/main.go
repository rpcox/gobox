package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	tool    string = "dummy"
	version string = "v0.2.0"
)

func ElevateCPU(done chan bool, id int) {

	fmt.Fprintf(os.Stderr, "[%d] enter cpu elevation\n", id)
	for {
		select {
		case <-done:
			fmt.Fprintf(os.Stderr, "[%d] exit cpu elevation\n", id)
			return
		default:
			// Perform a simple, non-blocking calculation
			_ = 12345 * 67890 / 98765 % 43210
		}
	}
}

func JournalMessage(exit, mark chan bool) {
	msgCount := 1

	for {
		select {
		case <-mark:
			fmt.Fprintf(os.Stderr, "mark %d\n", msgCount)
			msgCount++
		case <-exit:
			fmt.Fprintf(os.Stderr, "exit JournalMessage()\n")
			return
		}
	}
}

func main() {
	_cpu_check := flag.Bool("cpu-check", false, "Determine the number of system CPUs")
	_speed := flag.Float64("speed", 1.0, "The fraction of CPUs to use")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())
	if *_cpu_check {
		fmt.Fprintf(os.Stderr, "cpu count = %d", runtime.NumCPU())
		os.Exit(0)
	}

	max_cores := runtime.NumCPU()
	if *_speed < 0 || *_speed > 1.0 {
		fmt.Fprintf(os.Stderr, "speed must be > 0 and <= 1\n")
		os.Exit(1)
	} else {
		max_cores = int(math.Ceil(float64(runtime.NumCPU()) / *_speed))
		if max_cores < 1 {
			max_cores = 1
		}
	}

	fmt.Fprintf(os.Stderr, "%s %s pid=%d\n", tool, version, os.Getpid())
	fmt.Fprintf(os.Stderr, "cpu_count=%d\n", runtime.NumCPU())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

	cpuElevated := false
	done := make(chan bool, runtime.NumCPU())
	mark := make(chan bool)
	exit := make(chan bool)
	go JournalMessage(exit, mark)
	goRoutineCount := 0

	for {
		select {
		case sig := <-sigChan:
			if sig == syscall.SIGUSR1 {
				fmt.Fprintf(os.Stderr, "sigusr1: toggle cpu utilization\n")
				if !cpuElevated {
					for i := 1; i <= max_cores; i++ {
						go ElevateCPU(done, i)
					}
					cpuElevated = true
				} else {
					for i := 1; i <= max_cores; i++ {
						done <- true
					}
					cpuElevated = false
				}
			} else if sig == syscall.SIGUSR2 {
				fmt.Fprintf(os.Stderr, "sigusr2: mark\n")
				mark <- true
			} else if sig == syscall.SIGTERM || sig == syscall.SIGINT {
				close(done)
				close(sigChan)
				close(mark)
				close(exit)
				return
			}
		default:
			time.Sleep(1 * time.Second)
			if goRoutineCount%60 == 0 {
				fmt.Fprintf(os.Stderr, "goroutine count %d\n", runtime.NumGoroutine())
				goRoutineCount = 0
			}
			goRoutineCount++
		}
	}
}

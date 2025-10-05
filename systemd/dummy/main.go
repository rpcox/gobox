package main

import (
	"fmt"
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

func JournalMessage(done chan bool, mark chan bool) {
	msgCount := 1

	for {
		select {
		case <-mark:
			fmt.Fprintf(os.Stderr, "mark %d\n", msgCount)
			msgCount++
		case <-done:
			fmt.Fprintf(os.Stderr, "exit JournalMessage()\n")
			return
		}
	}
}

func main() {

	fmt.Fprintf(os.Stderr, "%s %s pid=%d\n", tool, version, os.Getpid())
	runtime.GOMAXPROCS(runtime.NumCPU())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	cpuElevated := false
	done := make(chan bool, runtime.NumCPU())
	mark := make(chan bool)

	for {
		select {
		case sig := <-sigChan:
			if sig == syscall.SIGUSR1 {
				fmt.Fprintf(os.Stderr, "sigusr1: toggle cpu utilization\n")
				if !cpuElevated {
					for i := 0; i < runtime.NumCPU(); i++ {
						go ElevateCPU(done, i)
					}
					cpuElevated = true
				} else {
					for i := 0; i < runtime.NumCPU(); i++ {
						done <- true
					}
				}
			} else if sig == syscall.SIGUSR2 {
				fmt.Fprintf(os.Stderr, "sigusr2: mark\n")
				mark <- true
			} else if sig == syscall.SIGTERM || sig == syscall.SIGINT {
				close(done)
				close(sigChan)
				close(mark)
				time.Sleep(1 * time.Second)
				return
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

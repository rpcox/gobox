package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coreos/go-systemd/journal"
)

var (
	tool string = "dummy"
	version string = "v0.1.0"
)


func main() {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	msgCount := 1

	for {
		select {
		case <- sigChan:
			return
		default:
			msg := fmt.Sprintf("dummy message [%d]", msgCount)
			err := journal.Send(msg, journal.PriInfo, nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}

			msgCount++
			time.Sleep(30 * time.Second)
		}
	}
}

// A cmdline filter for tcpdump to determine bytes per second on a system network interface
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"
	"time"
)

func sigHandler(sigChan chan os.Signal, done chan any) {
	for sig := range sigChan {
		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			close(done)
			return
		}
	}
}

func main() {
	_interval := flag.Float64("interval", 10, "capture interval in seconds")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	done := make(chan any)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go sigHandler(sigChan, done)

	// sudo tcpdump -i eth0 -l -e -n
	// 17:09:39.719946 dc:a6:32:9d:c0:f9 > 20:6d:31:02:0e:48, ethertype IPv4 (0x0800), length 566: ...
        // 17:09:39.723143 20:6d:31:02:0e:48 > dc:a6:32:9d:c0:f9, ethertype IPv4 (0x0800), length 66: ...
        // 17:09:39.823780 dc:a6:32:9d:c0:f9 > 20:6d:31:02:0e:48, ethertype IPv4 (0x0800), length 566: ...
        // 17:09:39.826450 20:6d:31:02:0e:48 > dc:a6:32:9d:c0:f9, ethertype IPv4 (0x0800), length 66: ...
	// ...
	r := regexp.MustCompile(`^[^ ]+ [^ ]+ . [^ ]+ [^ ]+ [^ ]+ [^ ]+ length (\d+):`)

	var n, bytes int64
	start := time.Now()
	for {
		select {
		case <-done:
			os.Exit(0)
		default:
			in, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stdout, "%s\n", err)
				os.Exit(1)
			}

			m := r.FindStringSubmatch(in)
			if len(m) > 0 {
				n, err = strconv.ParseInt(m[1], 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				bytes = bytes + n
				elapsed := time.Since(start).Seconds()

				if elapsed > *_interval {
					// bytes per second = float64(bytes) / float64(elapsed)
					fmt.Fprintf(os.Stdout, "%15.2f Bps\n", float64(bytes) / float64(elapsed))
					start = time.Now()
					bytes = 0
				}
			}
		}
	}
}

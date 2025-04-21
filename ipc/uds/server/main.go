package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func echoHandler(conn net.Conn, i int, done chan bool, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()
	defer fmt.Printf("handler[%d]: returning\n", i)

	for {
		select {
		case <-done:
			return
		default:
			buf := make([]byte, 512)
			n, err := conn.Read(buf)
			if err == io.EOF {
				fmt.Fprintf(os.Stderr, "handler[%d]: connection closed\n", i)
				return
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "handler[%d]: read error: %s\n", i, err)
				return
			}

			data := buf[0:n]
			fmt.Printf("handler[%d]: %s\n", i, string(data))

			_, err = conn.Write(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "handler[%d]: write error: %v\n", i, err)
				return
			}
		}
	}
}

func sigHandler(sigChan chan os.Signal, done chan bool, nl net.Listener) {

	for sig := range sigChan {
		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			fmt.Printf("sigHandler: exiting\n")
			err := nl.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "sigHandler: error closing listener: %v\n", err)
			}
			close(done)
			return
		}

	}
}

func StartReceiver(socketName string) net.Listener {
	os.Remove(socketName)
	nl, err := net.Listen("unix", socketName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "listen error: %v\n", err)
		os.Exit(2)
	}

	return nl
}

func main() {
	_socketPath := flag.String("socket-path", "/tmp/echo=", "Specify the full path of the domain socket")
	flag.Parse()

	done := make(chan bool)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	nl := StartReceiver(*_socketPath)
	defer nl.Close()
	go sigHandler(sigChan, done, nl)

	i := 0
	var wg sync.WaitGroup
	for {
		fmt.Println("accept")
		conn, err := nl.Accept()
		if err != nil {
			// accept unix /tmp/echo=: use of closed network connection
			// due to sigHandler exit
			break
		}

		wg.Add(1)
		go echoHandler(conn, i, done, &wg)
		i++
	}

	wg.Wait()
}

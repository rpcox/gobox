package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"math/big"
	"net"
	"os"
	"os/signal"
	"regexp"
	"sync"
	"syscall"
	"time"
)

func generateRandomString(length int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[num.Int64()]
	}
	return string(result), nil
}

func workHandler(done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	id, _ := generateRandomString(8)
	handle := `work-` + id
	defer fmt.Fprintf(os.Stderr, "%s: exiting\n", handle)

	for {
		select {
		case <-done:
			return
		default:
			time.Sleep(time.Duration(2) * time.Second)
			fmt.Printf("%s: working\n", handle)
		}
	}
}

func sigHandler(sigChan chan os.Signal, done chan bool) {

	for sig := range sigChan {
		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			fmt.Printf("sigHandler: exiting\n")
			_, ok := <-done
			if ok {
				close(done)
			}
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

func MessageHandler(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	buf := make([]byte, 64)
	n, err := conn.Read(buf)
	if err == io.EOF {
		return []byte{}, fmt.Errorf("message: connection closed: %s", err)
	} else if err != nil {
		return []byte{}, err
	}

	msg := buf[0:n]

	ackFail := false
	_, err = conn.Write([]byte(`ACK`))
	if err != nil {
		fmt.Fprintf(os.Stderr, "message: client ack fail: %s\n", err)
		ackFail = true
	}

	// if expecting a large volume of traffic there is a need for a more
	// complex regex, use Compile
	match, err := regexp.Match(`^(pause|resume|stop|work)`, msg)
	if err != nil {
		return []byte{}, fmt.Errorf("message: regex fail: %s", err)
	}

	if match {
		return msg, nil
	}

	if !ackFail {
		conn.Write([]byte(`bad command submitted`))
	}

	return []byte{}, fmt.Errorf("messsage: bad command received: '%s'", string(msg))
}

func main() {
	_socketPath := flag.String("socket-path", "/tmp/work=", "Specify the full path of the domain socket")
	flag.Parse()

	done := make(chan bool)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	nl := StartReceiver(*_socketPath)
	defer nl.Close()
	go sigHandler(sigChan, done)

	paused := false
	var wg sync.WaitGroup
	for {
		fmt.Println("accept")
		conn, err := nl.Accept()
		if err != nil {
			// Listener is toast. Most likely error is below
			// accept unix /tmp/echo=: use of closed network connection
			break
		}

		msg, err := MessageHandler(conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			continue
		}

		switch {
		case bytes.Equal(msg, []byte(`work`)):
			if !paused {
				wg.Add(1)
				go workHandler(done, &wg)
			}
		case bytes.Equal(msg, []byte(`pause`)):
			paused = true
			fmt.Fprintf(os.Stderr, "pausing for new work\n")
		case bytes.Equal(msg, []byte(`resume`)):
			paused = false
			fmt.Fprintf(os.Stderr, "resuming for new work\n")
		case bytes.Equal(msg, []byte(`stop`)):
			close(done)
			nl.Close()
		}
	}

	wg.Wait()
}

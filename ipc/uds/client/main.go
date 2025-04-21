package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

// read the returns
func Reader(r net.Conn) {
	defer fmt.Printf("reader: returning\n")

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}

		fmt.Printf("returned: %s\n", string(buf[0:n]))
	}
}

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

func main() {
	_socketPath := flag.String("socket-path", "/tmp/echo=", "Specify the full path of the domain socket")
	flag.Parse()

	conn, err := net.Dial("unix", *_socketPath)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer conn.Close()

	go Reader(conn)
	for {
		msg, _ := generateRandomString(10)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Fprintf(os.Stderr, "write error: %s\n", err)
			break
		}
		fmt.Println("send:", msg)
		time.Sleep(1 * time.Second)
	}
}

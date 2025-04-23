package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	_socketPath := flag.String("socket-path", "/tmp/work=", "Specify the full path of the domain socket")
	flag.Parse()
	cmd := flag.Args()

	conn, err := net.Dial("unix", *_socketPath)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(cmd[0]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "write error: %s\n", err)
		return
	}

	buf := make([]byte, 128)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %s\n", err)
		return
	}

	fmt.Fprintln(os.Stderr, "ok")
}

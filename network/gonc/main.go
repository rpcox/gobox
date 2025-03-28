// Send text strings over the network to a server.
// The intent was to send preformatted syslog messages for regex testing.
// Security folks often don't think mere mortals can handle netcat.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

const _version = "0.1"

func Version(b bool) {
	if b {
		_tool := filepath.Base(os.Args[0])
		fmt.Printf("%s v%s\n", _tool, _version)
	}
}

type Transport struct {
	msgQ    chan *[]byte
	address string
	network string
}

func NewTransport(network, address string, dataQCap, doneQCap int) *Transport {
	var t Transport
	t.network = network
	t.address = address
	t.msgQ = make(chan *[]byte, dataQCap)

	return &t
}

func (t *Transport) Start(wg *sync.WaitGroup) {
	conn, err := net.Dial(t.network, t.address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for data := range t.msgQ {

		if data == nil {
			conn.Close()
			wg.Done()
			return
		}

		_, err = conn.Write(*data)
		if err != nil {
			log.Println(err)
			continue
		}

		data = nil

	}
}

func main() {
	address := flag.String("ip", "", "Specify the destination IP or name. See net.Dial()")
	port := flag.Int("port", 514, "Specify the destination port. See net.Dial()")
	protocol := flag.String("proto", "tcp", "Specify the protocol udp, udp4, udp6, tcp, tcp4, tcp6.  See net.Dial()")
	dataQCap := flag.Int("data-chan-cap", 100, "Specify the data capacity")
	version := flag.Bool("version", false, "Display the version and exit")

	flag.Parse()

	Version(*version)

	transport := NewTransport(*protocol, *address+":"+strconv.Itoa(*port), *dataQCap, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go transport.Start(&wg)

	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 { // data is being piped to stdin

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			b := scanner.Bytes()
			transport.msgQ <- &b
		}

		var b *[]byte
		transport.msgQ <- b

	} else {
		files := flag.Args()
		if len(files) == 0 {
			log.Fatal("no files identified.  nothing to do.  exiting.")
		} else {
			for _, f := range files {
				fh, err := os.Open(f)
				if err != nil {
					log.Println(err)
					continue
				}
				scanner := bufio.NewScanner(fh)
				for scanner.Scan() {
					b := scanner.Bytes()
					transport.msgQ <- &b
				}
			}

			var b *[]byte
			transport.msgQ <- b
		}
	}

	wg.Wait()
	close(transport.msgQ)
}

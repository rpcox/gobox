// quick simple file server to move things
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sigHandler(sigChan chan os.Signal, server *http.Server) {
	for sig := range sigChan {
		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			log.Println("signal: shutting down")
			ctx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
			defer shutdownRelease()

			if err := server.Shutdown(ctx); err != nil {
				log.Fatalf("shutdown error: %v", err)
			}
			return
		}
	}
}


func main() {
	_ip := flag.String("ip", "0.0.0.0", "IP to bind on")
	_port := flag.Int("port", 8000, "Port to serve on")
	_directory := flag.String("d", ".", "The directory to serve")
	flag.Parse()

	if _, err := os.Stat(*_directory); os.IsNotExist(err) {
		log.Fatalf("directory '%s' does not exist", *_directory)
	}

	log.Printf("serving %s on port: %d\n", *_directory, *_port)
	ip := net.ParseIP(*_ip)
	if ip == nil {
		log.Fatalf("not a valid IP: '%s'\n", *_ip)
	}

	http.Handle("/", http.FileServer(http.Dir(*_directory)))

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", *_ip, *_port),
		Handler: nil,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go sigHandler(sigChan, server)


	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// Simple middleware nesting test
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Middleware func(http.Handler) http.Handler

// Next applies middleware in order. The first middleware becomes the outermost layer.
func Nest(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		fmt.Fprintf(os.Stdout, "+ enter Logging: %s - - %s \"%s %s %s %s\"\n",
			r.RemoteAddr, t0.Format(time.RFC822Z), r.Method, r.URL.Path, r.Proto, r.UserAgent())
		next.ServeHTTP(w, r)
		fmt.Fprintf(os.Stdout, "- exit Logging: %v\n", time.Since(t0))
	})
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, " + enter Auth\n")
		next.ServeHTTP(w, r)
		fmt.Fprintf(os.Stdout, " - exit Auth\n")
	})
}

func HeaderSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "  + enter HeaderSet\n")
		w.Header().Set("X-Server", "Go") // observed by receiver
		next.ServeHTTP(w, r)
		fmt.Fprintf(os.Stdout, "  - exit HeaderSet\n")
	})
}

// primary/main application handler for /
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, "   + enter mainHandler\n")
	w.Write([]byte("hello chain"))
	fmt.Fprintf(os.Stdout, "   - exit mainHandler\n")
}

func main() {
	mux := http.NewServeMux()

	finalHandler := Nest(
		http.HandlerFunc(mainHandler),
		Logging,   // Runs 1st
		Auth,      // Runs 2nd
		HeaderSet, // Runs 3rd
	)

	mux.Handle("/", finalHandler)

	fmt.Fprintf(os.Stdout, "server starting on :8080\n")
	http.ListenAndServe(":8080", mux)
}

// Simple middleware nesting test
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"slices"
	"time"
)

type contextKey string

const user contextKey = `userID`
const panik contextKey = `panic`
const validUser = `johndoe`

type Middleware func(http.Handler) http.Handler

// Next applies middleware in order. The first middleware becomes the outermost layer.
func Nest(handler http.Handler, middlewares ...Middleware) http.Handler {
	//for i := len(middlewares) - 1; i >= 0; i-- {
	for i, _ := range slices.Backward(middlewares) {
		handler = middlewares[i](handler)
	}
	return handler
}

func Recovery(next http.Handler) http.Handler {
	t0 := time.Now()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "+ enter Recovery\n")
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintf(os.Stderr, "%v:\nstack:\n%s\n", err, string(debug.Stack()))

				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
		fmt.Fprintf(os.Stdout, "- exit Recovery: %v\n", time.Since(t0))
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		fmt.Fprintf(os.Stdout, "+ enter Logging: %s - - %s \"%s %s %s %s\"\n",
			r.RemoteAddr, t0.Format(time.RFC822Z), r.Method, r.URL.Path, r.Proto, r.UserAgent())
		next.ServeHTTP(w, r)
		userID, ok := r.Context().Value(user).(string)
		if ok { // note: this won't fire. context started in verify json
			fmt.Fprintf(os.Stdout, " user action for %s completed\n", userID)
		}
		fmt.Fprintf(os.Stdout, "- exit Logging\n")
	})
}

type UserRequest struct {
	Name  string `json:"name"`
	Panic string `json:"panic"`
}

func VerifyJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, " + enter VerifyJson\n")
		if r.Method == http.MethodPost {
			var req UserRequest

			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				s := fmt.Sprintf("error: malformed json: %v", err)
				fmt.Fprintf(os.Stdout, " + %s\n", s)
				http.Error(w, s, http.StatusBadRequest) // sysntax/structure error
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, user, req.Name)
			ctx = context.WithValue(ctx, panik, req.Panic)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		fmt.Fprintf(os.Stdout, " - exit VerifyJson\n")
	})
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "  + enter Auth\n")
		userID, ok := r.Context().Value(user).(string)
		if !ok {
			s := `error: missing user ID`
			fmt.Fprintf(os.Stdout, " + %s\n", s)
			http.Error(w, s, http.StatusUnprocessableEntity) // violates logic and business rules
			return
		}

		if userID != validUser {
			s := `unauthorized user`
			fmt.Fprintf(os.Stdout, "  + %s\n", s)
			http.Error(w, s, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
		fmt.Fprintf(os.Stdout, "  - exit Auth\n")
	})
}

func HeaderSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "   + enter HeaderSet\n")
		w.Header().Set("X-Server", "Go") // observed by receiver
		next.ServeHTTP(w, r)
		fmt.Fprintf(os.Stdout, "   - exit HeaderSet\n")
	})
}

// primary/main application handler for /
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, "    + enter mainHandler\n")
	userID, _ := r.Context().Value(user).(string)
	s := fmt.Sprintf("usage valid for: %s\n", userID)
	w.Write([]byte("hello chain\n"))
	w.Write([]byte(s))
	panicSwitch, _ := r.Context().Value(panik).(string)
	if panicSwitch == "true" {
		panic(`throwing panic`)
	}
	fmt.Fprintf(os.Stdout, "    - exit mainHandler\n")
}

func main() {
	mux := http.NewServeMux()

	finalHandler := Nest(
		http.HandlerFunc(mainHandler),
		Recovery,
		Logging, // Runs 1st
		VerifyJson,
		Auth,      // Runs 2nd
		HeaderSet, // Runs 3rd
	)

	mux.Handle("/", finalHandler)

	fmt.Fprintf(os.Stdout, "server starting on :8080\n")
	http.ListenAndServe(":8080", mux)
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var healthy = true

func main() {
	sleep()

	http.Handle("/", http.HandlerFunc(root))

	port := 3333
	fmt.Printf("listening on port %v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	must(err)
}

func sleep() {
	sleepTimeString, set := os.LookupEnv("SLEEP")
	if !set {
		must(errors.New("env variable SLEEP is not set"))
	}
	sleepTime, err := strconv.Atoi(sleepTimeString)
	must(err)
	time.Sleep(time.Second * time.Duration(sleepTime))
}

func root(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()

	if !healthy {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = fmt.Fprintf(w, "%v is not healthy\n", hostname)
		return
	}

	if r.URL.Path == "/makeunhealthy" {
		healthy = false
		_, _ = fmt.Fprintf(w, "%v set to unhealthy...\n", hostname)
		return
	}

	w.WriteHeader(501)
	_, _ = fmt.Fprintf(w, "%v is healthy\n", hostname)
}

func must(err error) {
	if err != nil {
		// ignoring print error and written bytes count
		_, _ = fmt.Fprintf(os.Stderr, "%s: %s\n", time.Now().Format(time.RFC850), err)
		os.Exit(1)
	}
}

// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	version := "1.0.0"
	fmt.Fprintf(w, "<title>Koala %s</title>", version)
	fmt.Fprintf(w, "<pre>")
	fmt.Fprintf(w, "{\"Method\": \"%s\", \"Path\":\"%s\", \"time\":\"%s\", \"Hostname\":\"%s\", \"Version\":\"%s\"}",
		r.Method,
		r.URL.Path,
		time.Now().Format("Mon Jan 2 15:04:05 2006"),
		host,
		version,
	)
	log.Printf("{\"Method\": \"%s\", \"Path\":\"%s\", \"time\":\"%s\", \"Hostname\":\"%s\", \"Version\":\"%s\"}",
		r.Method,
		r.URL.Path,
		time.Now().Format("Mon Jan 2 15:04:05 2006"),
		host,
		version,
	)
}

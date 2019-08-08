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
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Koala Version --> %s \nHostname --> %s", "1.0.0", host)
	log.Printf("{\"Request\":\"%s\", \"time\":\"%s\", \"Hostname\":\"%s\"}",
		r.URL.Path,
		time.Now().Format("Mon Jan 2 15:04:05 2006"),
		host,
	)
}

// Copyright 2014 The gofsh Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":6060", "listening address")
	root = flag.String("root", ".", "root folder")
)

func main() {
	flag.Parse()
	log.Printf("Listening on %s (root=%s)\n", *addr, *root)
	fileHandler := http.FileServer(http.Dir(*root))
	if err := http.ListenAndServe(*addr, logHandler(fileHandler)); err != nil {
		log.Fatalln("Error:", err)
	}
}

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s\n", r.RemoteAddr, r.Method, r.RequestURI)
		handler.ServeHTTP(w, r)
	})
}

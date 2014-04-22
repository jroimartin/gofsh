// Copyright 2014 The gofsh Authors.  All rights reserved.
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
	if err := http.ListenAndServe(*addr, http.FileServer(http.Dir(*root))); err != nil {
		log.Fatalln("Error:", err)
	}
}

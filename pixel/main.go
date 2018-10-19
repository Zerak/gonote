// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			idx := x + y + 1
			fix := ""
			if y/10 == 0 {
				fix = "\n"
			}
			fmt.Printf("%v x:%v y:%v %v", idx, x, y, fix)
		}
	}
}

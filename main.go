package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"log"
	"net/http"
	"time"
)

const VERSION = `0.3`

var (
	folder  = flag.String("folder", "./static", "Static Folder Path")
	port    = flag.Int("port", 8082, "set port (default 8082) ")
	version = flag.Bool("v", false, "display version")
	help    = flag.Bool("h", false, "display usage")
)

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		elapsed := time.Since(start)

		url := r.URL.Path
		method := r.Method

		log.Println(fmt.Sprintf("Elapsed time %06s \t %7s \t %s", elapsed, method, url))
	})
}

func main() {

	flag.Parse()
	if *version {
		fmt.Printf("simple-http %s\n", VERSION)
		return
	}

	if *help {
		fmt.Println("simple-http Help")
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("simple-http -v                             => Display version")
		fmt.Println("simple-http -h                             => Display usage")
		fmt.Println("simple-http -folder=./static_folder        => Static Folder Path (default ./static)")
		fmt.Println("simple-http -port=8082                     => Port ( default 8082) ")
		fmt.Println("simple-http -folder=./static -port=8082    => Start s simple http server for ./static folder via 8082 port")

		return
	}

	rootDir := http.Dir(*folder)
	filerServer := http.FileServer(rootDir)

	http.Handle("/", middleware(filerServer))

	log.Printf("Simple-HTTP started to serve '%s' folder on port: %d\n", *folder, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	log.Println("Stoppping app")
}

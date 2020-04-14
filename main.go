package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"log"
	"net/http"
	"time"

	tr "github.com/cratonica/trayhost"
	"github.com/rayyildiz/simple-http/icon"
)

// VERSION of simple-http.
const VERSION = `0.5.0`

var (
	folder  = flag.String("folder", ".", "Static Folder Path")
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
		fmt.Printf("simple-http v%s Help \n", VERSION)
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("simple-http -v                                => Display version")
		fmt.Println("simple-http -h                                => Display usage")
		fmt.Println("simple-http -folder=./static_folder        => Static Folder Path (default  folder is .)")
		fmt.Println("simple-http -port=8082                        => Port (default 8082) ")
		fmt.Println("simple-http -folder=./static -port=8082    => Start s simple http server for ./static folder via 8082 port")

		return
	}

	f, err := os.OpenFile("access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.SetOutput(f)

	srv := startHTTPServer(*port, *folder)
	defer srv.Close()

	fmt.Printf("simple-http v%s started http://localhost:%d for '%s' folder. Access logs are saved in '%s' file.\n", VERSION, *port, *folder, "access.log")

	runtime.LockOSThread()

	go func() {
		tr.SetUrl(fmt.Sprintf("http://localhost:%d", *port))
	}()

	tr.EnterLoop(fmt.Sprintf("Simple Http v%s", VERSION), icon.Data)

	log.Println("Stoppping app")
	err = srv.Close()
	if err != nil {
		log.Printf("stopping error, %v", err)
		os.Exit(1)
	}
}

func startHTTPServer(port int, folder string) *http.Server {
	srv := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	rootDir := http.Dir(folder)
	filerServer := http.FileServer(rootDir)

	http.Handle("/", middleware(filerServer))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("simple-http: start server error: %s", err)
			fmt.Printf("simple-http: start server error: %s", err)
			os.Exit(1)
		}
	}()
	log.Printf("Simple-HTTP v%s started to serve '%s' folder on port: %d", VERSION, folder, port)

	// returning reference so caller can call Shutdown()
	return srv
}

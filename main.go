package main

import (
	"github.com/gin-gonic/gin"
	"flag"
	"fmt"
)

const VERSION = `0.1`

var (
	folder  = flag.String("folder", "./static", "Static Folder Path")
	port    = flag.Int("port", 8082, "display version info and exit")
	version = flag.Bool("v", false, "display version info and exit")
	help    = flag.Bool("h", false, "display usage")
)

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

	router := gin.Default()
	router.Static("/", *folder)
	router.Run(fmt.Sprintf(":%d", *port))
}

Simple Http Server
===

[![Build Status](http://img.shields.io/travis/rayyildiz/simple-http.svg?style=flat-square)](https://travis-ci.org/rayyildiz/simple-http)
[![Go Report Card](https://goreportcard.com/badge/github.com/rayyildiz/simple-http)](https://goreportcard.com/report/github.com/rayyildiz/simple-http)

Starts a simple http server to serve static files. Default port is ```8082``` and default folder is ```.``` (current folder)

Checkout and build:

    go get github.com/rayyildiz/simple-http
    go install github.com/rayyildiz/simple-http

Usage
---

| Parameter                               | Description                                                  |
|:----------------------------------------|:-------------------------------------------------------------|
| simple-http -v                          | Display version                                              |
| simple-http -h                          | Display usage                                                |
| simple-http -folder=./static_folder     | Static Folder Path (Default .)                               |
| simple-http -port=8082                  | Port (Default 8082)                                          |
| simple-http -folder=./static -port=8082 | Start s simple http server for ./static folder via 8082 port |


Credits
---

* https://github.com/cratonica/trayhost

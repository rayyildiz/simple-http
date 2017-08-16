Simple Http Server
===


Starts a simple http server to serve static files. Default port is ```8082``` and default folder is ```./static```

Checkout and build: 
    
    go get github.com/rayyildiz/simple-http
    go install github.com/rayyildiz/simple-http


Usage 
---


| Parameter                               | Description                                                  |
|:----------------------------------------|:-------------------------------------------------------------|
| simple-http -v                          | Display version                                              |
| simple-http -h                          | Display usage                                                |
| simple-http -folder=./static_folder     | Static Folder Path (Default ./static)                        |
| simple-http -port=8082                  | Port (Default 8082)                                          |
| simple-http -folder=./static -port=8082 | Start s simple http server for ./static folder via 8082 port |

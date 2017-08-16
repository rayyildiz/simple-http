GO=go.exe

windows:
	CGO_ENABLED=0 GOOS=windows 	$(GO) build -o ./dist/windows/simple-http.exe .

linux:
	CGO_ENABLED=0 GOOS=linux 	$(GO) build -o ./dist/linux/simple-http .

mac:
	CGO_ENABLED=0 GOOS=darwin 	$(GO) build -o ./dist/macos/simple-http .

bsd:
	CGO_ENABLED=0 GOOS=freebsd 	$(GO) build -o ./dist/freebsd/simple-http.
	CGO_ENABLED=0 GOOS=netbsd 	$(GO) build -o ./dist/netbsd/simple-http.
	CGO_ENABLED=0 GOOS=openbsd 	$(GO) build -o ./dist/openbsd/simple-http .

solaris:
	CGO_ENABLED=0 GOOS=solaris 	$(GO) build -o ./dist/solaris/simple-http .

all: windows linux mac bsd solaris
	@echo "Compiled for all pltaforms"
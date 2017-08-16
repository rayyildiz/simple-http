$env:GOOS = "windows"
$env:GOOS
rm ./dist/windows/simple-http.exe
go get ./...
go build -o ./dist/windows/simple-http.exe .
echo "Created Windows App"


$env:GOOS = "linux"
$env:GOOS
rm ./dist/linux/simple-http
go get ./...
go build -o ./dist/linux/simple-http .
echo "Created linux App"


$env:GOOS = "darwin"
$env:GOOS
rm ./dist/macos/simple-http
go get ./...
go build -o ./dist/macos/simple-http .
echo "Created darwin App"
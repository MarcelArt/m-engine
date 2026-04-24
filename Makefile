flappy:
	@go run examples/flappy/main.go

dev-flappy:
	@air examples/flappy/main.go

build-flappy-windows:
	@GOOS=windows GOARCH=amd64 go build -o builds/flappy/flappy.exe ./examples/flappy

build-flappy-linux:
	@GOOS=linux GOARCH=amd64 go build -o builds/flappy/flappy.x86_64 ./examples/flappy

# build-flappy-mac:
# 	@GOOS=darwin GOARCH=arm64 go build -o builds/flappy/flappy ./examples/flappy

build-flappy: build-flappy-windows build-flappy-linux
	
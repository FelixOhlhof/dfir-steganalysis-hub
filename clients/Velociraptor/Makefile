BINARY_NAME=grpc-client

build:
	set GOOS=windows&& set GOARCH=amd64&& go build -o $(BINARY_NAME)-windows-amd64.exe
	set GOOS=linux&& set GOARCH=amd64&& go build -o $(BINARY_NAME)-linux-amd64

clean:
	go clean
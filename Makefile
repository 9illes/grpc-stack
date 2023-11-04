protoc:
	protoc *.proto --go_out=./ --go-grpc_out=./

test:
	go test -v ./...

build:
	go build -o build/stack main.go
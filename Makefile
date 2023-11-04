CONTAINER_ENGINE=docker
CONTAINER_IMAGE=stack

protoc:
	protoc *.proto --go_out=./ --go-grpc_out=./

test:
	go test -v ./...

build:
	go build -o build/stack main.go

clean:
	rm -rf build

container:
	${CONTAINER_ENGINE} build -t ${CONTAINER_IMAGE} .

run:
	${CONTAINER_ENGINE} run --rm -p 8080:8080 ${CONTAINER_IMAGE}
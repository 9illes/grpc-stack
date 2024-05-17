# GRPC Stack

[![Go Reference](https://pkg.go.dev/badge/github.com/9illes/grpc-stack.svg)](https://pkg.go.dev/github.com/9illes/grpc-stack)

Sandbox project for discovering GRPC

## TODO

* Use a separate container for the build
* Optimize container size
* Plug a GUI (Giu)

## Usage

Help

```sh
go run . -h
```

Start server (default localhost:8080)

```sh
go run . serve [-i=hostname] [-p port]
```

Push int to the stack

```sh
go run . push -n 123
```

Pop int from the stack

```sh
go run . pop
```

## Setup

1. Install `protoc` from the package manager of your system.

```sh
dnf install -y protobuf-compiler
```

2. Install `protoc-gen-go`.

```sh
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Generate the go files from `.proto` files.

```sh
make protoc
```

Run tests

```sh
make test
```

Build the project (binary : build/stack)

```sh
make build
```

Create container

```sh
make container
```

Start container

```sh
make run
```

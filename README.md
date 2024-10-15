## Requirements

https://grpc.io/docs/languages/go/quickstart/

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Quit Start

1. Clone the repo

```bash
$ git clone https://github.com/cheesecrust/goraft.git
```

2. Build

```bash
$ protoc --go_out=. --go-grpc_out=. proto/sample.proto
$ go build .
```

3. Start Node

```bash
$ ./exmaple [-port <host_port>] [-client <client_port1>,<client_port2>...]
```

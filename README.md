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

현재는 실행 파일이 올라가 있기 때문에 go 와 grpc 가 있다면 **Strart Node** 명령어 만으로 실행 시킬 수 있습니다.

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



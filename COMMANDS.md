# Learning gRPC

### Commands to achieve these results

```sh
$ brew install protobuf

$ go mod init github.com/gustavolbs/learning-grpc

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

$ export PATH="$PATH:$(go env GOPATH)/bin"

$ source ~/.bashrc

# para gerar as stubs
$ protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.
```

下载go依赖包

```shell
go get -u google.golang.org/grpc@v1.29.1
```

生成protp文件

```shell
protoc --go_out=plugins=grpc:. ./proto/*.proto
```

调试 gRPC 接口

grpcurl 是一个命令行工具，可让你与 gRPC 服务器进行交互，安装命令如下：

```shell
go get github.com/fullstorydev/grpcurl
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

测试

```shell
grpcurl -plaintext localhost:8001 list
grpcurl -plaintext localhost:8001 list proto.TagService
grpcurl -plaintext -d '{\"name\":\"Go\"}' localhost:8001 proto.TagService.GetTagList

```
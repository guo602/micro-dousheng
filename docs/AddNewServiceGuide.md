# How To Add Services - 如何添加一个服务

1. 确定要添加的服务的名称（如：`user`, `auth`, `feed` 等）可参考 `api_gateway/router.go` 下的 hertz group 或 API URL
2. 编写 thrift IDL 文件，并将文件命名为`{服务名称}.thrift` ，放入 idl 目录
3. 打开终端 cd 到项目根目录，然后调用 `./add-kitex-service.sh {服务名称}`
4. 此时 kitex 生成的代码将被放入 `kitex_gen` 目录（无需改动）和 `rpc/{服务名称}` 
5. 完成服务业务逻辑并妥善修改 `rpc/{服务名称}/handler.go`
6. 在`api_gateway/biz/handler` 中新建对应的handler `{服务名称}.go` 如user.go feed.go
7. 在 `api_gateway/router.go`  中添加对 RPC 服务的调用
8. 想要运行这个服务需要先查看根目录下的 Makefile 文件中有该服务的make代码。如没有，类比Makefile中`user`的写一下，然后在根目录下运行`make {服务名}`
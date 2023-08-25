# douyin micro.
## 开发
 参考[如何添加新的微服务模块](docs/AddNewServiceGuide.md)来添加新的微服务模块
 ## 运行
 ### 启动 HTTP 服务

```shell
make api
```

### 启动 RPC 服务

```shell
make user
make feed
make {your_service}
...
```

 ## 技术栈

| 功能      | 实现                  |
|---------|---------------------|
| HTTP 框架 | Hertz               |
| RPC 框架  | Kitex               |
| 数据库     | MySQL、Redis |
| 身份鉴权    | jwt              |
| 对象存储    | Aliyun-oss               |


## 目录介绍

| 目录     | 介绍             |
|--------|----------------|
| api_gateway    | hertz http 框架         |
| Idl    | 项目所有服务的 IDL 文件 |
| kitex_gen | kitex自动生成代码          |
| rpc | kitex业务逻辑代码        |

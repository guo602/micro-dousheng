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
| 服务发现  | etcd               |
| 数据库     | MySQL、Redis |
| 身份鉴权    | jwt              |
| 对象存储    | Aliyun-oss               |
|    服务治理 |	OpenTelemetry|
|链路追踪	|Jaeger             |



## 目录介绍

| 目录     | 介绍             |
|--------|----------------|
| api_gateway    | hertz http 框架         |
| Idl    | 项目所有服务的 IDL 文件 |
| kitex_gen | kitex自动生成代码          |
| rpc | kitex业务逻辑代码        |

## Jaeger

Visit `http://127.0.0.1:16686/` on browser

#### Snapshots

![jaeger-tracing](./docs/images/Jaeger_0.png)

![jaeger-architecture](./docs/images/Jaeger_1.png)

## Grafana

Visit `http://127.0.0.1:3000/` on browser

#### Dashboard Example

![grafana-dashboard-example](./docs/images/grafana_trace.png)

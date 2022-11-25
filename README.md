### 项目目录
```
├── README.md
├── app // 服务
│   ├── order // 订单微服务
│   │   ├── api // 订单微服务 api 方式调用（对外）
│   │   │   └── order.api
│   │   ├── model // 订单微服务 api和rpc 服务共用 model
│   │   │   └── ddl.sql
│   │   └── rpc // 订单微服务 rpc 方式调用（对内）
│   │       └── order.proto
│   └── ucenter // 用户中心微服务
│       ├── api // 用户中心微服务 api 方式调用（对外）
│       │   ├── Dockerfile
│       │   ├── etc // 配置
│       │   │   ├── prod // 生产
│       │   │   │   └── ucenter.yaml
│       │   │   ├── test // 测试
│       │   │   │   └── ucenter.yaml
│       │   │   └── ucenter.yaml // 本地
│       │   ├── internal
│       │   │   ├── config
│       │   │   │   └── config.go
│       │   │   ├── handler // api
│       │   │   │   ├── loginhandler.go
│       │   │   │   ├── routes.go
│       │   │   │   └── userinfohandler.go
│       │   │   ├── logic // 业务逻辑
│       │   │   │   ├── loginlogic.go
│       │   │   │   └── userinfologic.go
│       │   │   ├── svc // 服务注册
│       │   │   │   └── servicecontext.go
│       │   │   └── types // 结构（自动生成，禁止改动）
│       │   │       └── types.go
│       │   ├── ucenter.api // 所有 api 接口定义
│       │   └── ucenter.go // 入口文件
│       ├── model // 用户中心微服务 api和rpc 服务共用 model
│       │   ├── ddl.sql // sql ddl，改动后生成 model 文件
│       │   ├── tkusermodel.go // 工具生成 goctl model mysql ddl -src ./model/ddl.sql -dir ./model -c
│       │   ├── tkusermodel_gen.go
│       │   └── vars.go
│       └── rpc // 用户中心微服务 rpc 方式调用（对内）
│           ├── Dockerfile
│           ├── etc // 配置，服务注册和发现采用 etcd
│           │   ├── prod
│           │   │   └── ucenter.yaml
│           │   ├── test
│           │   │   └── ucenter.yaml
│           │   └── ucenter.yaml
│           ├── internal
│           │   ├── config
│           │   │   └── config.go
│           │   ├── jobs
│           │   │   └── jobs.go // job 获取 queue
│           │   ├── logic
│           │   │   ├── jobslogic.go
│           │   │   └── userinfologic.go
│           │   ├── server
│           │   │   └── ucenterserver.go
│           │   └── svc
│           │       └── servicecontext.go
│           ├── types
│           │   └── ucenter
│           │       ├── ucenter.pb.go
│           │       └── ucenter_grpc.pb.go
│           ├── ucenter
│           │   └── ucenter.go
│           ├── ucenter.go // 入口文件
│           └── ucenter.proto // proto 定义
├── build.sh
├── common // 公共包
│   ├── const.go
│   ├── cryptx
│   ├── errorx
│   │   └── baseerror.go // 自定义错误
│   ├── response
│   │   └── response.go // 自定义接口输出格式
│   └── utils
├── go.mod
└── go.sum
```

### 修改 `handle` 模板

如果本地没有 `~/.goctl/${goctl版本号}/api/handler.tpl` 文件，可以通过模板初始化命令 `goctl template init` 进行初始化

修改模板 `vim ~/.goctl/${goctl版本号}/api/handler.tpl`

```
package handler

import (
    "net/http"
    "tt90.cc/ucenter/common/response"
    {{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        {{if .HasRequest}}var req types.{{.RequestType}}
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }{{end}}

        l := logic.New{{.LogicType}}(r.Context(), svcCtx)
        {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        {{if .HasResp}}response.Response(w, resp, err){{else}}response.Response(w, nil, err){{end}}
            
    }
}
```

### 根据 `DDL` 生成 `MODEL`

1. 修改 ddl `cd ${服务}/rpc && vim ./model/ddl.sql`
2. 在项目根目录执行 `goctl model mysql ddl -src ./model/ddl.sql -dir ./model -c`


### 生成 `api` 或者 `rpc` 代码

1. 进入 `${服务}/rpc` 或者 `${服务}/api`
2. 生成 api：`goctl api go -api ./ucenter.api -dir .`
3. 生成 rpc：`goctl rpc protoc ./ucenter.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.`

### 本地docker安装 `etcd-serve`

```shell
docker run -d --name etcd-server \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    bitnami/etcd:latest
```

### 服务端口设置规范

| 服务          | 端口           |
|-------------|--------------|
| ucenter-rpc | 8213         |
| ucenter-api | 8214         |

### 测试 api

```shell
curl --location --request POST 'http://localhost:8214/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{"username":"ncty","password":"123456"}'
```

### 测试在 api 中调用 rpc 服务

```shell
curl 'http://localhost:8214/user/info' --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjkxOTk1MzgsImlhdCI6MTY2OTE5MjMzOCwidXNlcklkIjoxfQ.pK06HqrU4qu0mC7Txje4h09rsRuYH2PelxEJ6sDMhoo' \
--header 'Content-Type: application/json'
```
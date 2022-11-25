### 项目目录
```
├── README.md
├── api
│   ├── Dockerfile
│   ├── etc
│   │   ├── prod
│   │   │   └── ucenter.yaml
│   │   ├── test
│   │   │   └── ucenter.yaml
│   │   └── ucenter.yaml
│   ├── internal
│   │   ├── config
│   │   │   └── config.go
│   │   ├── handler
│   │   │   ├── loginhandler.go
│   │   │   ├── routes.go
│   │   │   └── userinfohandler.go
│   │   ├── logic
│   │   │   ├── loginlogic.go
│   │   │   └── userinfologic.go
│   │   ├── svc
│   │   │   └── servicecontext.go
│   │   └── types
│   │       └── types.go
│   ├── ucenter.api
│   └── ucenter.go
├── build.sh
├── common
│   ├── const.go
│   ├── cryptx
│   ├── errorx
│   │   └── baseerror.go
│   ├── response
│   │   └── response.go
│   └── utils
├── go.mod
├── go.sum
├── logs
├── model
│   ├── ddl.sql
│   ├── tkusermodel.go
│   ├── tkusermodel_gen.go
│   └── vars.go
└── rpc
    ├── Dockerfile
    ├── etc
    │   ├── prod
    │   │   └── ucenter.yaml
    │   ├── test
    │   │   └── ucenter.yaml
    │   └── ucenter.yaml
    ├── internal
    │   ├── config
    │   │   └── config.go
    │   ├── jobs
    │   │   └── jobs.go
    │   ├── logic
    │   │   ├── jobslogic.go
    │   │   └── userinfologic.go
    │   ├── server
    │   │   └── ucenterserver.go
    │   └── svc
    │       └── servicecontext.go
    ├── types
    │   └── ucenter
    │       ├── ucenter.pb.go
    │       └── ucenter_grpc.pb.go
    ├── ucenter
    │   └── ucenter.go
    ├── ucenter.go
    └── ucenter.proto
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

1. 修改 ddl `vim ./model/ddl.sql`
2. 在项目根目录执行 `goctl model mysql ddl -src ./model/ddl.sql -dir ./model -c`


### 生成 `api` 或者 `rpc` 代码

1. 进入 `./rpc` 或者 `./api`
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
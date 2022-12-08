### 项目目录
```
├── README.md
├── api
│   ├── Dockerfile
│   ├── etc
│   │   ├── prod
│   │   ├── test
│   │   └── ucenter.yaml
│   ├── internal
│   │   ├── config
│   │   ├── handler
│   │   ├── logic
│   │   ├── svc
│   │   └── types
│   ├── t_test.go
│   ├── ucenter.api
│   └── ucenter.go
├── build.sh
├── common
│   ├── cryptx
│   ├── ctxdata
│   │   └── ctxData.go
│   ├── errorx
│   │   ├── baseerror.go
│   │   └── errcode.go
│   ├── globalkey
│   │   ├── constantKey.go
│   │   └── redisCacheKey.go
│   ├── response
│   │   └── response.go
│   ├── tool
│   │   ├── coinconvert.go
│   │   ├── encryption.go
│   │   ├── krand.go
│   │   ├── krand_test.go
│   │   ├── placeholders.go
│   │   ├── set.go
│   │   └── time.go
│   └── uniqueid
│       ├── sn.go
│       ├── sn_test.go
│       └── uniqueid.go
├── doc
├── go.mod
├── go.sum
├── model
│   ├── ddl.sql
│   ├── tkusermodel.go
│   ├── tkusermodel_gen.go
│   └── vars.go
└── rpc
    ├── Dockerfile
    ├── etc
    │   ├── prod
    │   ├── test
    │   └── ucenter.yaml
    ├── internal
    │   ├── config
    │   ├── jobs
    │   ├── logic
    │   ├── server
    │   └── svc
    ├── t_test.go
    ├── types
    │   └── ucenter
    ├── ucenter
    │   └── ucenter.go
    ├── ucenter.go
    └── ucenter.proto
```

### 替换 tpl 模板

下载并替换 `goctl` 模板

```shell
git clone https://github.com/tt90cc/goctl-template.git && rm -rf ~/.goctl/$(goctl -v|awk '{print $3}')/* && cd goctl-template && mv ./* ~/.goctl/$(goctl -v|awk '{print $3}')
```

### 根据 `DDL` 生成 `MODEL`

1. 修改 ddl `cd ./model && vim ./ddl.sql`
2. 在项目根目录执行 `goctl model mysql ddl -src ./ddl.sql -dir . -c`

##### 复杂查询

```go
squirrel.Or{squirrel.Expr("id=?", cast.ToInt64(req.Name)), squirrel.And{squirrel.Eq{"name": req.Name}}}
// squirrel.Or{squirrel.Eq{"id": cast.ToInt64(req.Name)}, squirrel.And{squirrel.Eq{"name": req.Name}}}

Where("FIND_IN_SET(?, platform_type)", req.PlatformType)
```

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

### 运行服务

##### 1.编译
```
./build.sh rpc prod
```

##### 2.运行容器
```
docker run -d --name serve.ucenter-api -p 8214:8214 -v /tmp/logs:/app/logs serve.ucenter-api
```

### 生成 kube

```
cd api
goctl kube deploy --name ucenter-api --image serve.ucenter-api --namespace default --port 8214 -o kube-ucenter-api.yaml
```

### 常用包

* cast类型转换：`go get github.com/spf13/cast`
* crontab任务：`go get github.com/robfig/cron/v3`
* err输出：`go get github.com/pkg/errors`
* copier：`go get github.com/jinzhu/copier`
* id生成：`go get github.com/sony/sonyflake`
* validator参数验证：`go get github.com/go-playground/validator/v10`
* 微信公众号小程序开发：`go get github.com/silenceper/wechat/v2`
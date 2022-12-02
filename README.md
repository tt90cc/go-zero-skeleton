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

#### 自定义错误

如果本地没有 `~/.goctl/${goctl版本号}/api/handler.tpl` 文件，可以通过模板初始化命令 `goctl template init` 进行初始化

修改模板 `vim ~/.goctl/${goctl版本号}/api/handler.tpl`

```
package handler

import (
    "net/http"
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

### 修改 `model` 模板

##### interface-insert.tpl
```
Insert(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result,error)
```

##### interface-delete.tpl
```
Delete(ctx context.Context, session sqlx.Session, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error
```

##### interface-update.tpl
```
Update(ctx context.Context, session sqlx.Session, newData *{{.upperStartCamelObject}}) error
```

##### model.tpl
```
package {{.pkg}}
{{if .withCache}}
import (
  "context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
{{else}}
import (
  "context"
  "github.com/zeromicro/go-zero/core/stores/sqlx"
)
{{end}}
var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
	// {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Model.
	{{.upperStartCamelObject}}Model interface {
		{{.lowerStartCamelObject}}Model
    Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	custom{{.upperStartCamelObject}}Model struct {
		*default{{.upperStartCamelObject}}Model
	}
)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf{{end}}) {{.upperStartCamelObject}}Model {
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn{{if .withCache}}, c{{end}}),
	}
}

func (c *custom{{.upperStartCamelObject}}Model) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return c.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}
```

##### insert.tpl
```

func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	{{if .withCache}}{{.keys}}
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    if session != nil {
      return session.ExecCtx(ctx, query, {{.expressionValues}})
    }
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    s := m.conn
    if session != nil {
      s = session
    }
    ret,err:=s.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return ret,err
}

```

##### update.tpl
```

func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, session sqlx.Session, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    if session != nil {
      return session.ExecCtx(ctx, query, {{.expressionValues}})  
    }
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    s := m.conn
    if session != nil {
      s = session
    }
    _,err:=s.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return err
}

```

##### delete.tpl
```

func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, session sqlx.Session, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, err {{if .containsIndexCache}}={{else}}:={{end}} m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
    if session != nil {
      return session.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})  
    }
		return conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
    s := m.conn
    if session != nil {
      s = session
    }
		_,err:=s.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}}){{end}}
	return err
}

```

### 根据 `DDL` 生成 `MODEL`

1. 修改 ddl `cd ./model && vim ./ddl.sql`
2. 在项目根目录执行 `goctl model mysql ddl -src ./ddl.sql -dir . -c`


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
docker run -d --name serve.ucenter_rpc -p 8213:8213 -v /tmp/logs:/app/logs serve.ucenter_rpc
```

### 常用包

* cast类型转换：`go get github.com/spf13/cast`
* crontab任务：`go get github.com/robfig/cron/v3`
* err输出：`go get github.com/pkg/errors`
* copier：`go get github.com/jinzhu/copier`
* id生成：`go get github.com/sony/sonyflake`
* validator参数验证：`go get github.com/go-playground/validator/v10`
* 微信公众号小程序开发：`go get github.com/silenceper/wechat/v2`
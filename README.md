- [简介](#简介)
- [Projects Catalog](#projects-catalog)
	- [1. 算法与数据结构](#1-算法与数据结构)
		- [Algorithms and Data Structures Go](#algorithms-and-data-structures-go)
	- [2. Web 网络编程](#2-web-网络编程)
		- [Gin 脚手架](#gin-脚手架)
			- [ORM Interface 继承](#orm-interface-继承)
			- [Http Timeout](#http-timeout)
			- [分布式链路 Jaeger](#分布式链路-jaeger)
			- [编译](#编译)
			- [运行](#运行)
			- [总结](#总结)
	- [3. Cli 工具](#3-cli-工具)
		- [scaffold](#scaffold)
		- [cobra](#cobra)
- [Appendix](#appendix)
	- [Golang Open Source Projects](#golang-open-source-projects)
- [Maintainer](#maintainer)

# 简介
goAlltime 是一个学习 golang (新手入门级) 的项目。它包含了一些 数据结构与算法、Web、Cli、Lib 包等脚手架 …

# Projects Catalog
## 1. 算法与数据结构
### Algorithms and Data Structures Go
[algo data structures](https://github.com/codeandcode0x/algo-data-structures)


## 2. Web 网络编程

### Gin 脚手架
Go 语言有非常的优秀的特性 (比如高并发、原生支持协程、泛型等等), 同时也贡献了非常多项目(可以 https://awesome-go.com/ 一览)，在 Web 开发这块也有非常多优秀的框架，如 Gin、Beego、Iris、Echo、Revel 等. Top Go Web Frameworks

**为了快速使用 Gin, 我提供了 [Gin Scaffold](go-web/gin-scaffold/README.md) 程序。 提供如下功能:**

- ORM 封装 (使用的 GORM, 对 Model Interface 可进行继承设计, 可方便的封装 DAO 层)
- Tracing (支持链路追踪)
- Http TimeOut (支持 Http 请求超时中断)
- Swagger


**支持部署:**
- Docker / Docker Compose
- K8s Manifest
- Helm

#### ORM Interface 继承
设计 BaseModel, 在后面的业务 Model 继承这个 BaseModel 即可抽象 DAO interface
```go

// base model
type BaseModel struct {
	ID        uint64 `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// base dao
type BaseDAO interface {
	Create(entity interface{}) error
	Update(entity interface{}, uid uint64) (int64, error)
	Delete(entity interface{}, uid uint64) (int64, error)
	FindAll(entity interface{}, opts ...DAOOption) error
	FindByKeys(entity interface{}, keys map[string]interface{}) error
	FindByPages(entity interface{}, currentPage, pageSize int) error
	FindByPagesWithKeys(entity interface{}, keys map[string]interface{}, currentPage, pageSize int) error
	SearchByPagesWithKeys(entity interface{}, keys, keyOpts map[string]interface{}, currentPage, pageSize int) error
	Count(entity interface{}, count *int64) error
	CountWithKeys(entity interface{}, count *int64, keys, keyOpts map[string]interface{}) error
}
```
如 UserDAO 继承 BaseDAO, 然后编写 UserDAO 独有的业务 dao 方法
```go

// instance entity
type User struct {
	ID           uint64         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty" gorm:"type:varchar(255)"`
	Password     string         `json:"password,omitempty" gorm:"type:varchar(1000)"`
	Email        string         `json:"email,omitempty" gorm:"type:varchar(255)"`
	Age          int            `json:"age,omitempty"`
	Birthday     time.Time      `json:"birthday,omitempty"`
	MemberNumber sql.NullString `json:"member_number,omitempty" gorm:"type:varchar(255)"`
	Role         string         `json:"Role,omitempty" gorm:"type:varchar(100)"`
	BaseModel
}

// user DAO
type UserDAO interface {
	BaseDAO
	FindAllByCount(count int) ([]User, error)
}
```

#### Http Timeout
设计 TimeoutHandler middleware 来处理超时请求
```go
const (
	TIME_DURATION = 10
)

func DefinitionRoute(router *gin.Engine) {
	// set run mode
	gin.SetMode(gin.DebugMode)
	// middleware
	router.Use(middleware.Tracing())
	router.Use(middleware.UseCookieSession())
	router.Use(middleware.TimeoutHandler(time.Second * TIME_DURATION))
	// no route
	router.NoRoute(NoRouteResponse)
	// home
	var userController *controller.UserController
	router.Static("/web/assets", "./web/assets")
	router.StaticFS("/web/upload", http.Dir("/web/upload"))
	router.LoadHTMLGlob("web/*.tmpl")
...

```

#### 分布式链路 Jaeger
引入 traceandtrace-go 实现 Tracing middleware, 上报链路信息到 jaeger
```go
import (
	"log"

	tracing "github.com/codeandcode0x/traceandtrace-go"
	"github.com/gin-gonic/gin"
)

// Tracing 中间件
func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("..... tracing header1 ", c.Request.Header)
		// add tracing
		pctx, cancel := tracing.AddHttpTracing(
			"ticket-manager",
			c.Request.URL.String()+" "+c.Request.Method,
			c.Request.Header,
			map[string]string{
				"component":      "gin-server",
				"httpMethod":     c.Request.Method,
				"httpUrl":        c.Request.URL.String(),
				"proto":          c.Request.Proto,
				"peerService":    c.HandlerName(),
				"httpStatusCode": "200",
				"spanKind":       "server",
			})
		defer cancel()

		// deliver parent context
		c.Set("parentCtx", pctx)
		c.Next()
		return
	}
}

```

#### 编译

```sh
cd /web/gin && ./docker/build.sh 
```
build.sh $1 $2 $3 可以传递三个参数
- $1: appVersion (版本号 如 1.0.0 、latest)
- $2: repoUser ( docker repo username)
- $3: repoHost ( docker repo host)

#### 运行
- docker compose
启动
```sh
 docker-compose --env-file deploy/config/.env -f docker/docker-compose.yml -p gin-scaffold  up
```
卸载
```sh
 docker-compose --env-file deploy/config/.env -f docker/docker-compose.yml -p gin-scaffold  down
```

- k8s manifest
```sh
kubectl apply -f k8s/manifest --recursive / -R
```
- helm
```sh
helm dep update k8s/helm/apps/web-app
helm upgrade --install gin-scaffold  k8s/helm/apps/web-app
```

#### 总结
- gin 在 go web 上使用非常方便, 并且性能非常不错 
- 使用 gin 脚手架可以快速构建开发项目


## 3. Cli 工具
### scaffold
[cli scaffold](cli/cli-tools/cli-scaffold/README.md)
### cobra
[cobra](github.com/spf13/cobra)

# Appendix
## Golang Open Source Projects
[golang open source projects](https://github.com/codeandcode0x/golang-open-source-projects) 

# Maintainer
- roancsu@163.com
- codeandcode0x@gmail.com
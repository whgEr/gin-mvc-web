## 包管理工具
`go get github.com/kardianos/govendor`

## 框架
`govendor fetch github.com/gin-gonic/gin@v1.4`

## 热启动工具
`go get github.com/pilu/fresh`  

`$ fresh`

## 项目结构

```

.

├── README.md
├── asset                       静态资源文件
│   ├── css
│   ├── js
│   └── image
├── command                     命令行
│   └── command.go
├── config                      全局配置
│   ├── mysql.go
│   ├── redis.go
│   └── env.go
├── frontend                    前端项目
│   ├── controller
│   ├── middleware
│   └── view
├── backend                     后端项目
│   ├── controller
│   ├── middleware
│   └── view
├── api                         移动端项目
│   ├── controller
│   └── middleware
├── middleware                  全局中间件
├── model                       全局model
├── service                     业务层
├── vendor                      govendor包
├── storage                     
│   ├── cache                   缓存文件
│   └── log                    项目日志
│       ├── info.log          
│       └── error.log
├── main.go                     项目入口
└── router.go                   路由
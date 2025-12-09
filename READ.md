# go实验室

## 原则

* 要学就好好学，别三天打鱼两天晒网的。

* 要想想别人是怎么做的，少自己自己为是

* 先框架，在细节

## 目录设计

```(plaintext)
golab/
├── go.mod                  // 必选：Go模块依赖（项目核心标识）
├── go.sum                  // 自动生成：依赖版本锁
├── .gitignore              // 基础版：过滤无用文件
├── README.md               // 简单说明：启动方式、核心功能
├── cmd/                    // 程序入口：仅保留Web服务入口
│   └── web/
│       └── main.go         // 项目启动文件（唯一入口）
├── config/                 // 配置管理：简化版（单配置文件）
│   ├── config.go           // 配置结构体 + 加载方法
│   └── app.yaml            // 配置文件（数据库、端口等，初期共用）
├── internal/               // 核心业务：私有目录（外部不可导入）
│   ├── handler/            // 接口处理层（控制器）
│   │   └── handler.go      // 初期所有接口都放这，后续拆模块
│   ├── service/            // 业务逻辑层：核心逻辑都在这
│   │   └── service.go      // 初期所有业务逻辑放这
│   └── model/              // 数据模型：实体/DTO通用
│       └── model.go        // 初期所有模型（用户、订单等）放这
└── pkg/                    // 通用工具：仅保留核心
    ├── db/                 // 数据库连接封装
    │   └── db.go
    └── utils/              // 通用工具函数（字符串、时间等）
        └── utils.go
```

## 项目初始化步骤

```(bash)
# 创建目录
mkdir golab && cd golab 
# 初始化项目
go mod init github.com/lnanmu/golab
# 根据设计的目录结构创建文件
mkdir -p ./{cmd/web,config,internal/{handler,service,model},pkg/{db,utils}}
# 创建核心文件
touch config/{config.go,app.yaml} cmd/web/main.go internal/{handler/handler.go,service/service.go,model/model.go} pkg/{db/db.go,utils/utils.go}
# 生成.gitignore
curl -o .gitignore https://raw.githubusercontent.com/github/gitignore/main/Go.gitignore
```

## 推送代码

```(plaintext)
在推送之前，先同步github上的代码，点击左下脚的双箭头图标。
然后进行暂存和推送
推送时候可以总左下角看到需要推送和拉取的代码。
```
```
.
├── LICENSE
├── README.md
├── configs         # 配置文件
├── docs            # 文档集合
├── global          # 全局变量
├── go.mod
├── go.sum
├── internal        # 内部模块
│         ├── dao             # 数据访问层（Database Access Object），所有与数据相关的操作都会在dao层进行，例如 MySQL、Elasticsearch等; 连接业务层和 model 层
│         ├── middleware      # HTTP中间件: 国际化(validate 报错转中文)
│         ├── model           # 模型层，用于存放model对象
│         ├── routers         # 路由相关的逻辑
│         └── service         # 项目核心业务逻辑: model参数校验
├── main.go
├── storage         # 项目生成的临时文件
├── scripts         # 各类构建、安装、分析等操作的脚本
├── third_party     # 第三方的资源工具，如Swagger UI
└── pkg             # 项目相关的模块包
    ├── app         # 分页处理、响应处理
    ├── convert     # string int 类型转换
    ├── email       # 发送报警 email
    ├── errcode     # 通用错误码定义
    ├── limiter
    ├── logger      # 日志封装 堆栈打印
    ├── setting     # viper 文件配置
    ├── tracer
    ├── upload      # 上传
    ├── util        # tools
    └── validator

```

```
model 的调用：router(api->handle) -> service -> dao -> model
```
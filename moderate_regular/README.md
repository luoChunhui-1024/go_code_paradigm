## 中规中矩普通版

```txt
├── README.md
├── biz
│   ├── config // 配置文件各个模块对应的结构体和从配置文件中读取配置到内存对象
│   ├── dal    // 数据访问层，放置数据库每个表的查询和修改代码
│   ├── handler // http 接口参数解析
│   ├── service // 封装具体业务逻辑
│   └── model // 业务model代码
├── build.sh
├── cmd     // 自动生成表的 model 和 query 的工具
├── common  // 定义了一些常量和常见错误状态码
│   ├── constdef // 常量
│   └── response // 常见错误状态码
├── conf    // 配置文件
├── entity  // 表实体
├── kitex_gen // 自动生成的 rpc 服务的 thrift 客户端代码
├── main.go   // 主函数入口
├── output    // 编译产物
├── pkg       // 各种第三方服务接口和工具的初始化和使用方法都统一放在 pkg 路径下
│   ├── auth    // 授权相关的接口工具
│   ├── bot     // 飞书机器人相关的接口
│   ├── db      // 管理数据库初始化以及所有 store 对象的创建
│   ├── metrics // 日志打点工具
│   ├── middlewares// 存放中间件
│   ├── rpc     // 第三方 rpc 服务的初始化以及相关接口的函数调用
│   └── utils   // 通用工具，比如类型转换、生成logid、json转换、发送 http 消息
├── router.go    // 各类路由
├── router_gen.go
└── script
```
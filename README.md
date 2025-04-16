project/
├── routes/               # 定义路由和接口
├── cmd/               # 启动入口
│   └── main.go
├── config/            # 配置文件
│   └── config.yaml
├── internal/          # 内部核心代码
│   ├── controllers/   # 控制器，负责处理请求
│   ├── services/      # 业务逻辑
│   ├── models/        # 数据模型层
│   ├── repositories/  # 数据访问层
│   └── middlewares/   # 中间件
├── pkg/               # 独立封装的公共库
│   └── logger/
├── docs/              # 项目文档
├── scripts/           # 脚本文件（如部署、数据迁移）
├── tests/             # 测试用例
└── go.mod             # 模块管理文件
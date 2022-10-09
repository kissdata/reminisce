# 叙旧服务模块

helloworld 典藏版

程序默认运行在 Ubuntu 2204上

## 技术栈

1. HTTP 服务(Gin)
2. ORM CUR 操作(bun)
3. PG14 操作

## 特色

- 成果显示, 能显示当月/年交流的人数。

- 使用 `归并排序` 找出急需交流的朋友

更多见

[![Documentation](https://img.shields.io/badge/doc-叙旧服务详设文档-information)](doc/design.md)



## 代码文件结构

```
├── cmd
│   ├── pgclient.go        # orm 操作 
│   └── main.go            # 启动入口
│   ├── statement.go       # 全局变量
│   └── test.go            # 测试
├── common
│   |── config.go          # 显示的信息结构
│   └── table              # 数据表结构体
│       ├── data.go
│       └── friend.go           
├── doc
│   ├── design.md          # 详设文档
│   ├── table4data.sql     # 建表语句示例
│   └── table4friend.sql
├── docker
│   └── Dockerfile         # docker构建 (*)
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## 其他

使用 VS Code 调试

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "golang",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "env": {},
            "program": "${workspaceFolder}/cmd",
            "cwd": "${workspaceFolder}"
        }
    ]
}
```


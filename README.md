# 叙旧服务模块

helloworld 典藏版

程序默认运行在 Ubuntu 2204上, 需要先安装 Go 1.19+

## 技术栈

1. HTTP 服务(Gin)
2. ORM CUR 操作(bun)
3. PG14 操作

## 特色

- 实现两个http服务, 一个在 `:2210` 端口，另一个用Gin写的在 `:2222` 端口;

- 成果显示, 能显示当月/年交流的人数。

- 使用 `归并排序` 找出急需交流的朋友

更多见

[![Documentation](https://img.shields.io/badge/doc-叙旧服务详设文档-information)](doc/design.md)



## 代码文件结构

```
├── cmd
│   ├── html.go            # 网页操作 
│   ├── pgclient.go        # orm 操作 
│   ├── main.go            # 启动入口
│   ├── statement.go       # 全局变量
│   ├── test.go            # 测试
│   └── usegin.go          # http服务(Gin)
├── common
│   |── config.go          # 显示的信息结构
│   │── table              # 数据表结构体
│   │   ├── data.go
│   │   └── friend.go
│   └── templates          # 静态文件
│       └── index.tmpl      
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


## 调试方法

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

e.g.1 浏览器打开 `localhost:2210` 查看。

e.g.2 浏览器打开 `localhost:2210/html/` 查看网页版。

e.g.3 Postman 设置需求URL `ip:2222/addchat` , 选择 **POST**, 菜单选择 Body -> x-www-form-urlencoded, 填写 key-value 参数, 点击 Send, 看响应情况。


## Reference 

- [7天用Go从零实现Web框架Gee教程](https://geektutu.com/post/gee-day3.html)

- [Go Web Examples Courses](https://gowebexamples.com/hello-world/)
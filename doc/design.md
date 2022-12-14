// 详细设计文档


## 一、需求和功能说明

1. 输入名字，查到上次聊天时间

2. 输入身份，查到同类身份已经聊过的人

3. 增加记录，比较日期，确实是不是（补记）

4. 设计模板页面，浏览器显示信息页面

### more plan

- 增加 friend_type 字段, 可选 大女神、女神、闺蜜、校友、师傅、死党、普通、初识。

- 向另一个服务模块（syncselfie 显示好友照片）发送请求, 同步服务器目录里最新的自拍照, 该服务模块每天更新一次


## 二、IPO

| 输入  | 处理    | 输出      |
| ----- | -------- | ------- |
| 人名  | 读取 PG | 结构体输出 name-昵称-身份-上次聊天时间      |
| 月份  | 读取 PG | 在该月自己聊过的所有人的信息  |
| 新增 | 识别    | 增加完成\|失败         |
| 身份 | 读取 PG | 同类身份所有人 name, 认识时间，上次聊天时间 |
| need  | 计算  | 接下来一周需要和哪些人聊天 |


## 三、接口设计

### 3.1 服务接口

主要能添加这一次交谈（顺带能够修改）和显示需要最近去沟通的人。

| 功能场景     | 调用方式 | 接口路径       |
| ------- | -------- | -------------- |
| 增加交流的记录 | Post| [service name]/add_talk  |
| 修改交流的记录 | Post | [service name]/update_talk |
| 提醒下周要聊的对象 | Get | [service_name]/next_talk |

无法新增和删除好友, 需要手动打开PG操作


### 3.2 输入接口

输入数据的结构
```json
{
    "pet_name": "妍大大",
    "data": "2022-09-30", // 本次聊天记录
    "content": "今天是在先维的最后一天"， // 聊天内容
    "remark": "补记",
    "creator": ""
}
```


### 3.3 响应接口

```json
type ResultData struct {
    Code    int `json:"code"`         // 状态码
    Msg     string `json:"msg"`       // 状态信息
    Display *Display `json:"data,omitempty"` // 数据
}
```

### 输入输出示例

【输入】 yz

【输出】

```json
{
	"code": 0, // ok
	"msg": "ok",
	"data": {
		"龚小琳", 
        "椰子",        // 昵称
        "闺蜜",        // 当前身份
        "2022-10-01", // 上次聊天日期
        "节日祝贺",    // 聊天话题
        "国庆快乐",    // 聊天内容
        "7",          // 距离上次聊天间隔, 单位天数
	}
}
```




## 附录 技术规格偏移表

```sql
id    技术指标                响应     偏移情况
===+========================+======+============
 1  提供好友昵称                Ok      无偏移
 2  支持增加记录                OK      无偏移
 3  支持显示当月记录            OK      无偏移
 4  支持显示下周需要聊天的对象   OK      无偏移 
 5  提供好友交流意向度          No      暂不支持
 6  提供好友情感经历            No      暂不支持
 7  程序能运行在Linux上         OK      无偏移
 8  提供输出接口                OK      无偏移
 9  提供PG数据库配置            OK      无偏移
10  使用时长不低于3年           OK      无偏移
11  维护更新不低于9年           OK      无偏移
```

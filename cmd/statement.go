package main

import "database/sql"

// 新增
type ChatData struct {
	PetName   string
	ChatDate  string // 交谈日期
	ChatTopic string // 话题
	Content   string // 内容
	Score     int    // 评价
}

var (
	DB2mysql *sql.DB // 连接 MySQL8
	DB2pg    *sql.DB // 连接 PostgreSQL14
)

var timeTemplates = []string{
	"2006-01-02 15:04:05", //常规类型
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
}

// 好友身份
var FriendTypeMap = map[string]int{
	"大女神": 0,
	"女神":  0,
	"闺蜜":  0,
	"校友":  0,
	"师傅":  0,
	"死党":  0,
	"普通":  0,
	"初始":  0,
}

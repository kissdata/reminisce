package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 功能: 查询当月叙旧记录
//
//	@param URI pet_name  通过好友昵称
func getChatInfo(c *gin.Context) {
	petName := c.Param("pet_name")
	if ok := CheckPetName(petName); !ok {
		c.JSON(200, gin.H{
			"pet_name": petName,
			"⚠ notice": "not this friend, pet name ok?",
		})
		return
	}

	year := c.DefaultQuery("year", time.Now().Format("2006"))
	month := c.Query("month")
	if month == "" {
		month = time.Now().Format("01")
	}

	chatLatest, exist := CheckChatLatest(petName)
	if !exist {
		c.JSON(200, gin.H{
			"time":     year + "-" + month,
			"⚠ notice": "not talk with " + petName + " more than 1 month...",
		})
		return
	}

	var output = ChatData{
		PetName:   petName,
		ChatDate:  chatLatest.Datetime.Format("2006-01-02"),
		ChatTopic: chatLatest.ChatTopic,
		Content:   chatLatest.Content,
	}
	c.JSON(200, gin.H{
		"date":        year + "-" + month,
		"chat_latest": output,
	})
}

// 功能: 获取同类身份的所有人的信息
//
//	@description: 包括姓名、初识时间、上次聊天时间、聊天话题、超过天数
func getInfoByFriendType(c *gin.Context) {
	friendType := c.Param("friend_type")
	if _, ok := FriendTypeMap[friendType]; !ok {
		c.JSON(200, gin.H{
			"friend_type": friendType,
			"⚠ notice":    "not this friend type, input ok?",
		})
		return
	}

	//TODO:
}

// 功能: 添加一条叙旧记录
//
//	@param pet_name  必须提供昵称
func addchat(c *gin.Context) {
	var chatdate ChatData

	chatdate.PetName = c.PostForm("pet_name")
	chatdate.ChatDate = c.PostForm("datetime")
	chatdate.ChatTopic = c.DefaultPostForm("chat_topic", "例行叙旧") // 未设置该参数
	chatdate.Content = c.PostForm("content")

	if chatdate.PetName == "" {
		c.JSON(200, gin.H{
			"state": "fail to add a new record! you need bring in a pet name (notnull) ",
			"date":  chatdate,
		})
		return
	}
	if chatdate.ChatTopic == "" {
		chatdate.ChatTopic = "例行叙旧" // 参数为空时
	}

	if ok := AddChatLatest(chatdate); !ok {
		c.JSON(200, gin.H{
			"state": "fail to add a new record!",
			"date":  chatdate,
		})
		return
	}

	result, _ := CheckChatLatest(chatdate.PetName)

	c.JSON(200, gin.H{
		"message": result,
	})

}

// 功能: Gin开发http服务, 使用端口2222
func GinMain() {
	router := gin.Default() // 携带基础中间件启动

	router.GET("/friend/:pet_name", getChatInfo)
	router.GET("/reminisce/:friend_type", getInfoByFriendType)
	router.POST("/addchat", addchat)

	router.Run(":2222")
}

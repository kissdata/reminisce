// 存放所有测试案例, 确保更新后支持以往的测试案例
package main

import (
	"log"
)

// 功能: check ORM section is ok
func test221009() {
	friendInfo, _ := QueryFriendByName("潘瑜妍")
	contactArr, total := QueryContact(friendInfo.FriendId.String())
	log.Printf("have %d contact ways:%+v\n", total, contactArr)

	newTalk := ChatData{
		"椰子",
		"2022-10-08",
		"QQtag",
		"`晚上看到她10月改了签名是“待会去码头整点薯条” , 搜了这个梗的由来，问她，她说就是想当个海鸥`",
		1}
	if ok := AddChatLatest(newTalk); !ok {
		log.Println("fail to add a new talk record!")
	}
	log.Println("succeed in adding a new talk record!")
}

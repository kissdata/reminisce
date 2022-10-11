package main

import (
	"fmt"
	"log"
	"net/http"
	"reminisce/common"
	"time"
)

func main() {

	if err := connPG14(); err != nil {
		log.Println("fail to connect to pg14, stop running! err=", err)
		return
	}

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":2210", nil); err != nil {
		log.Println(err)
	}

}

// 显示当月叙旧情况
func handler(w http.ResponseWriter, r *http.Request) {
	curYear, curMonth := time.Now().Format("2006"), time.Now().Format("01")
	fmt.Fprintf(w, "叙旧情况:%s年%s月\n", curYear, curMonth)

	chatList, count := QueryChatLastest(curYear, curMonth)
	if count == 0 {
		fmt.Println("not result")
		return
	}

	var outputArr []common.OutputDefault
	var tmpRecord common.OutputDefault
	for _, chatInfo := range chatList {
		// 先由pet_name获取friend_id
		friendInfo, err := QueryFriend(chatInfo.FriendId)
		if err != nil {
			fmt.Println("该好友已被删除, friend_id=", chatInfo.PetName)
			continue
		}
		tmpRecord.FirstKnowTime = friendInfo.FirstKnow
		tmpRecord.AcquaintedTime = friendInfo.FirstKnow
		tmpRecord.PetName = chatInfo.PetName
		tmpRecord.LatestTime = chatInfo.Datetime.Format("2006-01-02")
		tmpRecord.TalkTopic = chatInfo.ChatTopic
		tmpRecord.Content = chatInfo.Content

		outputArr = append(outputArr, tmpRecord)
	}

	for _, record := range outputArr {
		fmt.Fprintf(w, "%q\n", record)
	}
	fmt.Fprintf(w, "本月共和%d个好友叙旧。", count)
}

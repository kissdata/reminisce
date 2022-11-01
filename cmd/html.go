package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func handlerHtml(w http.ResponseWriter, r *http.Request) {
	// 通过执行程序找到cmd目录路径
	e, err := os.Executable()
	if err != nil {
		return
	}
	cmdDir := filepath.Dir(e)
	projDir := strings.TrimSuffix(cmdDir, "/cmd")

	tmplPath := filepath.Join(projDir, "/common/templates/index.tmpl")

	var iterate = Iterate
	t := template.New("index.tmpl").Funcs(template.FuncMap{
		"Iterate": iterate,
	})

	t, err = t.ParseFiles(tmplPath)
	if err != nil {
		log.Panicln("解析html模板失败, err: ", err)
		return
	}

	t.Execute(w, nil)
	q := r.URL.Query()
	curYear, curMonth := q.Get("year"), q.Get("month")
	chatList, count := QueryChatLastest(curYear, curMonth)
	if count == 0 {
		fmt.Println("not result")
		return
	}

	var displayData = make([]map[string]string, 0)
	for _, chatInfo := range chatList {
		// 先由pet_name获取friend_id
		friendInfo, err := QueryFriend(chatInfo.FriendId)
		if err != nil {
			fmt.Println("该好友已被删除, friend_id=", chatInfo.PetName)
			continue
		}
		// 必须在for内make, 否则append后输出数据都一样
		tmpRecord := make(map[string]string, 6)
		tmpRecord["firstknow"] = friendInfo.FirstKnow
		tmpRecord["petname"] = chatInfo.PetName
		tmpRecord["latesttime"] = chatInfo.Datetime.Format("2006-01-02")
		tmpRecord["talktopic"] = chatInfo.ChatTopic
		tmpRecord["content"] = chatInfo.Content

		displayData = append(displayData, tmpRecord)
	}

	t.Execute(w, displayData) // 数据替换, 只需要执行一次

}

// 模板中使用
func Iterate(count int) []int {
	var Items []int
	for i := 0; i < count; i++ {
		Items = append(Items, i)
	}
	return Items
}

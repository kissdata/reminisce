// 映射pg数据类型的表
package table

import (
	"time"

	"github.com/uptrace/bun"
)

// 最新交流记录表
type ChatLatest struct {
	bun.BaseModel `bun:"table:chat_latest"`

	Ts        time.Time `json:"ts"         bun:"ts,default:current_timestamp"`
	FriendId  string    `json:"friend_id"  bun:"friend_id,pk"`
	PetName   string    `json:"pet_name"   bun:"pet_name"`
	Datetime  time.Time `json:"datetime"   bun:"datetime,type:date"`
	ChatTopic string    `json:"chat_topic" bun:"chat_topic"`
	Content   string    `json:"content"    bun:"content"`
	Remark    string    `json:"remark"     bun:"remark"`
	Creator   string    `json:"creator"    bun:"creator"`
}

// 叙旧历史记录
type ChatHistory struct {
	bun.BaseModel `bun:"table:chat_history"`

	Ts        time.Time `json:"ts"         bun:"ts,default:current_timestamp"`
	ID        int64     `json:"id"         bun:",pk,autoincrement"`
	FriendId  string    `json:"friend_id"  bun:"friend_id,notnull"`
	PetName   string    `json:"pet_name"   bun:"pet_name"`
	Datetime  time.Time `json:"datetime"   bun:"datetime,type:date"`
	ChatTopic string    `json:"chat_topic" bun:"chat_topic"`
	Content   string    `json:"content"    bun:"content"`
	Remark    string    `json:"remark"     bun:"remark"`
	Creator   string    `json:"creator"    bun:"creator"`
}

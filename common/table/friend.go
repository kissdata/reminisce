// 映射pg好友信息类型的表
package table

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// 好友信息表
type Friend struct {
	bun.BaseModel `bun:"table:friend"`

	Ts        time.Time `json:"ts"         bun:"ts,default:current_timestamp"`
	FriendId  uuid.UUID `json:"friend_id"  bun:"friend_id,pk,type:uuid,default:uuid_generate_v4()"`
	Name      string    `json:"name"       bun:"name,notnull"`
	FirstKnow string    `json:"first_know" bun:"first_know"` // 初识时间
	Remark    string    `json:"remark"     bun:"remark"`
	Creator   string    `json:"creator"    bun:"creator"`
	Status    string    `json:"status"` // field ignored
}

// 好友称呼表
type Greet struct {
	bun.BaseModel `bun:"greet"`

	Ts         time.Time `json:"ts"          bun:"ts,default:current_timestamp"`
	Id         string    `json:"id"          bun:"id,notnull"`
	Name       string    `json:"name"        bun:"name"`
	TailLetter string    `json:"tail_letter" bun:"tail_letter"`
	PetName    string    `json:"pet_name"    bun:"pet_name"`
	Remark     string    `json:"remark"      bun:"remark"`
	Creator    string    `json:"creator"     bun:"creator"`
}

// 好友联系方式表
type Contact struct {
	bun.BaseModel `bun:"contact"`

	Ts         time.Time `json:"ts"         bun:"ts,default:current_timestamp"`
	ID         int64     `json:"id"         bun:"id,pk,autoincrement"`
	FriendId   uuid.UUID `json:"friend_id"  bun:"friend_id,type:uuid"`
	Name       string    `json:"name"       bun:"name"`
	CallType   string    `json:"call_type"  bun:"call_type"`
	Number     string    `json:"number"     bun:"number"`
	GetTime    string    `json:"get_time"   bun:"get_time"`
	Preference bool      `json:"preference" bun:"preference"`
	Remark     string    `json:"remark"     bun:"remark"`
	Creator    string    `json:"creator"    bun:"creator"`
}

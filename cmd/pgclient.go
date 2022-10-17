package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"reminisce/common"
	"reminisce/common/table"
	"strconv"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// 功能: 时间格式字符串转换
func String2Time(tm string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		if nil == err && !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}

// 功能: 合并交流的内容，通过“§§”分隔
func contentCombine(contents ...string) string {
	var res bytes.Buffer
	for _, content := range contents {
		res.WriteString(content)
		res.WriteString("§§")
	}
	return strings.TrimSuffix(res.String(), "§§")
}

// 功能: 检查拼接的文本中的老内容有没有新内容
func hasDuplicate(combineStrs, x string) bool {
	contentArr, exist := strings.Split(combineStrs, "§§"), false
	for _, content := range contentArr {
		if content == x {
			exist = true
			break
		}
	}
	return exist
}

// 功能: 连接mysql8
func connMySQL() (err error) {

	return
}

// 功能: 连接pg14
func connPG14() (err error) {

	ip, port := common.PG14info.IP, strconv.Itoa(common.PG14info.Port)
	dbName := common.PG14info.DbName

	// 连接串格式 postgres://postgres:@localhost:5432/test?sslmode=disable
	dsn := "postgres://postgres:@" + ip + ":" + port + "/" + dbName + "?sslmode=disable"
	DB2pg = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return DB2pg.Ping()
}

// 功能: 查询好友
//
//	@param field[string]  姓名或者id
func QueryFriend(field string) (friendInfo *table.Friend, err error) {
	ctx := context.Background()
	ormDb := bun.NewDB(DB2pg, pgdialect.New())

	friendInfo = new(table.Friend)
	if len(field) < 26 { // 将输入字段当姓名查
		err = ormDb.NewSelect().Model(friendInfo).Column("friend.*").
			Where("name = ?", field).
			Scan(ctx)
		if err != nil {
			log.Panicln("fail to find frind info, err=", err)
			friendInfo = nil
		}
	} else {
		err = ormDb.NewSelect().Model(friendInfo).Column("friend.*").
			Where("friend_id = ?", field).
			Scan(ctx)
		if err != nil {
			log.Panicln("fail to find frind info, err=", err)
			friendInfo = nil
		}
	}
	return friendInfo, err
}

// 功能: 查询联系方式
func QueryContact(friendID string) ([]table.Contact, int) {
	ctx := context.Background()
	ormDb := bun.NewDB(DB2pg, pgdialect.New())

	//contacts := make([]table.Contact, 0)
	var contacts []table.Contact
	count, err := ormDb.NewSelect().Model(&contacts).
		Where("friend_id = ?", friendID).
		ScanAndCount(ctx)
	if err != nil {
		log.Println("fail to find contact info, err=", err)
		contacts = nil
	}
	return contacts, count
}

// 功能: 查询个人招呼信息
func QueryGreet(petName string) (*table.Greet, error) {
	ctx := context.Background()
	ormDb := bun.NewDB(DB2pg, pgdialect.New())

	greetRecord := new(table.Greet)
	err := ormDb.NewSelect().Model(greetRecord).Column("greet.*").
		Where("pet_name = ?", petName).
		Scan(ctx)
	if err != nil {
		log.Println("fail to find pet name, err=", err)
		greetRecord = nil
	}
	return greetRecord, err
}

// 功能: 确认私人昵称有对应的好友
func CheckPetName(petName string) bool {
	ctx := context.Background()
	ormDb := bun.NewDB(DB2pg, pgdialect.New())

	nameLike := "'%" + petName + "%'"
	exist, err := ormDb.NewSelect().Model((*table.Greet)(nil)).
		Where("pet_name LIKE " + nameLike).
		Exists(ctx)
	if err != nil {
		log.Println(err)
		exist = false
	}

	return exist
}

// 功能: 查询最新聊天记录
func CheckChatLatest(petName string) (*table.ChatLatest, bool) {
	ctx := context.Background()
	ormDb := bun.NewDB(DB2pg, pgdialect.New())

	chatRecord := new(table.ChatLatest)
	err := ormDb.NewSelect().Model(chatRecord).Column("chat_latest.*").
		Where("pet_name = ?", petName).
		Scan(ctx)
	if err != nil {
		log.Println(err)
		return nil, false
	}

	return chatRecord, true
}

// 功能: 查询最新聊天记录
//
//	@ return [int] 按年月查询的记录总数
//
// e.g. SELECT * FROM "chat_latest" WHERE datetime::VARCHAR LIKE '2022-10%';
func QueryChatLastest(year, month string) ([]table.ChatLatest, int) {
	ctx := context.Background()
	ormDb := bun.NewDB(DB2pg, pgdialect.New())

	dateLike := "'" + year + "-" + month + "%'"
	var latestChatList []table.ChatLatest
	err := ormDb.NewRaw("SELECT * FROM chat_latest WHERE datetime::VARCHAR LIKE "+dateLike).
		Scan(ctx, &latestChatList)
	if err != nil {
		log.Println("fail to find latest chat table, err=", err)
		latestChatList = nil
	}
	return latestChatList, len(latestChatList)
}

// 功能: 查询历史聊天记录
func QueryChatHistory(friendID string) ([]table.ChatHistory, int) {
	ctx := context.Background()
	ormDb := bun.NewDB(DB2pg, pgdialect.New())

	var oldChatList []table.ChatHistory

	count, err := ormDb.NewSelect().Model(&oldChatList).
		Where("friend_id = ?", friendID).
		ScanAndCount(ctx)
	if err != nil {
		log.Println("fail to find contact info, err=", err)
		oldChatList = nil
	}
	return oldChatList, count
}

// 功能: 增加一条数据
func InsertChatLatest(dataI any) (ok bool) {
	ctx := context.Background()
	ormDb := bun.NewDB(DB2pg, pgdialect.New())

	ok = true
	if data, exist := dataI.(ChatData); exist {
		greetp, err := QueryGreet(data.PetName)
		if err != nil {
			panic(err) // accidental error
		}
		friendInfo, err := QueryFriend((*greetp).Name)
		if err != nil {
			panic(err) // accidental error
		}

		latestRecord := table.ChatLatest{
			Ts:        time.Now(),
			FriendId:  friendInfo.FriendId.String(),
			PetName:   data.PetName,
			ChatTopic: data.ChatTopic,
			Content:   data.Content,
		}
		latestRecord.Datetime = String2Time(data.ChatDate)
		// fix bug : 不写，date类型会自动减少8小时
		repairHour, _ := time.ParseDuration("8h")
		latestRecord.Datetime = latestRecord.Datetime.Add(repairHour)

		result, err := ormDb.NewInsert().Model(&latestRecord).
			On("CONFLICT (friend_id) DO UPDATE").
			Exec(ctx)
		if err != nil {
			log.Println("fail to put new record in table chat_latest, err=", err)
			ok = false
		} else {
			rows, err := result.RowsAffected()
			if err != nil || int(rows) != 1 {
				ok = false
			}
		}
	}

	return
}

// 将一条最新交流记录插入到历史记录表中
func InsertChatHistory(lastRecord *table.ChatLatest) (int64, bool) {
	ctx := context.Background()
	ormDb := bun.NewDB(DB2pg, pgdialect.New())

	// 不重复添加
	oldChatList, total := QueryChatHistory(lastRecord.FriendId)
	if total != 0 {
		for _, record := range oldChatList {
			if record.Content == lastRecord.Content {
				log.Println("have back up at past")
				return record.ID, true
			}
		}
	}

	oldRecord := table.ChatHistory{
		Ts:        lastRecord.Ts,
		FriendId:  lastRecord.FriendId,
		PetName:   lastRecord.PetName,
		Datetime:  lastRecord.Datetime,
		ChatTopic: lastRecord.ChatTopic,
		Content:   lastRecord.Content,
		Remark:    lastRecord.Remark,
		Creator:   lastRecord.Creator,
	}
	if _, err := ormDb.NewInsert().Model(&oldRecord).Exec(ctx); err != nil {
		log.Println("fail to put a record in table `chat_history`, err=", err)
		return 0, false
	}

	checkRecord := new(table.ChatHistory)
	ormDb.NewSelect().Model(checkRecord).Column("id").Where("ts = ?", oldRecord.Ts).
		Scan(ctx)
	return checkRecord.ID, true
}

/*
功能: 增加一条最新交流记录

	@description: 如果没有聊过，直接增加到最新记录表；
	如果聊过，先确认有没有当天同类话题，有就合并，没有就将旧记录移入到历史表
*/
func AddChatLatest(newData ChatData) bool {
	ctx := context.Background()
	ormDb := bun.NewDB(DB2pg, pgdialect.New())

	if exist := CheckPetName(newData.PetName); !exist {
		fmt.Println("not a old friend, you have to put info in friend table.")
		return false
	}
	if lastRecord, exist := CheckChatLatest(newData.PetName); exist {
		// 同一天交流的内容，是同类话题合并，不是同类话题的就移动到历史
		if lastRecord.Datetime.Format("2006-01-02") == strings.Split(newData.ChatDate, " ")[0] &&
			lastRecord.ChatTopic == newData.ChatTopic {

			if !hasDuplicate(lastRecord.Content, newData.Content) {
				newData.Content = contentCombine(lastRecord.Content, newData.Content)
			} else {
				fmt.Printf("新记录的内容 %v 是最新记录表里的重复数据, 忽略。", newData.Content)
				return false
			}
		} else if _, ok := InsertChatHistory(lastRecord); !ok {
			return false
		}
		_, err := ormDb.NewDelete().Model(lastRecord).WherePK().Exec(ctx)
		if err != nil {
			log.Println("fail to delete: ", err)
		}
	}
	log.Printf("A new talk record will be added, data: %+v", newData)
	return InsertChatLatest(newData)
}

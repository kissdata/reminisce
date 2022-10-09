// 输出的数据
package common

type PGInfo struct {
	IP       string
	Port     int
	UserName string
	Password string
	DbName   string
}

var PG14info = PGInfo{
	"192.168.126.129",
	15432,
	"postgres",
	"like2022",
	"postgres",
}

type OutputDefault struct {
	PetName        string
	FirstKnowTime  string // 初识的时间
	AcquaintedTime string // 结交的时间
	LatestTime     string // 上次交谈时间
	TalkTopic      string // 交流话题
	Content        string // 交谈内容

}

type FriendInfo struct {
}

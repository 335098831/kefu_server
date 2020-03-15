package models

// Robot struct
type Robot struct {
	ID               int64  `orm:"auto;pk;type(bigint);column(id)" json:"id"`                        // 机器人ID
	NickName         string `orm:"unique;type(char);column(nickname)" json:"nickname"`               // 机器人名称
	Avatar           string `orm:"type(char);column(avatar)" json:"avatar"`                          // 机器人头像
	Welcome          string `orm:"column(welcome)" json:"welcome"`                                   // 欢迎语
	Understand       string `orm:"column(understand)" json:"understand"`                             // 不明白语句
	Artificial       string `orm:"column(artificial)" json:"artificial"`                             // 关键词转人工
	KeyWord          string `orm:"column(keyword)" json:"keyword"`                                   // 知识库默认匹配词
	TimeoutText      string `orm:"column(timeout_text)" json:"timeout_text"`                         // 超时提示语
	NoServices       string `orm:"column(no_services)" json:"no_services"`                           // 无人工在线提示语
	LoogTimeWaitText string `orm:"column(loog_time_wait_text)" json:"loog_time_wait_text"`           // 长时间等待提示
	Switch           int    `orm:"default(0);column(switch)" json:"switch"`                          // 是否开启
	System           int    `orm:"default(0);column(system)" json:"system"`                          // 系统内置
	Platform         int64  `orm:"type(bigint);column(platform)" json:"platform"`                    // 服务那个平台
	Delete           int    `orm:"default(0);column(delete)" json:"delete"`                          // 是否已删除 0 false  1 true
	UpdateAt         int64  `orm:"type(bigint);column(update_at)" json:"update_at"`                  // 更新时间
	CreateAt         int64  `orm:"auto_now_add;type(int64);null;column(create_at)" json:"create_at"` // 创建时间
}

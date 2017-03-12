package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Message struct {
	Messageid  string `orm:"pk"`
	Senderid   int64
	Receiverid int64
	Createtime string
	Content    string
}

type Notice struct {
	Noticeid     string `orm:"pk"`
	Receivertype int64
	Createtime   string
	Status       int
	Title        string
	Content      string
}

//用于查询消息返回
type MessageList struct {
	Flag int64
	Msgs []Message
}

func (m *Message) TableName() string {
	return TableName("message")
}

func GetMessageById(userId int64) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("message"))

	cond := orm.NewCondition()

	cond1 := cond.And("receiverid", userId)
	num, _ := qs.SetCond(cond1).Count()
	return num
}

func GetMessageList(userId int64) (Msgl MessageList) {
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("message"))

	cond := orm.NewCondition()
	cond1 := cond.And("receiverid", userId)

	var msgs []Message
	_, err := qs.SetCond(cond1).OrderBy("-createtime").Limit(5).All(&msgs)

	if err != nil {
		logs.Error(err.Error())
		Msgl.Flag = 0
	}
	Msgl.Msgs = msgs
	Msgl.Flag = 1

	return Msgl
}

func (n *Notice) TableName() string {
	return TableName("notice")
}

// func GetNoticeById(userId int64) int64 {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable(TableName("notice"))

// 	cond := orm.NewCondition()
// 	cond.And("receivertype__in", getPositionById(userId))
// 	num, _ := qs.SetCond(cond).Count()
// 	return num
// }

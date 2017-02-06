package models

import (
	"log"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64        `orm:"pk;column(userid);"`
	Profile  *UserProfile `orm:"rel(one);"`
	Username string
	Password string
	Status   int
}

type UserProfile struct {
	Id         int64 `orm:"pk;column(userid);"`
	Realname   string
	Sex        int
	Birth      string
	Email      string
	Phone      string
	Address    string
	Positionid int64
	lognum     int64
	Ip         string
	Lasted     int64
}

func init() {
	orm.RegisterModel(new(User), new(UserProfile))

}

func (u *User) TableName() string {
	return TableName("user")
}

func (p *UserProfile) TableName() string {
	return TableName("UserProfile")
}

func LoginUser(username, password string) (User, error) {
	orm.Debug = true
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("user"))
	cond := orm.NewCondition()
	cond = cond.And("username", username).And("password", password).And("status", 1)
	var user User
	err := qs.SetCond(cond).One(&user)

	if err != nil {
		log.Println(err.Error())
	}

	/**
		user := User{Username: username, Password: password}
		err := o.Read(&user)

		if err == orm.ErrNoRows {
			return user, err
		} else if err == orm.ErrMissPK {
			return user, err
		}
	**/
	return user, err

}

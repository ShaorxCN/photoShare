package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"sync"
	"time"
)

var seq = 0
var once sync.Once
var idChannel chan int

type User struct {
	Id       int64        `orm:"pk;column(userid);"`
	Profile  *UserProfile `orm:"rel(one);"`
	Username string
	Password string
	Status   int
}

type UserProfile struct {
	Id       int64 `orm:"pk;column(userid);"`
	Realname string
	Sex      int
	Birth    string
	Email    string
	Phone    string
	Address  string
	Position string
	Lognum   int64
	Ip       string
	Lasted   int64
}

func init() {
	orm.RegisterModel(new(User), new(UserProfile), new(Message), new(Notice))

}

func (u *User) TableName() string {
	return TableName("user")
}

func (p *UserProfile) TableName() string {
	return TableName("user_profile")
}

func LoginUser(username, password string) (User, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(TableName("user"))
	cond := orm.NewCondition()
	cond = cond.And("username", username).And("password", password)
	var user User
	err := qs.SetCond(cond).One(&user)

	if err != nil {
		logs.Error(err.Error())
		return user, errors.New("登录失败,请确认账户密码")
	}

	if user.Status == 3 {
		profile := new(UserProfile)
		profile.Id = user.Id
		_ = o.Read(profile)

		timePoint, _ := strconv.ParseInt(time.Now().Add(time.Second*-300).Format("20060102150405"), 10, 64)
		if timePoint >= profile.Lasted {
			return user, nil
		}
		return user, errors.New("该帐号已经登录,请核实或联系管理员")
	}

	user.Status = 3
	_, err = o.Update(&user)
	if err != nil {
		err = errors.New("登录状态异常")
	}

	return user, err

}

func checkErr(err error, message string) error {
	if err != nil {
		log.Println(err)
		return errors.New("message")
	} else {
		return nil
	}
}

func singleIdChanel() chan int {
	if idChannel == nil {
		once.Do(func() { idChannel = make(chan int, 100) })
		go userIdPro(idChannel)
	}

	return idChannel
}

func userIdPro(idChannel chan int) {
	for {
		if seq == 9999 {
			seq = 1
		} else {
			seq++
		}

		idChannel <- seq
	}
}

//get userId
func getUserId() (int64, error) {
	idChannel := singleIdChanel()

	var pre = time.Now().Format("20060102150405")
	seqt := <-idChannel

	res, err := strconv.ParseInt(pre+strconv.Itoa(seqt), 10, 64)

	if err != nil {
		return 0, errors.New("获取用户id失败")
	}

	return res, nil
}

//get ProfileId
// func GetProfileId(id int64) (int64, error) {
// 	profileId, err := strconv.ParseInt("1"+strconv.FormatInt(id, 10), 10, 64)
// 	return profileId, err
// }

func RegisterUser(username, password string) (*User, error) {
	user := new(User)
	id, err := getUserId()
	if err != nil {
		return user, err
	}

	o := orm.NewOrm()
	exist := o.QueryTable(TableName("user")).Filter("username", username).Exist()
	if exist {
		return user, errors.New(fmt.Sprintf("用户 %s 已存在,请更换用户名.", username))
	} else {

		user.Username = username
		user.Password = password
		user.Status = 1
		user.Id = id

		profile := new(UserProfile)
		profile.Id = id
		profile.Position = "1"
		user.Profile = profile
		err = o.Begin()
		if err != nil {
			log.Println(err)
			return user, errors.New("注册失败")
		}
		_, err = o.Insert(user)
		if err != nil {
			log.Println(err)
			err = o.Rollback()
			log.Println(err)
			return user, errors.New("注册失败")
		}
		_, err = o.Insert(profile)
		if err != nil {
			log.Println(err)
			err = o.Rollback()
			log.Println(err)
			return user, errors.New("注册失败")
		}

	}
	err = o.Commit()
	if err != nil {
		log.Println(err)
		return user, errors.New("注册失败")
	}
	return user, nil
}

func GetNickById(userId int64) string {
	o := orm.NewOrm()
	user := new(User)

	user.Id = userId

	err := o.Read(user)

	if err != nil {
		return "nil"
	}

	return user.Username
}

func ChangeLoginStatus(user User) error {
	o := orm.NewOrm()

	user.Status = 3

	_, err := o.Update(&user)

	return err
}

func SignOut(userId int64) error {
	o := orm.NewOrm()
	user := new(User)
	user.Id = userId
	err := o.Read(user)
	if err != nil {
		return err
	}

	user.Status = 1
	_, err = o.Update(user)

	return err

}

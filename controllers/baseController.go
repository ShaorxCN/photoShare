package controllers

import (
	"github.com/astaxie/beego"
	//"net/http"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"photoShare/models"
	"strconv"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller

	IsLogin      bool
	UserUserId   int64
	UserUserName string
}

func (b *BaseController) Prepare() {
	isLogin := b.GetSession("isLogin")
	if isLogin == nil {
		b.IsLogin = false
	} else {
		temp := strings.Split(b.GetSession("isLogin").(string), "||")
		if 2 == len(temp) {
			b.IsLogin = true
			b.UserUserId, _ = strconv.ParseInt(temp[0], 10, 64)
			b.UserUserName = temp[1]

			o := orm.NewOrm()
			user := new(models.User)

			user.Id = b.UserUserId
			err := o.Read(user)

			if err != nil {

				b.Data["error"] = "用户信息查询失败,请重新<a href='/login'>登录</a>"

				b.Redirect("/error", 302)
			}
			profile := new(models.UserProfile)
			profile.Id = user.Id

			_ = o.Read(profile)
			profile.Lasted, _ = strconv.ParseInt(time.Now().Format("20060102150405"), 10, 64)
			logs.Info(profile.Lasted)
			_, err = o.Update(profile)
			if err != nil {
				logs.Error(err.Error())
			}
			b.Data["user"] = user
		}
	}

}

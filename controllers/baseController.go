package controllers

import (
	"github.com/astaxie/beego"
	//"net/http"
	//"log"
	"strconv"
	"strings"
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
		}
	}

}

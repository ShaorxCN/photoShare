package routers

import (
	"photoShare/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &user.MainController{})

	//登录模块
	beego.Router("/login", &user.LoginInController{})

	//用户模块
	beego.Router("/user/:id", &user.UserController{})

}

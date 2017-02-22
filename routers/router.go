package routers

import (
	"github.com/astaxie/beego"
	"photoShare/controllers/account"
	"photoShare/controllers/user"
)

func init() {
	//主页
	beego.Router("/", &user.MainController{})

	//登录模块
	beego.Router("/login", &user.LoginInController{})

	//用户模块
	beego.Router("/user/:id([0-9]+)", &user.UserController{})

	//账户管理模块
	beego.Router("/register", &account.RegisterController{})
	beego.Router("/profile/:id([0-9]+)", &account.ProfileController{})

}

package routers

import (
	"github.com/astaxie/beego"
	"photoShare/controllers/account"
	"photoShare/controllers/errorCtr"
	"photoShare/controllers/user"
)

func init() {
	//主页
	beego.Router("/", &user.MainController{})

	//登录模块
	beego.Router("/login", &user.LoginInController{})
	beego.Router("/loginOut", &user.LoginOutController{})

	//用户模块
	beego.Router("/user/:id([0-9]+)", &user.UserController{})

	//搜索
	beego.Router("/search", &user.SearchController{})

	//账户管理模块
	beego.Router("/register", &account.RegisterController{})
	beego.Router("/profile/:id([0-9]+)", &account.ProfileController{})

	//错误处理模块
	beego.Router("/error", &errorCtr.ErrorController{})

}

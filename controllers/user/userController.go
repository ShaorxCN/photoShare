package user

import (
	//"github.com/astaxie/beego"
	"fmt"
	"log"
	"photoShare/controllers"
	"photoShare/models"
	"strconv"
)

type MainController struct {
	controllers.BaseController
}

func (this *MainController) Get() {
	check := this.BaseController.IsLogin

	if check {
		this.Redirect(fmt.Sprintf("/user/%s", strconv.FormatInt(this.UserUserId, 10)), 301)
	} else {
		this.Redirect("/login", 301)
	}

}

type LoginInController struct {
	controllers.BaseController
}

func (this *LoginInController) Get() {
	this.TplName = "users/login.tpl"

}

func (this *LoginInController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	user, err := models.LoginUser(username, password)

	if err != nil {
		log.Println(err.Error())
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败，请您确认账户或者密码"}
	} else {

		this.SetSession("isLogin", strconv.FormatInt(user.Id, 10)+"||"+username)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "恭喜登录您登录成功," + user.Username, "redirect": fmt.Sprintf("user/%d", user.Id)}
	}
	this.ServeJSON()
}

type UserController struct {
	controllers.BaseController
}

func (this *UserController) Get() {
	check := this.BaseController.IsLogin

	if !check {
		this.Redirect("/login", 301)

	} else {
		this.TplName = "users/user.tpl"
	}
}

package account

import (
	//"log"
	"photoShare/controllers"
	"photoShare/models"
)

type RegisterController struct {
	controllers.BaseController
}

func (this *RegisterController) Get() {
	this.TplName = "account/register.tpl"
}

func (this *RegisterController) Post() {
	//todo
	username := this.GetString("username")
	password := this.GetString("password")
	_, err := models.RegisterUser(username, password)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": err.Error()}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "恭喜您注册成功", "redirect": "login"}
	}

	this.ServeJSON()
}

package account

import (
	//"log"
	"photoShare/controllers"
	//"photoShare/models"
)

type RegisterController struct {
	controllers.BaseController
}

func (this *RegisterController) Get() {
	this.TplName = "account/register.tpl"
}

func (this *RegisterController) Post() {
	//todo
	this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	this.ServeJSON()
}

package errorCtr

import (
	"github.com/astaxie/beego"
	"html/template"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Get() {
	this.Data["error"] = template.HTML("查询信息失败,请重新<a href = '/login'>登录</a>")
	this.TplName = "error/error.html"
}

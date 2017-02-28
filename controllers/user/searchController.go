package user

import (
	//"fmt"
	//"github.com/astaxie/beego"
	"html/template"
	//"github.com/astaxie/beego/orm"
	"log"
	"photoShare/controllers"
	//"photoShare/models"
	//"strconv"
)

type SearchController struct {
	controllers.BaseController
}

func (this *SearchController) Get() {
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	key := this.GetString("searchKey")
	log.Println(key)

	this.Redirect("/", 302)

}

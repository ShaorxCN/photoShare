package account

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	//"net/http"
	"photoShare/controllers"
	"photoShare/models"
	"strconv"
)

type RegisterController struct {
	controllers.BaseController
}

func (this *RegisterController) Get() {
	this.Data["xsrf_token"] = this.XSRFToken()
	this.TplName = "account/register.html"
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

type ProfileController struct {
	controllers.BaseController
}

func (this *ProfileController) Get() {
	this.Data["xsrf_token"] = this.XSRFToken()
	check := this.BaseController.IsLogin
	if !check || this.Ctx.Input.Param(":id") != strconv.FormatInt(this.BaseController.UserUserId, 10) {
		this.Redirect("/login", 302)
	}

	userId := this.BaseController.UserUserId
	userProfile := new(models.UserProfile)

	userProfile.Id = userId

	o := orm.NewOrm()
	err := o.Read(userProfile)
	if err != nil {
		log.Println(err)
		this.Data["response"] = map[string]interface{}{"code": 0, "profile": userProfile}
	} else {
		this.Data["response"] = map[string]interface{}{"code": 1, "profile": userProfile}
	}
	this.TplName = "account/profile.tpl"
}

func (this *ProfileController) Post() {

	check := this.BaseController.IsLogin
	if !check || this.Ctx.Input.Param(":id") != strconv.FormatInt(this.BaseController.UserUserId, 10) {
		this.Redirect("/login", 302)
	}

	realname := this.GetString("realname")
	sex := this.GetString("sex")
	birth := this.GetString("birth")
	phone := this.GetString("phone")
	email := this.GetString("email")
	address := this.GetString("address")

	userId := this.Ctx.Input.Param(":id")
	//userProfile := new(models.UserProfile)
	userProfile := new(models.UserProfile)
	id, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {
		log.Println(err)
		this.Data["json"] = map[string]interface{}{"code": 0, "profile": userProfile, "message": "服务异常,请稍后再试"}

	} else {

		userProfile.Id = id
		o := orm.NewOrm()
		err = o.Read(userProfile)
		if err != nil {
			log.Println(err)
			this.Data["json"] = map[string]interface{}{"code": 0, "profile": userProfile, "message": "服务异常,请稍后再试"}
		} else {
			userProfile.Address = address
			userProfile.Birth = birth
			userProfile.Email = email
			userProfile.Phone = phone
			userProfile.Realname = realname
			sexInt, err := strconv.Atoi(sex)
			userProfile.Sex = sexInt
			_, err = o.Update(userProfile)

			if err != nil {
				log.Println(err)
				this.Data["json"] = map[string]interface{}{"code": 0, "message": "服务异常,请稍后再试"}
			} else {
				this.Data["json"] = map[string]interface{}{"code": 1, "message": "资料填充完成"}
			}
		}
	}
	this.ServeJSON()
}

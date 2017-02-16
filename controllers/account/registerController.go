package account

import (
	"github.com/astaxie/beego/orm"
	"log"
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

type ProfileController struct {
	controllers.BaseController
}

func (this *ProfileController) Get() {

	check := this.BaseController.IsLogin
	if !check {
		this.Redirect("/login", 302)
	}

	userId := this.BaseController.UserUserId
	userProfile := new(models.UserProfile)
	profileId, err := models.GetProfileId(userId)

	if err != nil {
		log.Println(err)
		this.Data["response"] = map[string]interface{}{"code": 0, "profile": userProfile}
	}

	userProfile.Id = profileId

	o := orm.NewOrm()
	err = o.Read(userProfile)
	if err != nil {
		log.Println(err)
		this.Data["response"] = map[string]interface{}{"code": 0, "profile": userProfile}
	} else {
		this.Data["response"] = map[string]interface{}{"code": 1, "profile": userProfile}
	}
	this.TplName = "account/profile.tpl"
}

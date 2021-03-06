package beegoLearn

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get(){
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail"
	this.TplName = "index.html"
}

func StartBeego()  {
	beego.Router("/json",&BaseController{})
	beego.Router("/",&MainController{})
	beego.Router("/index",&AddController{})
	beego.Run("localhost:8081")
}

package routers

import (
	"github.com/astaxie/beego"
	"goTest/beegoLearn"
)

func init()  {
	beego.Router("/json",&beegoLearn.BaseController{})
}


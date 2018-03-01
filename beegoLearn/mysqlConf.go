package beegoLearn

import "github.com/astaxie/beego"

//获取 conf/app.conf 文件中的配置信息

func BeeConf()  {
	beego.AppConfig.String("mysqluser")
}

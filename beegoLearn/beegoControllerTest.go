package beegoLearn

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"github.com/axgle/mahonia"
)


type BaseController struct {
	beego.Controller
}

type User struct {
	Name string
	Age  string
}
func (this *BaseController) Get() {
	userJson,err := json.Marshal(User{"张三","23"})
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(userJson))
	ContentType := this.GetString("Content-Type")
	fmt.Println(ContentType)
	this.Data["json"] = string(userJson)
	this.TplName = "index.html"
	//ctx := this.Ctx
	//ctx.WriteString(string(userJson))
	//this.Ctx.ResponseWriter.Flush()
}

type user struct {
	Id int `form:"-"`
	Name interface{} `form:"username"`
	Age int `form:"age"`
	Email string
}

func (c *BaseController) Post() {
	u := user{}
	c.ParseForm(&u)
	fmt.Println(u)
}



type AddController struct {
	beego.Controller
}

func (c *AddController) Prepare() {

}

func (c *AddController) Get() {

	session := c.GetSession("session")
	if session == nil{
		c.SetSession("session","123456789")
	}
	//c.XSRFExpire = 7200
	//xsrf设置方式一
	//c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	//xsrf设置方式二
	//c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["content"] = "value"
	c.Layout = "index.html"
	c.TplName = "index.html"
	t,_ := template.ParseFiles("views/index.html")
	t.Execute(c.Ctx.ResponseWriter,"hello world")
	c.Ctx.ResponseWriter.Flush()
	t.Clone()

}

func (c *AddController) Post() {
	//c.XSRFExpire = 7200
	//c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	session := c.GetSession("session")
	fmt.Println(session)
	ss := c.StartSession()
	sid := ss.SessionID()
	fmt.Println(sid)
	ss.SessionRelease(c.Ctx.ResponseWriter)
	file,head,err := c.GetFile("file")
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(head.Size)
	byt := make([]byte,1024)
	for {
		_,err := file.Read(byt)
		if err != nil {
			fmt.Println(err)
		}
		if err == io.EOF {
			break
		}
		enc := mahonia.NewEncoder("utf8")
		result := enc.ConvertString(string(byt))
		fmt.Println(result)
	}
	defer file.Close()

}



package template

import (
	"net/http"
	"html/template"
	"fmt"
)

type Article struct {
	Title   string 	   //标题
	Content string     //内容
	Author  string     //作者
	Tab     []string   //标签
	PublishTime string //发表时间
	ViewNum int        //浏览量
}

func EditHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("===========")
	article := Article{Title:"标题",Content:"内容",Author:"作者",Tab:[]string{"练习","实践"},PublishTime:"2018-1-13",ViewNum:1}
	t,err :=template.ParseFiles("template/template.html")
	if err!=nil {
		fmt.Println(err)
		fmt.Fprintf(w,"页面没有准备好，请稍后访问。。。。。")
		return
	}
	t.Execute(w,article)
}


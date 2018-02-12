package template

import (
	"net/http"
	"fmt"
	"html/template"
	"time"
)

func ForbiddenRepeat()  {
	mux := http.NewServeMux()
	mux.Handle("/login",&myRoutor{})
	http.ListenAndServe(":8080",mux)
}

type myRoutor struct {
	name string
}

func (my *myRoutor) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	if r.Method =="GET" {
		t,_ := template.ParseFiles("template/login.html")
		t.Execute(w,nil)
	} else {
	values := r.Form
		for k,v := range values{
		fmt.Println(k,"====",v)
			if k == "Token"&&v == "" {
				time.Now()
			}
		}
	}
	defer r.Body.Close()

}
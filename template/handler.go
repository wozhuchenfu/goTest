package template

import (
	"net/http"
	"fmt"
)

func Handler()  {
	serverMux := http.NewServeMux()
	serverMux.Handle("/",&MyHandler{})
	http.ListenAndServe(":8080",serverMux)

	//http.HandleFunc("/admin",myHandler)
	//http.ListenAndServe(":8080",nil)
}

type MyHandler struct {

}

func (myHandler *MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"<a href='https://www.baidu.com'>百度</a>")
	fmt.Println(r.Method)
	EditHandler(w,r)

}
func myHandler(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	requestMethod := r.Method
	fmt.Println(requestMethod)
	fmt.Fprintf(w,"%s","<a herf=''>hahha</a>")
}



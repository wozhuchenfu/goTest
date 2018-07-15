package baseLearn

import (
	"net/http"
	"fmt"
	"log"
	"context"
	"time"
	"strconv"
	"io"
)

func GoContextTest() {
	mux := http.NewServeMux()
	mux.Handle("/hello",&goContextServer{"goContextServer"})
	mux.HandleFunc("/login",loginHandleFunc)
	goContextServer := AddContextSupport(mux)
	http.ListenAndServe(":9090",goContextServer)
}

func loginHandleFunc(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	defer r.Body.Close()
	fmt.Println(r.Form)
}

func LoginHandler(w http.ResponseWriter,r *http.Request){
	expitation := time.Now().Add(24*time.Hour)
	var username string
	if username=r.URL.Query().Get("username");username==""{
		username = "guest"
	}
	cookie:=http.Cookie{Name:"username",Value:username,Expires:expitation}
	http.SetCookie(w,&cookie)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().AddDate(0, 0, -1)
	cookie := http.Cookie{Name: "username", Value: "alice_cooper@gmail.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func StatusHandler(w http.ResponseWriter,r *http.Request){

	if username:=r.Context().Value("username"); username!=nil{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hi username:"+username.(string)+"\n"))
	}else{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}
type goContextServer struct {
	serverName string
}

func (g *goContextServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	defer r.Body.Close()
	w.Write([]byte("hello,world!"))
}
/*
AddContextSupport是一个中间件，用来绑定一个context到原来的handler中，
所有的请求都必须先经过该中间件后才能进入各自的路由处理中
 */
func AddContextSupport(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("username")
		if cookie != nil {
			ctx := context.WithValue(r.Context(), "username", cookie.Value)
			// WithContext returns a shallow copy of r with its context changed
			// to ctx. The provided ctx must be non-nil.
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
/*
超时处理
对于简单的保存和传递对象，使用context的确很方便，但是该库的使用不仅仅是保存变量，
还可以创建一个超时和取消的行为，比如说我们web端去请求了其他的资源，但是该资源的处理比较耗时，
我们无法预见什么时候能够返回，如果让用户超时的话，实在是不太好，所以我们需要创建一个超时的操作，
主动判断是否超时，然后传递一个合适的行为给用户。

这里我们现在路由中增加一个长期运行的job路由

mux.HandleFunc("/longjob",jobWithCancelHandler)
1
具体的处理如下，我们的handler会利用WithCancel() 返回一个新的（如果没有创建）或者原来已保存的上下文，
还有一个cancel对象，这个对象可以用来手动执行取消操作。另外我们的url中可以指定这个任务模拟执行的长度，
比如/longjob?jobtime=10则代表模拟的任务将会执行超过10秒。 执行任务的函数longRunningCalculation（）
返回一个chan该chan会在执行时间到期后写入一个Done字符串。

handler中我们就可以使用select语句监听两个非缓存的channel，阻塞直到有数据写到任何一个channel中。
比如代码中我们设置了超时是5秒，而任务执行10秒的话则5秒到期后ctx.Done()会因为cancel()的调用而写入数据，
这样该handler就会因为超时退出。否则的话则执行正常的job处理后获得传递的“Done”退出。
 */
func longRunningCalculation(timeCost int)chan string{

	result:=make(chan string)
	go func (){
		time.Sleep(time.Second*(time.Duration(timeCost)))
		result<-"Done"
	}()
	return result
}
func jobWithCancelHandler(w http.ResponseWriter, r * http.Request){
	var ctx context.Context
	var cancel context.CancelFunc
	var jobtime string
	if jobtime=r.URL.Query().Get("jobtime");jobtime==""{
		jobtime = "10"
	}
	timecost,err:=strconv.Atoi(jobtime)
	if err!=nil{
		timecost=10
	}
	log.Println("Job will cost : "+jobtime+"s")
	ctx,cancel = context.WithCancel(r.Context())
	defer cancel()

	go func(){
		time.Sleep(5*time.Second)
		cancel()
	}()

	select{
	case <-ctx.Done():
		log.Println(ctx.Err())
		return
	case result:=<-longRunningCalculation(timecost):
		io.WriteString(w,result)
	}
	return
}
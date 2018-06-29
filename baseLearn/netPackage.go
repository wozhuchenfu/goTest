package baseLearn

import (
	"log"
	"net"
	"io/ioutil"
	"fmt"
	"time"
	"net/http"
	"html/template"
	"io"
	"os"
	"encoding/json"
)

func chkError(err error)  {
	if err != nil{
		log.Fatal(err)
	}
}
func TcpRequest()  {
	//模拟网易服务器
	//net参数是"tcp4","tcp6","tcp"
	//addr表示域名或IP地址加端口
	tcpaddr,err := net.ResolveTCPAddr("tcp4","www.163.com:80")
	fmt.Println("访问地址:",tcpaddr)
	chkError(err)
	//DialTCP建立一个TCP链接
	//net参数是"tcp4","tcp6","tcp"
	//laddr表示远程地址
	tcpconn,err2 := net.DialTCP("tcp",nil,tcpaddr)
	chkError(err2)
	//向tcpconn中写入数据
	_,err3 := tcpconn.Write([]byte("GET / HTTP/1.1 \r\n\r\n"))
	chkError(err3)
	//读取tcpconn中的所有数据
	data,err4 := ioutil.ReadAll(tcpconn)
	chkError(err4)
	//打印数据
	fmt.Println(string(data))
}

func TcpServer()  {
	//创建一个TCP服务
	tcpaddr,err := net.ResolveTCPAddr("tcp4","127.0.0.1:8080")
	chkError(err)
	//监听端口
	tcplisten,err2 := net.ListenTCP("tcp",tcpaddr)
	chkError(err2)
	//处理客户端请求
	for {
		//等待客户端连接（这里无法并发处理多个请求）
		conn,err3 := tcplisten.Accept()
		chkError(err3)
		if err3 != nil {
			continue
		}
		//向客户端发送数据，并关闭链接
		conn.Write([]byte("hello,client \r\n"))
		conn.Close()
	}
}

//对客户端改进
func TcpServer2()  {
	//创建一个TCP服务端
	tcpaddr,err := net.ResolveTCPAddr("tcp4","127.0.0.1:8080")
	chkError(err)
	//监听端口
	tcplisten,err2 := net.ListenTCP("tcp",tcpaddr)
	chkError(err2)
	for {
		conn,err3 := tcplisten.Accept()
		chkError(err3)
		if err3 != nil{
			continue
		}
		go clientHandle(conn)
	}


}

func clientHandle(conn net.Conn)  {
	defer conn.Close()
	//设置3分钟内无数据请求时，自动关闭conn
	conn.SetDeadline(time.Now().Add(time.Minute*3))
	n,_ := conn.Write([]byte("hello"+time.Now().String()))
	fmt.Println(n)
}

//Go http
/*
Go语言中处理HTTP请求主要跟两个东西相关：ServerMux和Handler。
ServerMux本质上是一个HTTP请求路由器（或者叫多路复用器，Multiple）他把收到的一组请求与预先定义的URL路径
列表做对比，然后在匹配到路径的时候嗲欧勇关联的处理器（Handler）。
处理器（Handler）负责输出HTTP响应的头和正文。任何满足了http.Handler接口的对象都可以作为一个处理器。
通俗的说，对象只要有个如下签名的ServerHTTP方法即可
 */
func HttpTest()  {
	serverMux := http.NewServeMux()//创建一个空的ServerMux
	rh := http.RedirectHandler("http://example.org",307)//创建一个新的处理器，这个处理器会对收到的所有请求，都执行307重定向操作到http://example.org。
	serverMux.Handle("/foo",rh)//处理器注册到新创建的ServerMux，所以它再URL路径/foo上接收到所有的请求都交给这个处理器。
	log.Println("Listening...")
	http.ListenAndServe(":9090",serverMux)//函数监听所有进入的请求，通过传递刚才创建的ServerMux来为请求去匹配对应的处理器

}

type IndexUser struct {
	Name string
	Age int
}

func customerHandler(w http.ResponseWriter,r *http.Request)  {
	err := r.ParseForm()
	if err!=nil {
		fmt.Println(err)
	}
	byts := []byte{}
	r.Body.Read(byts)
	fmt.Println(string(byts))
	t,err := template.ParseFiles("views/index.html")
	if err!=nil {
		fmt.Println(err)
	}
	indexUser := IndexUser{Name:"张三",Age:34}
	t.Execute(w,indexUser)
}

//创建一个自定义的处理器，功能是将以特定格式输出当前的本地时间
type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is:"+ tm))
}


func CustomerHandler()  {
	mux := http.NewServeMux()

	th := &timeHandler{format:time.RFC1123}
	mux.Handle("/time",th)
	log.Println("Listening...")
	http.ListenAndServe(":9090",mux)
}

func timeHandler2(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func defaultHandler() {
	// Note that we skip creating the ServeMux...

	var format = time.RFC1123
	th := timeHandler2(format)

	// We use http.Handle instead of mux.Handle...
	http.Handle("/time", th)

	log.Println("Listening...")
	// And pass nil as the handler to ListenAndServe.
	http.ListenAndServe(":3000", nil)
}

func RunaHttpTest() {
	fmt.Println("=============")
	http.HandleFunc("/index",handle)
	http.ListenAndServe(":8080",nil)

}

func handle(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	t,err:= template.ParseFiles("views/index.html")
	if err!=nil {
		fmt.Println(err)
	}

	t.Execute(w,"httpTest")
	defer r.Body.Close()
}

type CustomerServerHttp struct {
	ReponseValue interface{}
}

func (c *CustomerServerHttp) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	//fmt.Fprint(w,c.ReponseValue)
	err := r.ParseForm()
	if err!=nil {
		log.Fatal(err)
	}
	t,err := template.ParseFiles("views/index.html")
	if err!=nil && err!=io.EOF {
		log.Fatal(err)
	}
	//repsons := &IndexUser{Name:"张三",Age:34}
	t.Execute(w,c.ReponseValue)


	//传递json数据ajax传递数据类型设置成json返回数据也要设置为json格式才能被ajax接收
	w.Header().Add("Content-Type","application/json")
	m := make(map[string]string)
	m["key"] = "hello"
	b,err:=json.Marshal(m)
	if err!=nil{
		fmt.Println(err)
	}
	w.Write(b)


}

func CustomerHTTPTest()  {
	customer := &CustomerServerHttp{ReponseValue:IndexUser{Name:"张三",Age:34}}
	serverMux := http.NewServeMux()
	serverMux.Handle("/index",customer)
	http.ListenAndServe(":9090",serverMux)
}


type CustomerServMux struct {

}

//实现ServeHTTP(w ResponseWriter, r *Request)
func (c *CustomerServMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		err:=r.ParseForm()
		fmt.Println("======================")
	if err!=nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	w.WriteHeader(200)
	f,err := os.Open("C:\\Users\\edianzu\\Desktop\\人生六戒.txt")
	if err!=nil || err==io.EOF {
		log.Fatal(err)
	}
	bys,err:=ioutil.ReadAll(f)
	fmt.Fprint(w,string(bys))
	w.Write([]byte("hello"))
}

func HttpTesat()  {
	c := &CustomerServMux{}//自定义处理器
	mux := http.NewServeMux()//路由
	mux.Handle("/hello",c)
	//路由注册
	http.ListenAndServe(":9090",mux)
}
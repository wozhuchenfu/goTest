package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
	"reflect"
	"goTest/baseLearn"
	"os"
	//"goTest/template"
	//"goTest/sessionHandler"
	//"net/http"
	//"html/template"
	"goTest/goconfig"
	"github.com/davecgh/go-spew/spew"
)

/*var globalSessions *sessionHandler.SessionManager

func init() {
	globalSessions,_ = sessionHandler.NewManager("memeory","gosessionid",3600)
	//go globalSessions.GC()
}

func Count(w http.ResponseWriter,r *http.Request)  {
	sess := globalSessions.SessionStart(w,r)
	createTime := sess.Get("createtime")
	if createTime == nil{
		sess.Set("createtime",time.Now().Unix())
	} else if (createTime.(int64)+360 <(time.Now().Unix())) {
		globalSessions.SessionDestory(w,r)
		sess = globalSessions.SessionStart(w,r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum",1)
	} else {
		sess.Set("countnum",(ct.(int)+1))
	}
	t,_ := template.ParseFiles("template/countnum.html")
	w.Header().Set("Content-Type","text/html")
	t.Execute(w,sess.Get("countnum"))
}

func Login(w http.ResponseWriter,r *http.Request)  {
	sess := globalSessions.SessionStart(w,r)
	r.ParseForm()
	if r.Method == "GET" {
		t,_ := template.ParseFiles("template/session.html")
		w.Header().Set("Content-Type","text/html")
		t.Execute(w,sess.Get("username"))
	} else {
		sess.Set("username",r.Form["username"])
		http.Redirect(w,r,"/",302)
	}

}*/



func main()  {

	type user struct {
		name string
		age string
	}

	user1 := user{name:"zhangsan",age:"23"}
	spew.Dump(user1)
	fmt.Println("+++++++++++++")
	baseLearn.JsonTest2()
	//fmt.Println(os.Environ())
	//add := os.Getenv("ADDR")
	//fmt.Println(add)
	fmt.Println("============")
	goconfig.ReadConf()
	baseLearn.ListTest()
	/*http.HandleFunc("/count",Count)
	http.HandleFunc("/",Login)
	http.ListenAndServe(":8080",nil)*/
	//template.Handler()
	//handlers.Reads()
	//handlers.Writes()
	//baseLearn.TcpServer2()
	//baseLearn.CustomerHandler()
	fmt.Println(os.Args[0])//打印命令行信息

	baseLearn.PointerTest()
	fmt.Println(os.Args[0])
	/*rangeChannel()
	fmt.Println("hello,world")
	baseLearn.SliceTest()
	n := baseLearn.Fact(10)
	fmt.Println(n)
	fmt.Println(time.Second)*/
	/*done := make(chan bool,1)
	done<-false
	go worker(done)//先给通道赋值通知worker执行否则worker被阻塞
	<-done
	fmt.Println(<-done)*/
	/*deferTest()
	pase_student()
	goroutineTest()
	t := Teacher{}
	t.ShowA()
	exceptionTest()*/

	//继续研究
	/*a:=1
	b:=2
	defer calc("1",a,calc("10",a,b))
	a = 0
	defer calc("2",a,calc("20",a,b))
	b = 1
*/
	/*sliceTest()

	user := UserAges{make(map[string]int),sync.Mutex{}}
	user.Add("zhangsan",23)
	age := user.Get("zhangsan")
	fmt.Println(age)

	//没理解
	c := []int{1,2,3,4,5,6}
	thred := threadSafeSet{sync.RWMutex{}, c}
	fmt.Println(<-thred.Iter())

	//....没理解
	if live() == nil {
		fmt.Println("AAAAAAAA")
	}else {
		fmt.Println("BBBBBBB")
		fmt.Println(live())
	}*/

}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice  {
		panic("toslice arr not slic")
	}
	l := v.Len()
	ret := make([]interface{},l)
	for i :=0;i<l ; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

func test(i int)  {
	var x int = 100
	x = i + 100
	fmt.Println(x)
	var s *int  = &i
	fmt.Println(s)
}

func worker(done chan bool)  {
	time.Sleep(5*time.Second)
	done <- true
}

func rangeChannel()  {
	queue := make(chan string,2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem:=range queue{
		fmt.Println(elem)
	}

}

func deferTest()  {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()
}

type student struct {
	Name string
	Age int
}

func pase_student()  {
	m := make(map[string]*student)
	stus := []student{{Name:"zhou",Age:24},{Name:"li",Age:23},{Name:"wang",Age:22}}
	for _,stu := range stus{
		m[stu.Name] = &stu//将地址赋给了map
		fmt.Println(stu)
	}
	fmt.Println(m)
}
//考察闭包 阻塞
func goroutineTest()  {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i:=0;i<15;i++ {
		go func() {
			fmt.Println("i:",i)
			wg.Done()
		}()
	}
	for i :=0;i<5 ;i++  {
		go func(i int) {
			fmt.Println("i:",i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
//结构考察
type People struct {
}

func (p *People) ShowA()  {
	fmt.Println("showA")

}
func (p *People) ShowB(){
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher ShowB")
}
//异常处理
func exceptionTest()  {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int,1)
	string_chan := make(chan string,1)
	int_chan <-1
	string_chan <-"hello"
	select {
	case value :=<-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)//抛出异常
	}
}

//考察defer 和函数内部调用返回值
func calc(index string,a,b int) int  {
	ret := a+b
	fmt.Println(ret)
	return ret
}

func sliceTest()  {
	s := make([]int ,5)//默认元素赋值为0
	s = append(s,1,2,3)//向数组尾部添加元素顺序为123
	fmt.Println(s)
}

type UserAges struct {
	ages map[string]int//map初始值为nil赋值时需要对map进行初始化
	sync.Mutex
}

func (ua *UserAges) Add(name string,age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age ,ok := ua.ages[name];ok {
		return age
	}
	return -1
}

type threadSafeSet struct {
	sync.RWMutex
	s []int
}
//这样得到的元素应该是重复的s中的最后一个元素
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})//创建的是没有缓冲的通道
	go func() {
		set.RLock()
		for elem := range set.s{
			ch<-elem
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

type People2 interface {
	Speak(string) string
}
type Student2 struct {

}

func(stu *Student2) Speak(think string)(talk string){
	if think == "bitch" {
		talk = "You are a good boy"
	}else {
		talk = "hi"
	}
	return
}

type People3 interface {
	Show()
}
type Student3 struct {

}

func (stu *Student3) Show() {

}
func live() People3 {
	var stu *Student3
	return stu
}






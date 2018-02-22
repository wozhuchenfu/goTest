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
	//"goTest/goconfig"
	//"github.com/davecgh/go-spew/spew"
	//"goTest/NSQ"
	//"goTest/database"
	//"goTest/io"
	_"goTest/beegoLearn"
	"errors"
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

	//template.ForbiddenRepeat()
	//suanfa.SortedByChoice()
	//suanfa.SortedByQuick()
	//suanfa.SortedByInsert()
	//baseLearn.UnsafePointerTest()
	//beegoLearn.StartBeego()
	//io.IOTest()
	//database.MongoDBTest()
	//NSQ.Producer()
	//NSQ.ConsumerTest()
	//err:=baseLearn.Defer()
	//fmt.Println(err.Error())
	//database.RedisTest()
	//baseLearn.DeferTest()
	//database.RedisPoolTest()
	//database.MysqlRegister()
	//type user struct {
	//	name string
	//	age string
	//}
	//
	//user1 := user{name:"zhangsan",age:"23"}
	//spew.Dump(user1)
	//fmt.Println("+++++++++++++")
	//baseLearn.EncodeAndDecod()
	////baseLearn.JsonTest2()
	////fmt.Println(os.Environ())
	////add := os.Getenv("ADDR")
	////fmt.Println(add)
	//fmt.Println("============")
	//goconfig.ReadConf()
	//baseLearn.ListTest()
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
	//deferTest()
	pase_student()
	goroutineTest()
	/*t := Teacher{}
	t.ShowA()
	t.ShowB()
	exceptionTest()*/

	//继续研究
	/*a:=1
	b:=2
	//先执行内部函数调用再按照defer的规则执行函数
	defer calc("1",a,calc("10",a,b))
	a = 0
	defer calc("2",a,calc("20",a,b))
	b = 1*/
	/*sliceTest()

	user := UserAges{make(map[string]int),sync.Mutex{}}
	user.Add("zhangsan",23)
	age := user.Get("zhangsan")
	fmt.Println(age)

	//没理解
	c := []int{1,2,3,4,5,6}
	thred := threadSafeSet{sync.RWMutex{}, c}
	fmt.Println(<-thred.Iter())
*/
	//返回结构的指针不为nil而不是结构数据为nil
	if live() == nil {
		fmt.Println("AAAAAAAA")
	}else {
		fmt.Println("BBBBBBB")
		fmt.Println(live())
	}
	//var peo People2 = Student2{}//People2与Student2是不同类型的数据不能直接等号赋值
	//fmt.Println(size,max_size)
	//常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，
	//println(&cl,cl)
	println(&bl,bl)

	//for i:=0;i<10 ;i++  {
	//	loop :
	//		println(i)
	//}
	//
	////goto不能跳转到其他函数或者内层代码
	//goto loop


	/*
	解析
考点：**Go 1.9 新特性 Type Alias **
基于一个类型创建一个新类型，称之为defintion；基于一个类型创建一个别名，称之为alias。
	MyInt1为称之为defintion，虽然底层类型为int类型，但是不能直接赋值，需要强转； MyInt2称之为alias，可以直接赋值。
	 */
	type MyInt1 int
	type MyInt2 = int
	var i int =9
	//var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i2)

	/*
	考点：变量作用域
	因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量，结果：
	<nil>
	<nil>
	 */
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))

	/*
	for循环复用局部变量i，每一次放入匿名函数的应用都是同一个变量。 结果：

	0xc042046000 2
	0xc042046000 2
	 */
	funs:=test2()
	for _,f:=range funs{
		f()
	}

	/*
	解析
	考点：panic仅有最后一个可以被revover捕获
	触发panic("panic")后顺序执行defer，但是defer中还有一个panic，所以覆盖了之前的panic("panic")

	defer panic
	 */
	/*defer func() {
		if err:=recover();err!=nil{
			fmt.Println("++++")
			f:=err.(func()string)
			fmt.Println(err,f(),reflect.TypeOf(err).Kind().String())
		}else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return  "defer panic"
		})
	}()
	panic("panic")*/

	//list := new([]int)//new方法传给变量的是指针类型变量
	list:=make([]int,0)//make方法传给变量的是类型变量
	list = append(list, 1)
	fmt.Println(list)


	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)//append切片时不要忘了s2后面的三个...
	fmt.Println(s1)


	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	fmt.Println(&sm1,&sm2)
	fmt.Println(sm1.m,sm2.m)

	/*if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}*/

	/*sn3:= struct {
		name string
		age  int
	}{age:11,name:"qq"}*/
	/*
	考点:结构体比较
	进行结构体比较时候，只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
	还有一点需要注意的是结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice。 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。
	 */

	//可以使用reflect.DeepEqual进行比较

	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 ==sm2")
	}else {
		fmt.Println("sm1 !=sm2")
	}

	//考察interface内部结构
	var x1 *int = nil
	Foo(x1)

	fmt.Println(x,y,z,k,p)
	//打印结果为：0,1，zz,4

}

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exists := m[id]; exists {
		return "存在数据", true
	}
	return /*nil*/"", false
	/*
	考点：函数返回值类型
	nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。
	但是如果不特别指定的话，Go 语言不能识别类型，所以会报错。
	报:cannot use nil as type string in return argument.
	 */
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

var ErrDidNotWork = errors.New("did not work")

/*func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		fmt.Println(result)
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err//没有给err赋值返回的err都是nil
}*/
func DoTheThing(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		result, err = tryTheThing()
		fmt.Println(result)
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string,error)  {
	return "",ErrDidNotWork
}

func test2() []func()  {
	var funs []func()
	for i:=0;i<2 ;i++  {
		funs = append(funs, func() {
			println(&i,i)
		})
	}
	return funs
}
/*
如果想不一样可以改为：

func test() []func()  {
    var funs []func()
    for i:=0;i<2 ;i++  {
        x:=i
        funs = append(funs, func() {
            println(&x,x)
        })
    }
    return funs
}
 */

func main1()  {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err)
		}else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
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
	panic("抛出异常")
	/*
	defer 是后进先出。
	panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，
	最后才会执行panic。
	 */
}

type student struct {
	Name string
	Age int
}

func pase_student()  {
	m := make(map[string]*student)
	stus := []student{{Name:"zhou",Age:24},{Name:"li",Age:23},{Name:"wang",Age:22}}
	for _,stu := range stus{
		m[stu.Name] = &stu
		fmt.Println(stu)
	}
	fmt.Println(m)
}
/*

因为for range创建了每个元素的副本，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针，
就会导致错误，在迭代时，返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以值的地址总是相同的，导致结果不如预期。


考点：foreach
解答：
这样的写法初学者经常会遇到的，很危险！ 与Java的foreach一样，都是使用副本的方式。
所以m[stu.Name]=&stu实际上一直指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝。
就像想修改切片元素的属性：

for _, stu := range stus {
    stu.Age = stu.Age+10
}
也是不可行的。 大家可以试试打印出来：

func pase_student() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
    // 错误写法
    for _, stu := range stus {
        m[stu.Name] = &stu
    }

    for k,v:=range m{
        println(k,"=>",v.Name)
    }

    // 正确
    for i:=0;i<len(stus);i++  {
        m[stus[i].Name] = &stus[i]
    }
    for k,v:=range m{
        println(k,"=>",v.Name)
    }
}
 */
//考察闭包 阻塞
func goroutineTest()  {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i:=0;i<10;i++ {
		go func() {
			fmt.Println("i:",i)
			wg.Done()
		}()//i没有传入函数内的i变量为最后一次执行时i的值
	}
	for i :=0;i<10 ;i++  {
		go func(i int) {
			fmt.Println("i:",i)
			wg.Done()
		}(i)//最后一次的i值先打印出来然后按照i传入到函数中按照传入的先后顺序打印（go1.9版本如此）
	}
	wg.Wait()
}
//结构考察
type People struct {
}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()

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
		panic(value)//抛出异常,当程序运行时，如果遇到引用空指针、下标越界或显式调用panic函数等情况，则先触发panic函数的执行，然后调用延迟函数。
	}
}

//考察defer 和函数内部调用返回值
func calc(index string,a,b int) int  {
	ret := a+b
	fmt.Println("=========++++++++++=========")
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
type Student2 struct {}

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

var(
	/*
	考点:变量简短模式
	变量简短模式限制：
	定义变量同时显式初始化
	不能提供数据类型
	只能在函数内部使用
	 */
	//size := 1024//var声明的变量不能用:=
	//max_size = size*2
)

const cl  = 100
var bl    = 123



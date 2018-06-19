package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
	"reflect"
	//"goTest/baseLearn"
	//"os"
	"errors"
	"goTest/sessionHandler"
	"net/http"
	"html/template"
	"goTest/suanfa"
	"goTest/baseLearn"
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/smartystreets/assertions/should"
)

var globalSessions *sessionHandler.SessionManager

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

}

type A struct {
	Name string
}

type B struct {
	Age uint8
}

type C struct {
	Name string `json:"name"`
	Age uint8 `json:"age, omitempty"`
}
/**
json tag 有很多值可以取，同时有着不同的含义，比如：

tag 是 "-"，表示该字段不会输出到 JSON.

tag 中带有自定义名称，那么这个自定义名称会出现在 JSON 的字段名中，比如上面小写字母开头的 name.

tag 中带有 "omitempty" 选项，那么如果该字段值为空，就不会输出到JSON 串中.

如果字段类型是 bool, string, int, int64 等，而 tag 中带有",string" 选项，那么该字段在输出到 JSON 时，会把该字段对应的值转换成 JSON 字符串.
 */

func main()  {
    a := &C{Name:"张三",Age:23}
	byt,err := jsoniter.Marshal(a)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(byt))
	err = jsoniter.Unmarshal(byt,&C{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byt))
	fmt.Println("=========jsoniter=========")

	//处理字符串和数字类型不对的问题了。比如

	var val1 int
	jsoniter.UnmarshalFromString(`100`, &val1)
	fmt.Println(val1)
	var val2 float32
	jsoniter.UnmarshalFromString(`"1.23"`, &val2)
	fmt.Println(val2)
	extra.RegisterTimeAsInt64Codec(time.Microsecond)
	output, err := jsoniter.Marshal(time.Unix(1, 1002))
	should.Equal("1000001", string(output))

	//t := time.Local.String()
	//fmt.Println(t)
	//fmt.Println("时间格式化输出")
	now := time.Now().Format("2006-01-02 03:04:05 PM")
	fmt.Println(now)

	//baseLearn.WGTest()
	//baseLearn.SyncMapTest()
	//baseLearn.SyncCondTest()
	baseLearn.CustomerHTTPTest()

	treeNode := baseLearn.TreeNode{}
	treeNode.SetValue(3)
	treeNode.Print()
	treeNode.SetValueByPtr(4)//指针方法才可以改变对象的值
	treeNode.Print()
	treeNode2 := baseLearn.CreateTreeNode(5)
	treeNode2.Print()
	treeNode2.SetValue(8)
	treeNode2.Print()
	treeNode3 := baseLearn.CreateTreeNode2(9)
	treeNode3.Print()
	treeNode3.SetValue(90)
	treeNode3.Print()


	fmt.Println("本机逻辑CPU核数",runtime.NumCPU())

	baseLearn.RuneLearn()
	maxLength := suanfa.LengthOfNonRepeatingSubStr("qweasdarerasda")
	fmt.Println(maxLength)

	//baseLearn.SingletonTest()
	//baseLearn.CreateCar()
	//var invoker = baseLearn.Invoker{baseLearn.OpenComond{baseLearn.TV{"TV"}}}
	//invoker.Do()
	//var person = baseLearn.New2()
	//fmt.Println(person)
	//person.Name = "张三"
	//person.Age = 123
	//
	//var person2 = baseLearn.New2()
	//fmt.Println(person2.Name,"=====",person2.Age)
	//var p = baseLearn.Pepole{}
	//p.SetAge(12)
	//p.SetName("lisi")
	//fmt.Println(p)

	//baseLearn.TestNew()
	//a := &A{"zhangsan"}
	//b := &B{23}
	//fmt.Println(a,b)
	//var c *A
	//var d *B
	//fmt.Println(c,d)
	//c = &A{"lisi"}
	//d = &B{34}
	//fmt.Println(*c,*d)
	//fmt.Println("============PointerTest================")
	//
	//var err baseLearn.MyError
	//baseLearn.ErrorTest2(err)
	//ginLearn.GinTest1()
	//http.HandleFunc("/img",baseLearn.Test)
	//http.ListenAndServe(":8080",nil)
	//elasticSearchTest.FindAll()
	//elasticSearchTest.FindByName("Smith")
	//elasticSearchTest.FindByFilter(30,"Smith")
	//elasticSearchTest.FindByJson("Smith")
	//elasticSearchTest.Delete("megacorp/employee/4")
	//elasticSearchTest.Search("megacorp/employee/3")
	//elasticSearchTest.Add()
	//baseLearn.HandleImage()
	//beegoLearn.StartBeego()
	//template.ForbiddenRepeat()
	//suanfa.SortedByChoice()
	//suanfa.SortedByQuick()
	//suanfa.SortedByInsert()
	//baseLearn.UnsafePointerTest()
	//baseLearn.RunaHttpTest()
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
	//baseLearn.SliceTest()
	//baseLearn.DiGui(0)
	//baseLearn.GoTest()
	//baseLearn.Bibao()
	//baseLearn.ChannelTest()
	//baseLearn.ChannelTest2()
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
	//fmt.Println(os.Args[0])//打印命令行信息

	//baseLearn.PointerTest()
	//fmt.Println(os.Args[0])
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
	//pase_student()
	//goroutineTest()
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
	//if live() == nil {
	//	fmt.Println("AAAAAAAA")
	//}else {
	//	fmt.Println("BBBBBBB")
	//	fmt.Println(live())
	//}

	/*
	考点：interface内部结构
解答：
很经典的题！ 这个考点是很多人忽略的interface内部结构。 go中的接口分为两种一种是空的接口类似这样：

var in interface{}
另一种如题目：

type People interface {
    Show()
}
他们的底层结构如下：

type eface struct {      //空接口
    _type *_type         //类型信息
    data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type iface struct {      //带有方法的接口
    tab  *itab           //存储type信息还有结构实现方法的集合
    data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type _type struct {
    size       uintptr  //类型大小
    ptrdata    uintptr  //前缀持有所有指针的内存大小
    hash       uint32   //数据hash值
    tflag      tflag
    align      uint8    //对齐
    fieldalign uint8    //嵌入结构体时的对齐
    kind       uint8    //kind 有些枚举值kind等于0是无效的
    alg        *typeAlg //函数指针数组，类型实现的所有方法
    gcdata    *byte
    str       nameOff
    ptrToThis typeOff
}
type itab struct {
    inter  *interfacetype  //接口类型
    _type  *_type          //结构类型
    link   *itab
    bad    int32
    inhash int32
    fun    [1]uintptr      //可变大小 方法集合
}
可以看出iface比eface 中间多了一层itab结构。 itab 存储_type信息和[]fun方法集，从上面的结构我们就可得出，因为data指向了nil 并不代表interface 是nil， 所以返回值并不为空，这里的fun(方法集)定义了接口的接收规则，在编译的过程中需要验证是否实现接口 结果：
	 */


	//var peo People2 = Student2{}
	//think := "bitch"
	//fmt.Println(peo.Speak(think))
	//fmt.Println(size,max_size)
	//常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，
	//println(&cl,cl)
	//println(&bl,bl)

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
	//type MyInt1 int
	//type MyInt2 = int
	//var i int =9
	//var i1 MyInt1 = i
	//var i2 MyInt2 = i
	//fmt.Println(i2)

	/*
	考点：变量作用域
	因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量，结果：
	<nil>
	<nil>
	 */
	//fmt.Println(DoTheThing(true))
	//fmt.Println(DoTheThing(false))

	/*
	for循环复用局部变量i，每一次放入匿名函数的应用都是同一个变量。 结果：

	0xc042046000 2
	0xc042046000 2
	 */
	//funs:=test2()
	//for _,f:=range funs{
	//	f()
	//}

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
	//list:=make([]int,0)//make方法传给变量的是类型变量
	//list = append(list, 1)
	//fmt.Println(list)
	//
	//
	//s1 := []int{1, 2, 3}
	//s2 := []int{4, 5}
	//s1 = append(s1, s2...)//append切片时不要忘了s2后面的三个...
	//fmt.Println(s1)
	//
	//
	//sn1 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: "qq"}
	//sn2 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: "qq"}
	//
	//if sn1 == sn2 {
	//	fmt.Println("sn1 == sn2")
	//}
	//
	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//
	//fmt.Println(&sm1,&sm2)
	//fmt.Println(sm1.m,sm2.m)

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

	//if reflect.DeepEqual(sm1, sm2) {
	//	fmt.Println("sm1 ==sm2")
	//}else {
	//	fmt.Println("sm1 !=sm2")
	//}
	//
	////考察interface内部结构
	//var x1 *int = nil
	//Foo(x1)
	//
	//fmt.Println(x,y,z,k,p)
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
		panic(value)
	}
	/*
	考点：select随机性
	解答：
	select会随机选择一个可用通道做收发操作。 所以代码是有可能触发异常，也有可能不会。 单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：

	select 中只要有一个case能return，则立刻执行。
	当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
	如果没有一个case能return则可以执行”default”块。
	 */
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



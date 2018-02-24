package baseLearn

import (
	"fmt"
	"math"
	"errors"
	"time"
	"sync/atomic"
	"sync"
	"math/rand"
	"sort"
	"encoding/json"
	"regexp"
	"bytes"
	"strconv"
	"net/url"
	"net"
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
	"os"
	"io"
	"bufio"
	"strings"
	"os/exec"
	"syscall"
	"os/signal"
	"container/list"
	"unsafe"
	"reflect"
)

var size = 10
var s = make([]string,size)
//常量
const cons = "const"
func SliceTest()  {
	s = append(s,"a")
	s2 := []string{"a","b","c","d","e","f","g","h","i"}
	n := copy(s,s2)
	fmt.Println(n)
	fmt.Println(s)
	for i := 0;i<len(s);i++{
		for j := 0;j<=i ;j++  {
			fmt.Print(s[j])
		}
		fmt.Println()
	}

	for i := 0;i<len(s);i++{
		for k :=0;k< len(s)-i-1; k++ {
			fmt.Print(s[k])
		}
		fmt.Println()
	}

	var map1 = make(map[int]string)
	map1[1] = "a"
	map1[2] = "b"

	map2 := map[int]string{0:"a",1:"b",2:"c"}
	mapTest(map2)

	var a = []string{"a","b","c"}
	rangTest(a)

}

func mapTest(a map[int]string) {
	for i := 0;i<len(a) ;i++  {
		fmt.Println(a[i])
		delete(a,i)
		fmt.Println(len(a))
	}

}

func rangTest(a []string) {
	for i,s := range a {
		fmt.Println(i,s)
	}
}

func plus(a int,b int) (int,int) {
	sum := a+b
	multy := a*b
	return sum ,multy

}

func add(a ...int)  {
	var sum int
	for _,v := range a {
		sum += v
	}
	fmt.Println(sum)
}

func closePakageTest() func() int {
//每次执行该函数i都会保存计算后的值
	var i = 0
	return func() int {
		i++
		return i
	}
}

//递归
func Fact(n int) int  {
	if n == 0{
		return 1
	}
	return n*Fact(n-1)
}

//结构体
type people struct {
	name string
	age int
}

func (p *people) SetName(name string) {
	p.name = name
}

func (p *people) SetAge(age int) {
	p.age = age
}

func (p *people) GetName(name string) string {
	return p.name
}

func (p *people) GetAge(age int) int {
	return p.age
}

func showStructField()  {
	p := &people{"zhangsan",23}
	p.age = 34
	name := p.GetName("zhangsan")
	fmt.Println(name)
}


//接口
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width,length float64
}

type circle struct {
	radius float64
}
//实现接口：struct实现接口中的所有方法即实现该接口
func (re *rect) area() float64 {
	return re.width * re.length
}

func (re *rect) perim() float64 {
	return 2*re.width + 2*re.length
}

func (c *circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func (c *circle) perim() float64 {
	return 2*math.Pi*c.radius
}

//异常处理
/*
Golang 有2个内置的函数 panic() 和 recover()，用以报告和捕获运行时发生的程序错误，
与 error 不同，panic-recover 一般用在函数内部。一定要注意不要滥用 panic-recover，
可能会导致性能问题，我一般只在未知输入和不可靠请求时使用。

golang 的错误处理流程：当一个函数在执行过程中出现了异常或遇到 panic()，正常语句就会立即终止，然后执行 defer 语句，再报告异常信息，最后退出 goroutine。如果在 defer 中使用了 recover() 函数,则会捕获错误信息，使该错误信息终止报告。
 */

func f1(flag int) (int,error) {
	if flag == 34 {
		return flag, errors.New("can't work with it")
	}
	flag = flag + 1
	return flag ,nil
}

type argError struct {
	arg int
	prob string
}
//实现error 接口
func (e *argError) Error() string {
	return fmt.Sprint("%d - %s",e.arg,e.prob)

}

func f2(arg int) (int,error) {
	if arg == 42 {
		return -1,&argError{arg,"can't work with it"}
	}
	return arg + 3, nil
}

func ErrorTest() {
	for _,i := range []int{7,42}{
		if r,e:=f1(i);e!=nil{
			fmt.Println("f1 failed:",e)
		} else {
			fmt.Println("f1 worked:",r)
		}
	}

	for _,i := range []int{7,42}{
		if r,e := f2(i);e!=nil {
			fmt.Println("f2 failed:",e)
		}else {
			fmt.Println("f2 worked:",r)
		}
	}

	_,e := f2(42)
	if ae,ok := e.(*argError);ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}

//协程:轻量级线程
func lightThread(from string)  {
	for i:=0;i<3 ;i++  {
		fmt.Println(from,":",i)
	}
}

func GoTest()  {
	go lightThread("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

//channel通道
var ch = make(chan int)

func ChannelTest()  {
	go func(i int) {
		ch <- 1
	}(3)
	i := <-ch
	fmt.Println("channel:",i)
}
//通道缓冲
var cha = make(chan string,2)

func ChannelTest2()  {
	cha <- "a"
	cha <- "b"
	for k := 0; k<len(cha); k++ {
		select {
		case i := <-cha:
			fmt.Println(i)
		case j:=<-cha:
			fmt.Println(j)
		}
	}


}
//通道同步
/**
我们可以使用通道来同步 Go 协程间的执行状态。这里是一个使用阻塞的接受方式来等待一个 Go 协程的运行结束。

这是一个我们将要在 Go 协程中运行的函数。done 通道将被用于通知其他 Go 协程这个函数已经工作完毕。发送一个值来通知我们已经完工啦。

运行一个 worker Go协程，并给予用于通知的通道。程序将在接收到通道中 worker 发出的通知前一直阻塞。

如果你把 <- done 这行代码从程序中移除，程序甚至会在 worker还没开始运行时就结束了

 */
func worker(done chan bool) {
	fmt.Print("working....")
	//睡眠1秒
	time.Sleep(time.Second)
	fmt.Println("done")
	//发送一个值来通知这里已经做完
	done <- true
}

func doWork()  {
	done := make(chan bool,1)
	go worker(done)
	<-done
	fmt.Println(<-done)
}

//通道方向
//接收通道
func ping(pings chan <-string,msg string)  {
	pings <- msg
}
//发送通道pings <-chan string

func pong(pings <-chan string,pongs chan <-string)  {
	msg := <-pings
	pongs <- msg
}
//双向通道var channel = make(chan string,1)

//通道选择器
func chose(){
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
//注意从第一次和第二次 Sleeps 并发执行，总共仅运行了两秒左右。
	for i:=1;i<2 ;i++  {
		select {
		case msg1 := <-c1:
			fmt.Println("received",msg1)
		case msg2 := <-c2:
			fmt.Println("received",msg2)
		}
	}
}
/**
超时 对于一个连接外部资源，或者其它一些需要花费执行时间的操作的程序而言是很重要的。
得益于通道和 select，在 Go中实现超时操作是简洁而优雅的。
 */
func timeOut()  {
	c1 := make(chan string,1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
/**
这里是使用 select 实现一个超时操作。res := <- c1 等待结果，<-Time.After 等待超时时间 1 秒后发送的值。由于 select 默认处理第一个已准备好的接收操作，如果这个操作超过了允许的 1 秒的话，将会执行超时 case。
Go 的 select 让你能够等待多个 channel 操作。通过 select 结合 goroutine 和 channel 是 Go 的重要特色。
 */
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string,1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <- c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
	
}
//Non-Blocking channel operations
//channel上简单的发送和接受是阻塞的。然而，我们可以使用select和default字句来实现非阻塞发送，接收甚至非阻塞的多路选择。

func noblockingChannel()  {
	messages := make(chan string)
	signals := make(chan bool)
	//这是一个非阻塞的接收。如果message的只可以获取，select将随即进入<-message字句否则将立刻进入default事件。
	select {
	case msg := <- messages:
		fmt.Println("recevied message",msg)
	default:
		fmt.Println("no message recived")
	}
	msg := "hi"
	//类似的有非阻塞发送
	select {
	case messages <- msg:
		fmt.Println("sent message",msg)
	default:
		fmt.Println("no message sent")
	}


	//我们可以在default上使用多个事件来实现多路非阻塞select。
	//这里我们试图在message和signal上均进行非阻塞接收
	select {
	case msg := <-messages:
		fmt.Println("recived message",msg)
	case sig := <-signals:
		fmt.Println("recived signal",sig)
	default:
		fmt.Println("no activity")
	}
}
//closing Channels

func closeChannel()  {
	jobs := make(chan int,5)
	done := make(chan bool)
	//下面是工人goroutine。他通过j，more：=<-jobs反复获取作业
	//在这个2返回值的接收中，如果作业关闭，所有值都已接收，more会变为false
	//我们用其在完成所有作业时进行已完成通知
	go func() {
		for {
			j,more:=<-jobs
			if more {
				fmt.Println("recived job",j)
			}else {
				fmt.Println("recived all jobs")
				done<-true
				return
			}
		}
	}()
	//此处向工人发送了3个作业，然后关闭它
	for j:=1;j<=3 ;j++  {
		jobs<-j
		fmt.Println("send job",j)
	}
	close(jobs)
	fmt.Println("send all jobs")
	<-done
}
//Range over channels
func rangeChannel()  {
	queue := make(chan string,2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem:=range queue{
		fmt.Println(elem)
	}
}

func switchCase() {
	i := 2
	fmt.Print("write",i,"as")
	switch i {
	case 1:fmt.Println("one")
	case 2:fmt.Println("two")
	case 3:fmt.Println("three")
	default:
		fmt.Println("zero")
	}

	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
	t := time.Now()
	switch {
	case t.Hour()<12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
}
//数组go中只有定长数组
var array1 = [...]string{"a","c","b","d"}
//不设定元素会默认设置成0
var array2 = [6]int{}
//切片（可以改变长度的数组）
var slice1 = make([]int,6,8)
var slice2 = make([]string,3)
//map
var map1 = make(map[int]string)
//追加、复制切片，用的是内置函数append和copy，copy函数返回的是最后所复制的元素的数量。
func arraySliceTest()  {
	slice3 := []string{}
	slice3[0] = "a"
	slice3 = append(slice3,"a")
	fmt.Println(slice3)
	copy(slice3,slice2)
	map1[1] = "a"
	map1[2] = "b"
	fmt.Println(map1[1])
	delete(map1,1)
	_,v := map1[2]
	fmt.Println(v)
	map2 := map[int]string{1:"a",2:"b"}
	fmt.Println(map2)
	for k,v:= range map1{
		fmt.Println(k,v)
	}
}
func sum(nums ...int)  {
	total := 0
	for _,num:=range nums{
		total += num
	}
	fmt.Println(total)
}

//使用channel进行跨goroutine同步执行
func worker2(done chan bool)  {
	fmt.Print("working.....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func doWork2()  {
	//启动一个worker goroutine,赋予它用以通知的channel
	done := make(chan bool,1)
	go worker(done)
	<-done

}

//timers
func timerTest(){
	//timer代表未来的一个单独事件。你要告诉它要等多久，它提供一个通道，在指定时间发出通知。下面这个timer将等待2秒钟
	timer := time.NewTimer(time.Second*2)
	//定时器通道由于操作<-timer.c发生阻塞，直到它发送一个值来表明定时器到时
	<-timer.C
	fmt.Println("timer 1 expired")
	//如果你仅仅想等待一段时间，可以用time.sleep,使用timer的一个原因是，你可以在计时结束前取消，如下例：
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

}

//timer用来在将来的某一次时间做某事一次。而ticker会在一个指定时间间隔重复做某事。这里是一个ticker的例子：它会在我们停止之前定期触发

func tickerTest()  {
	//ticker与timer的机制相似，都是发送值的通道
	//这里我们使用channel内置的range来遍历每500ms到达的值
	ticker := time.NewTicker(time.Millisecond*500)
	go func() {
		for t:=range ticker.C {
			fmt.Println("tick at",t)
		}
	}()
	time.Sleep(time.Millisecond*1600)
	ticker.Stop()
	fmt.Println("ticker stopped")
}

//速率限制是控制资源利用和维护服务质量的重要机制。Go 通过 goroutine，channel 和 ticker 可以优雅的支持速率控制
func rateLimit()  {
	//首先，我们看下基本的速率控制。
	//假设我们想要控制处理的输入请求，我们通过同一个通道来为这些请求提供服务
	requests := make(chan int,5)
	for i:=1;i<=5 ;i++  {
		requests<-i
	}
	close(requests)
	//limiter通道每过200毫秒接收一次数据，这是速率控制策略中的调节器
	limiter := time.Tick(time.Millisecond*200)
	//在服务每个请求之前，通过limiter通道阻塞接收，我们将自己限制在每200毫秒处理一个请求上
	for req := range requests {
		<-limiter
		fmt.Println("request",req,time.Now())
	}
	//我们可能希望在我们的速率限制方案中允许短时间的请求，同时保留整体速率限制。
	//我们可以通过缓冲限制器通道来实现这一点。
	//这个burstyLimiter通道将允许多达3个事件的突发。
	burstyLimiter := make(chan time.Time,3)
	//填充通道，来展示可允许的突发
	for i:=0;i<3 ;i++ {
		burstyLimiter <- time.Now()
	}
	//每隔200毫秒将试图添加一个新值到burstyLimiter，最多3个
	go func() {
		for t:=range time.Tick(time.Millisecond*200){
			burstyLimiter<-t
		}
	}()
	//模拟5个输入请求。前3个将受益于burstyLimiter的突发能力
	burstyRequests := make(chan int,5)
	for i:=1;i<=5 ;i++  {
		burstyRequests<-i
	}
	close(burstyRequests)
	for req:= range burstyRequests{
		<-burstyLimiter
		fmt.Println("request",req,time.Now())
	}
}

//原子计数器
func atomicCounter()  {
	var ops uint64 = 0
	for i:=0;i<50 ; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops,1)

				time.Sleep(time.Millisecond*200)
			}
		}()
	}
	time.Sleep(time.Second)

	//为了安全地使用计数器，当它仍被其他goroutine更新时，我们通过LoadUint64将当前值的副本提取到opsFinal中。
	//如上所述，我们需要给出这个函数来获取值的内存地址和操作。
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops",opsFinal)
}

//mutexes互斥

func mutexes() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var readOps uint64 = 0
	var writeOps uint64 = 0
	for r:=0;r<100 ;r++{
		go func() {
			total := 0
			for {
				key:=rand.Intn(5)//返回一个随机整数
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for w:=0;w<10;w++ {
		go func() {
			key:=rand.Intn(5)
			val:=rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			atomic.AddUint64(&writeOps,1)
			time.Sleep(time.Millisecond)
		}()
	}
	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)//用于对变量值进行原子增操作，并返回增加后的值
	fmt.Println("readOps:",readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writOps:",writeOpsFinal)
	mutex.Lock()
	fmt.Println("state:",state)
	mutex.Unlock()

}

//指针pointer
func zeroval(ival int)  {
	ival = 0
}

//zeroptr有一个*int类型的参数，代表它接收的是一个指针
func zeroptr(iptr *int)  {
	//*iptr解引用，从内存地址中获取存放的值
	//对解引用指针的赋值将改变指定地址上的值
	*iptr = 0
}
func PointerTest() {
	i := 1
	fmt.Println("initial:",i)
	zeroval(i)
	fmt.Println("zeroval:",i)
	//&语法将获得变量i的内存地址，也就是指向变量i的指针
	zeroptr(&i)
	fmt.Println("zeroptr:",i)
	//指针也可以被打印
	fmt.Println("pointer:",&i)
	//zeroval没有改变PointerTest函数中i的值，而zeroptr会，因为它拥有指向变量i的内存地址。
}
/*
Go语言的函数调用参数全部是传值的, 包括 slice/map/chan 在内所有类型, 没有传引用的说法.
什么叫传引用?
比如有以下代码:

var a Object
doSomething(a) // 修改a的值
print(a)
如果函数doSomething修改a的值, 然后print打印出来的也是修改后的值,
那么就可以认为doSomething是通过引用的方式使用了参数a.

*/
//使用goroutine和channel内置的同步功能锁定多个goroutine同步访问共享数据

type readOp struct {
	key int
	resp chan int
}
type writeOp struct {
	key int
	val int
	resp chan bool
}

func gochansync(){
	//记录执行次数
	var readOps uint64 = 0
	var writeOps uint64 = 0
	//reads和writes通道将被用于其他goroutine分别发送读写请求
	reads := make(chan *readOp)
	writes := make(chan *writeOp)
	/**
	这里是拥有状态值的goroutine，与之前一样是个map，但被私有化
	这个goroutine反复选择reads和writes通道，响应到达的请求。
	首先执行所有请求的操作然后在响应通道上发送值来表示唱功执行响应
	 */
	 go func() {
		 var state = make(map[int]int)
		 for  {
			 select {
			 case read:= <-reads:
				 read.resp <- state[read.key]
			 case write := <-writes:
				 write.resp <- true
			 }
		 }
	 }()
	 /**
	 启动100个goroutine，通过读取通道来读取有状态的groutine
	 每次读取需要构建一个readOp，通过reads发送给它再通过所提供的的resp通道获取结果
	  */
	for r := 0;r < 100 ; r++ {
		go func() {
			read := &readOp{
				key:rand.Intn(5),
				resp:make(chan int),
			}
			reads <- read
			<-read.resp
			atomic.AddUint64(&readOps,1)
			time.Sleep(time.Millisecond)
		}()
	}

	//启动10个写操作
	for w := 0;w < 10 ; w++ {
		go func() {
			for {
				write := &writeOp{
					key:rand.Intn(5),
					val:rand.Intn(100),
					resp:make(chan bool),
				}
				writes <- write
				<- write.resp
				atomic.AddUint64(&writeOps,1)
				time.Sleep(time.Second)
			}
		}()
	}
	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:",readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:",writeOpsFinal)
}
//sort包实现了内置和自定义类型的排序。首先看看内类型排序。

func sortTest()  {
	//sort改变了给定的slice，而不是返回一个新的
	strs := []string{"c","z","v"}
	sort.Strings(strs)
	fmt.Println("Strings:",strs)
	ints := []int{3,6,3}
	sort.Ints(ints)
	fmt.Println(ints)
	//可以使用sort检查一个slice是不是已经排好序了
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:",s)
}

//有时候我们想要对一个集合进行非自然顺序的排序。
//为了根据自定义函数排序，我们需要相应的类型
//这里我们创建了一个ByLength类型
type ByLength []string
//我们在ByLength上实现了sort接口的Len，less和Swap方法
//这里我们想要按照字符串长度曾序排列
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}
func (s ByLength) Less(i,j int) bool {
	return len(s[i]) < len(s[j])
}
func sortAction()  {
	fruits := []string{"peach","banana","Kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
//json
type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page int   //'json:"page"'
	Fruits []string  //'json:"fruits"'
}

func JsonTest(){
	bolB,_:=json.Marshal(true)
	fmt.Println(string(bolB))
	intB,_:=json.Marshal(1)
	fmt.Println(string(intB))
	fltB,_:=json.Marshal(2.34)
	fmt.Println(string(fltB))
	slcD:=[]string{"apple","peach","pear"}
	slcB,_:=json.Marshal(slcD)
	fmt.Println(slcB)
	mapD:=map[string]int{"apple":5,"lettuce":7}
	mapB,_:=json.Marshal(mapD)
	fmt.Println(mapB)

	res1D:=&Response1{
		Page:1,
		Fruits:[]string{"apple","peach","pear"},
	}
	res1B,_:=json.Marshal(res1D)
	fmt.Println(res1B)

	//byt := []byte('{"num":6.33,"strs":["a","b"]}')
	
	var dat map[string]interface{}
	byt := []byte{'a':1,'b':'e'}
	if err := json.Unmarshal(byt,&dat);err!=nil{
		panic(err)
	}
}
//regexp
func regexpTest()  {
	//测试模式是否符合字符串
	math,_ := regexp.MatchString("p([a-z]+)ch","peach")
	fmt.Printf("%s",math)
	r,_ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	//找到一个匹配
	fmt.Println(r.FindString("peach punch"))
	//寻找第一个匹配，但返回起止的索引而非字符串
	fmt.Println(r.FindStringIndex("peach punch"))
	//submatch包含整串匹配，页包含内部匹配
	fmt.Println(r.FindStringSubmatch("peach punch"))
	//返回整串匹配和内部匹配的索引信息
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	//这些all修饰的将返回输入中所有匹配的，不仅是第一个
	fmt.Println(r.FindAllString("peach punch pich",-1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch",-1))
	//第二个参数如果是非负数，则将限制最多匹配的个数
	fmt.Println(r.FindAllString("peach punch pinch",2))
	//提供[]byte参数，并将参数中的string去掉
	fmt.Println(r.Match([]byte("peach")))
	//一个纯compile不能用于常量，因为它有2个返回值。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	//regexp包也能用于使用其他值替换字符串的子集
	fmt.Println(r.ReplaceAllString("a peach","<fruit>"))
	//func修饰允许使用一个给定的函数修改匹配的字符串
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in,bytes.ToUpper)
	fmt.Println(string(out))
}
//time 获取时间
func getTeime()  {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)
	millis := nanos/1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	fmt.Println(time.Unix(secs,0))
	fmt.Println(time.Unix(0,nanos))
	//时间格式化
	p := fmt.Println
	//这里是一个根据RFC3339基本的格式化时间的例子，使用响应的布局常量
	t := time.Now()
	p(t.Format(time.RFC3339))
	//时间解析使用格式化相同的布局值
	t1,e := time.Parse(time.RFC3339, "2018-1-1T22:08:23+00:00")
	p(t1)
	p(e)
	//格式化和解析基于示例的布局
	//通常你是用常量在进行布局，但你也可以提供自定义的格式
	//但你必须使用Mon Jan 2 15;22:23 MST 2008来作为示例
	p(t.Format("3:03PM"))
	p(t.Format("Mon Jan _2 12:09:09 2006"))
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second())
	//解析返回一个错误来说明是什么问题
	ansic := "Mon Jan _2 12:09:09 2006"
	_,e2 := time.Parse(ansic,"8:43PM")
	p(e2)
}
//number parsing从字符串中解析数字是一个常见的任务
func parsNum()  {
	//64代表要解析的浮点数精度
 	f,_ := strconv.ParseFloat("1.234",64)
	fmt.Println(f)
	//0代表根据字符串 推断基数，64要求结果要适应64位
	i,_ := strconv.ParseInt("123",0,64)
	fmt.Println(i)
	//ParseInt可以识别十六进制
	d,_ := strconv.ParseInt("0x1c8",0,64)
	fmt.Println(d)
	u,_ := strconv.ParseUint("789",0,64)
	fmt.Println(u)
	//atoi是十进制整数解析的简便函数
	k,_ := strconv.Atoi("123")
	fmt.Println(k)
	//不合法的输入将导致解析函数返回一个错误
	_,err := strconv.Atoi("wqas")
	fmt.Println(err)
}
//url parsing
func parUrl()  {
	//这个URL解析示例，包含一个协议，授权信息，地址，端口，路径，查询参数以及查询拆分
	s := "postgres://user:pass@host.com:5423/path?k=v#f"
	//解析这个URL并保证没有错误
	u,err := url.Parse(s)
	if err!=nil {
		panic(err)
	}
	//可以直接访问协议
	fmt.Println(u.Scheme)
	//User包含所有授权信息，调用Username和Password可以得到单独的值
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p,_ := u.User.Password()
	fmt.Println(p)
	//Host包含地址和端口，使用SplitHostPort来抽取他们
	fmt.Println(u.Host)
	host,port,_ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	//为了以k=v的格式得到查询参数，使用RawQuery
	//也可以将查询参数解析到一个map中
	//解析的查询参数是从字符串到字符串的片段，故索引0可以只得到一个值
	fmt.Println(u.RawQuery)
	m,_ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
//SHA1 hashes
//SHA1 哈希经常用于计算二进制或者文本块的短表识。例如git版本控制系统使用SHA1来标示文本和目录。
//这里是Go如何计算SHA1哈希值
func GoHashTest()  {
	s := "sha1 this string"
	//产生一个哈希的模式是sha1.New(),sha1.Write(bytes)然后sha1.Sum([]bytes{})
	h := sha1.New()
	h.Write([]byte(s))
	//这里得到最终的哈希结果字节片段值，参数用于向已存在的字节片段追加，通常不需要
	bs := h.Sum(nil)
	//SHA1值经常用于打印成十六进制，如git提交时。使用%x格式参数来转换为十六进制
	fmt.Println(s)
	fmt.Printf("%x\n",bs)

}
//base64 encoding
func base64Test()  {
	data := "abc123!$*&()'-=@~"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	sDec,_ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec,_ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
//reading files 文件读写
func check(e error)  {
	if e!=nil {
		panic(e)
	}
}
func RFiles()  {
	//最基本的一个文件读取任务是将所有内容放入内存中
	dat,err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))
	//如果你想对文件的哪部分进行读取有更多的控制
	//首先你需要打开它
	f,err := os.Open("/tmp/dat")
	check(err)
	//从文件开头读取一些字节，允许到5，同时也是实际读取了的字节。
	b1 := make([]byte,5)
	n1,err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n",n1,string(b1))
	//也可以找到一个已知的位置并从哪里开始读取
	o2,err := f.Seek(6,0)
	check(err)
	b2 := make([]byte,2)
	n2,err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n",n2,o2,string(b2))
	//io包提供了一些函数，对于文件读取可能很有帮助
	o3,err := f.Seek(6,0)
	check(err)
	b3 := make([]byte,2)
	n3,err := io.ReadAtLeast(f,b3,2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n",n3,o3,string(b3))
	//没有内置的退回，但是seek（0,0）完成了这个事情
	_,err = f.Seek(0,0)
	check(err)
	//bufio包实现了一个带缓冲区的读取，它对于一些小的读取以及由于它所提供的额外方法很有帮助
	r4 := bufio.NewReader(f)
	b4,err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n",string(b4))
	//在完成时关闭文件（通常会在打开时通过defer计划执行）
	f.Close()
}

func WFiles()  {
	//写文件writing files
	d1 := []byte("hello\ngo\n")
	//perm := os.FileMode.Perm
	err := ioutil.WriteFile("tmp/dat1",d1,0644)
	check(err)
	//创建一个文件
	f,err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()
	d2 := []byte{115,111,123,122,12}
	n2,err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n",n2)
	n3,err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n",n3)
	f.Sync()
	w := bufio.NewWriter(f)
	n4,err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n",n4)
	w.Flush()
}
//过滤器 line filters
func filtersTest()  {
	//使用一个带缓冲的scanner可以方便的上使用scan方法来直接读取一行
	//每次调用该方法可以让scanner读取下一行
	scanner := bufio.NewScanner(os.Stdin)
	//text方法返回当前的token，现在是输入下一行
	for scanner.Scan(){
		ucl := strings.ToUpper(scanner.Text())
		//输出大写的行
		fmt.Println(ucl)
	}
	//检查scanner的错误，文件结束符不会当做是一个错误
	if err := scanner.Err();err!=nil{
		fmt.Fprintln(os.Stderr,"error:",err)
		os.Exit(1)
	}
}
//读取command-line arguments
func argsTest()  {
	//os.args提供原始命令行参数访问功能
	//切片的第一个值是程序的路径
	argsWithProg := os.Args
	//os.Args[1:]保存程序的所有参数
	argsWithoutProg := os.Args[1:]
	//可以通过自然索引获取到每个单独的参数
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
//Enviroment Variables环境变量
func envVar(){
	//使用os,Setenv来设置一个键值对
	//使用os.Getenv来获取一个环境变量，如果不存在，返回空字符串
	os.Setenv("FOO","1")
	fmt.Println("FOO:",os.Getenv("FOO"))
	fmt.Println("BAR",os.Getenv("BAR"))
	fmt.Println()
	//使用os.Eniron来列出所有环境变量键值对
	for _,v := range os.Environ(){
		pair := strings.Split(v,"=")
		fmt.Println(pair[0])
	}
}
//spawning processes

func SpawProcess()  {
	//exec.Command函数帮助我们创建一个表示这个外部进程的对象
	dataCmd := exec.Command("data")
	//output 等待命令运行完成，并收集命令的输出
	dataOut,err := dataCmd.Output()
	if err != nil{
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dataOut))
	grepCmd := exec.Command("grep","hello")
	//获取输入输出管道
	grepIn,_ := grepCmd.StdinPipe()
	grepOut,_ := grepCmd.StdoutPipe()
	//运行进程，写入输入信息，读取输出结果，等待程序运行结束
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepByte,_ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println(string(grepByte))
	//通过bash命令的-c选项来执行一个字符串包含的完整命令
	IsCmd := exec.Command("bash","-c","ls -a -l -h")
	IsOut,err := IsCmd.Output()
	if err != nil{
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(IsOut))
}

//exec'ing processes
func execProcessing()  {
	//通过LookPath得到需要执行的可执行文件的绝对路径
	binary,lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	//Exec 需要的参数是切片形式的，第一个参数为执行程序名
	args := []string{"ls","-a","-l","-h"}
	env := os.Environ()
	execErr := syscall.Exec(binary,args,env)
	if execErr != nil {
		panic(execErr)
	}

}
//signals
func signalTest(){
	//go通过向一个通道发送os.Signal值来进行信号通知
	sigs := make(chan os.Signal,1)
	//同时创建一个用于在程序可以结束时进行通知的通道
	done := make(chan bool,1)
	//注册给定通道用于接收特定信号
	signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)
	//go协程执行一个阻塞的信息号接收操作，当它得到一个值时，打印并通知程序可以退出
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("waiting signal")
	<- done
	fmt.Println("exiting")
	//运行，使用ctrl-c发送信号
}
//程序退出Exit
func exitTest()  {
	//当时用os.Exit时，defer将不会执行
	defer fmt.Println("!")
	//退出，并且状态为3
	os.Exit(3)
}

func ListTest()  {
	list1 := list.New()
	for i := 0;i < 5; i++{
		list1.PushBack(i)
	}
	//取值
	for j := list1.Front();j!=nil ;j = j.Next()  {
		fmt.Println(j.Value)
	}
	//取出首部元素的值
	fmt.Println(list1.Front().Value)
	//取出尾部元素的值
	fmt.Println(list1.Back().Value)
	//在首部元素后插入元素值为3的元素
	element := list1.InsertAfter(3,list1.Front())
	fmt.Println(element.Value)
	fmt.Println(list1.Len())
}

type User struct {
	Name string
	Age string
	ID string
}

func JsonTest2()  {
	user := User{
		Name:"张三",
		Age:"23",
		ID:"23456",
	}
	json1,_ := json.Marshal(user)
	fmt.Println(string(json1))
	fmt.Printf("%s\n",json1)
	json2,_ := json.MarshalIndent(user,"","")
	fmt.Printf("%s\n",json2)
	fmt.Println(string(json2))
	user2 := User{
		Name:"李四",
		Age:"22",
		ID:"1234",
	}
	user3 := User{
		Name:"王五",
		Age:"33",
		ID:"678345",
	}
	users := []User{user,user2,user3}
	json3,_ := json.Marshal(users)
	fmt.Println("==============")
	fmt.Println(string(json3))

	var user4 User
	json.Unmarshal(json1,&user4)
	fmt.Println("--------------------")
	fmt.Println(user4.Name)

	fmt.Println(os.Stdin.Name(),"==========")

}

func EncodeAndDecod()  {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	var v map[string]interface{}
	//控制台输入格式{"BPM":1234}
	fmt.Println("json decode")
	dec.Decode(&v)
	fmt.Println(v)
	for k := range v{
		fmt.Println(k)
		if k!="Name" {
			delete(v,k)
		}
	}
	fmt.Println("json encode")
	enc.Encode(&v)
}
//函数恐慌和恢复panic 和 recover

//判断函数是否会产生恐慌
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x:=recover();x!=nil {
			b = true
		}
	}()
	f()
	return
}


func DeferTest()  {
	defer func() {
		fmt.Println("a")
	}()
	defer func() {
		fmt.Println("b")
	}()
	defer func() {
		fmt.Println("c")
	}()
	defer func() {
		fmt.Println("d")
	}()
}

func UnsafePointerTest()  {
	var a = [4]int{1,2,3,4}
	ptr := unsafe.Pointer(&a[0])
	fmt.Println(ptr)
	//将指针转成int类型整数
	fmt.Println(uintptr(ptr))
	ptr = unsafe.Pointer(uintptr(ptr))
	fmt.Println(ptr)

	//从slice中得到一块内存地址是很容易的：
	s := make([]byte, 200)
	ptr = unsafe.Pointer(&s[0])
	fmt.Println(ptr)
	/*//从一个内存指针构造出Go语言的slice结构相对麻烦一些，比如其中一种方式：
	var ptr2 unsafe.Pointer
	//将ptr2强转成*[1<<10]byte类型的指针再取前200个元素组成slice
	s2 := ((*[1<<10]byte)(ptr2))[:200]
	fmt.Println(s2)*/

	//第二种从一个内存指针构造出Go语言的slice结构方法
	/*var ptr3 unsafe.Pointer
	var s1 = struct {
		addr uintptr
		len int
		cap int
	}{ptr3, length, length}
	s := *(*[]byte)(unsafe.Pointer(&s1))*/

	//使用reflect.SliceHeader的方式来构造slice，比较推荐这种做法：
	var ptr3 unsafe.Pointer
	var length int = 200
	var o []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(ptr3)
	fmt.Println(sliceHeader)
	fmt.Println("======================")

}



















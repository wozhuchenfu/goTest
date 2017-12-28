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
)

var size = 10
var s = make([]string,size)
//常量
const cons = "const"
func SliceTest()  {
	s = append(s,"a")
	s2 := []string{"a","b","c","d","e","f","g","h","i"}
	copy(s,s2)
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
	for _,i:=range []int{7,42}{
		if r,e:=f1(i);e!=nil{
			fmt.Println("f1 failed:",e)
		} else {
			fmt.Println("f1 worked:",r)
		}
	}

	for _,i:=range []int{7,42}{
		if r,e:=f2(i);e!=nil {
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

func goTest()  {
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

func channelTest()  {
	go func(i int) {
		ch <- 1
	}(3)
	i := <-ch
	fmt.Println(<- ch,i)
}
//通道缓冲
var cha = make(chan string,2)

func channelTest2()  {
	cha <- "a"
	cha <- "b"
	fmt.Println(<-cha)
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



















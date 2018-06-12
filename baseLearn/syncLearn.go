package baseLearn

import (
	"sync"
	"fmt"
	"time"
)

var (
	WG sync.WaitGroup
 	syncmux sync.Mutex
 	syncCond sync.Cond //条件变量条件等待通过 Wait 让协程等待，通过 Signal 让一个等待的协程继续，通过 Broadcast 让所有等待的协程继续。
	//在 Wait 之前应当手动为 c.L 上锁，Wait 结束后手动解锁。为避免虚假唤醒，需要将 Wait 放到一个条件判断循环中
 	syncMap sync.Map
 	syncPool sync.Pool //线程安全对象池
 	syncOnce sync.Once //只执行一次以后不再触发
 	syncRW sync.RWMutex //读写锁
 	)

var (
	synccha = make(chan int,6)
	syncchaInterface = make(chan SyncPerson)
)
func rountine1(i int)  {

	fmt.Println("hello world",i)
	synccha <- i
	WG.Done()

}

func SyncPoolTest()  {
	syncPool.Put(1)
}

type SyncPerson struct {
	Name string
	Age int
}

var syncinstance SyncPerson
func creatPerson(){
	syncinstance = *new(SyncPerson)
}


func SyncOnceTest() SyncPerson {
	syncOnce.Do(creatPerson)
	return syncinstance
}

func SyncMapTest()  {
	syncMap.Store(1,"123")
	v,ok := syncMap.Load(1)
	if ok {
		fmt.Println(v)
	}
	fmt.Println(syncMap.Load(1))
	fmt.Println(syncMap.LoadOrStore(2,"234"))
	fmt.Println(syncMap.Load(2))
}


func WGTest()  {
	for i := 0;i<6 ;i++  {
		WG.Add(1)
		go func(i int) {
			syncmux.Lock()
			rountine1(i)
			syncmux.Unlock()
		}(i)
	}
	/*for k := 0;k<6 ;k++  {
		 func(i int) {
			rountine1(i)
		}(k)
	}*/
	/*for j:=1;j<5 ;j++  {
		WG.Done()
	}*/

	//go rountine1(1)
	defer close(synccha)

	fmt.Println("syncTest")
	WG.Wait()
	for n:=0;n<len(synccha) ;n++  {
		fmt.Println(<-synccha)
	}
	/*for v := range synccha{
		fmt.Println(v)
	}*/
	fmt.Println("syncTest")
	fmt.Println("synccha'length",len(synccha))
	/*select {
	case  re,ok := <- synccha:
		if ok {
			fmt.Println(re)
		}
	}*/

}

var (
	synclock = new(sync.Mutex)
	synccond = sync.NewCond(synclock)
)

func rountine2(s int)  {
	synccond.L.Lock()
	synccond.Wait()
	fmt.Println("syncCond=====",s)
	synccond.L.Unlock()
}

func SyncCondTest()  {

	for i := 0;i<6 ;i++  {
		go rountine2(i)
		fmt.Println("==============")
	}

	synccond.Signal()
	synccond.Broadcast()
	time.Sleep(time.Second*3)
	//synccond.L.Unlock()
}





package baseLearn

import (
	"runtime"
	"fmt"
)

func TestMoreCPUS() {
	cpuNums := runtime.NumCPU()//cpu的核的数量
	runtime.GOMAXPROCS(cpuNums)//设置cpu的核的数量，从而实现高并发
	c := make(chan bool)

	for i := 0; i < 10; i++ {
		go testCPU(c, i)
	}
	<-c
	fmt.Println("main ok")
}

func testCPU(c chan bool, n int)  {
	x := 0
	for i :=0;i<100000;i++{
		x += i
	}
	print(c,x)
	if n == 9 {
		c <- true
	}
}
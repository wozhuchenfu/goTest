package baseLearn

import (
	"fmt"
	"time"
)
//本例中我们将看到如何使用 goroutine 和 channel 来实现一个工人池

func worker1(id int,jobs <-chan int,results chan <- int)  {
	for j:=range jobs{
		fmt.Println("worker",id,"started job",j)
		time.Sleep(time.Second)
		fmt.Println("worker",id,"finished job",j)
		results <- j*2
	}
}

func WorkerPool()  {
	jobs := make(chan int,100)
	results := make(chan int,100)

	for w:=1;w<=3;w++{
		go worker1(w,jobs,results)
	}
	for j:=1;j<=5 ;j++  {
		jobs<-j
	}
	close(jobs)
	for a:=1;a<=5;a++ {
		<-results
	}
}
/**
运行的项目展示了有 5 个作业得以被不同的工人执行。尽管总共有 5 秒钟的时间，这个程序只需要 2 秒钟，因为有 3 名工作人员同时进行操作。
 */
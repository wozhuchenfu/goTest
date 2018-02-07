package NSQ

import (
	"github.com/bitly/go-nsq"
	"fmt"
	"bufio"
	"os"
	"time"
)

/*
基于go-nsq的客户端实现
几个值得注意的地方
Producer断线后不会重连，需要自己手动重连，Consumer断线后会自动重连
Consumer的重连时间配置项有两个功能(这个设计必须吐槽一下，分开配置更好一点)

Consumer检测到与nsqd的连接断开后，每隔x秒向nsqd请求重连
Consumer每隔x秒，向nsqlookud进行http轮询，用来更新自己的nsqd地址目录
Consumer的重连时间默认是60s(...菜都凉了)，我改成了1s
Consumer可以同时接收不同nsqd node的同名topic数据，为了避免混淆，就必须在客户端进行处理
在AddConurrentHandlers和 AddHandler中设置的接口回调是在另外的goroutine中执行的
Producer不能发布(Publish)空message，否则会导致panic
 */
var producer *nsq.Producer
func Producer()  {
	strIP1 := "127.0.0.1:4150"
	strIP2 := "127.0.0.1:4152"
	InitProducer(strIP1)
	running := true
	//读取控制台输入
	reader := bufio.NewReader(os.Stdin)
	for running{
		data,_,_:=reader.ReadLine()
		command := string(data)
		fmt.Println(command)
		if command == "stop" {
			running = false
			fmt.Println(running)
		}
		for err:= Publish("test",command);err!=nil;err=Publish("test",command) {
			//切换IP重连
			strIP1,strIP2 = strIP2,strIP1
			InitProducer(strIP1)
		}

	}

}

//初始化生产者
func InitProducer(str string)  {
	var err error
	fmt.Println("address:",str)
	producer,err = nsq.NewProducer(str,nsq.NewConfig())
	if err!=nil {
		panic(err)
	}
	defer func() {
		fmt.Println(recover())
	}()
}

//发布消息
func Publish(topic string,message string) error {
	var err error
	if producer != nil{
		if message == "" {//不能发送空串，否则会导致error
			return nil
		}
		err = producer.Publish(topic,[]byte(message))//发布消息
		return nil
	}
	return fmt.Errorf("producer is nil",err)
}


//消息接收者测试

//消费者
type ConsumerT struct {

}

func ConsumerTest() {
	InitConsumer("test","test-channel","127.0.0.1:4150")
	fmt.Println("Consumer")
	for {
		time.Sleep(time.Second*10)
	}

}
//处理消息
func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive",msg.NSQDAddress,"message:",string(msg.Body))
	return nil
}
//初始化消费者
func InitConsumer(topic string,channel string,address string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second//设置重连时间
	c,err := nsq.NewConsumer(topic,channel,cfg)//新建一个消费者
	if err != nil{
		panic(err)
	}
	c.SetLogger(nil,0)//屏蔽系统日志
	c.AddHandler(&ConsumerT{})//添加消费者接口
	//建立NSQLookupd链接
	if err:=c.ConnectToNSQLookupd(address);err!=nil {
		panic(err)
	}
	//建立多个nsq链接
	/*if err:=c.ConnectToNSQDs([]string{"127.0.0.1:4150","127.0.0.1:4152"});err!=nil{
		panic(err)
	}*/
	//建立一个nsq链接
	if err := c.ConnectToNSQD("127.0.0.1:4150");err!=nil {
		panic(err)
	}
}





























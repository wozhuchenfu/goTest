package database

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

type NSQHandler struct {

}

func (this *NSQHandler) HandleMessage(message *nsq.Message) error  {
	log.Println("recv:",string(message.Body))
	return nil
}

func testNSQ()  {
	waiter := sync.WaitGroup{}
	waiter.Add(1)
	go func() {
		defer waiter.Done()
		consumer,err := nsq.NewConsumer("test","ch1",nsq.NewConfig())
		if err != nil {
			log.Println(err)
			return
		}
		consumer.AddHandler(&NSQHandler{})
		err = consumer.ConnectToNSQD("127.0.0.1:6379")
		if err != nil {
			log.Println(err)
			return
		}
		select {

		}
	}()
	waiter.Wait()
}

type MsgQueue struct {
	addr string
	producer *nsq.Producer
}
func (this *MsgQueue) Init(addr string) error{
	var err error
	this.addr = addr
	cfg := nsq.NewConfig()
	this.producer,err = nsq.NewProducer(addr,cfg)
	if err!=nil {
		return err
	}
	err = this.producer.Ping()
	if err!=nil {
		this.producer.Stop()
		this.producer = nil
		return err
	}
	return nil
}





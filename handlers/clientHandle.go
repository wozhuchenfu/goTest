package handlers

import (
	"net"
	"fmt"
	"os"
	"time"
)

//客户端发送封包

func sender9(conn net.Conn)  {
	for i := 0; i < 100; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write(Packet([]byte(words)))
	}
	fmt.Println("send over")
}

func sendAction()  {
	server := "127.0.0.1:9988"
	tcpAddr,err := net.ResolveTCPAddr("tcp4",server)
	if err != nil{
		fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
		os.Exit(1)
	}
	conn,err := net.DialTCP("tcp",nil,tcpAddr)
	if err != nil{
		 fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
		 os.Exit(1)
	}
	defer conn.Close()
	go sender9(conn)
	for  {
		time.Sleep(1*time.Second)
	}
}




















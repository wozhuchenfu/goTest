package handlers

import (
	"net"
	"fmt"
	"time"
	"os"
)

//粘包问题演示服务端

func Server()  {
	netListen,err := net.Listen("tcp",":9988")
	Checkerror(err)

	defer netListen.Close()

	Log("Waiting for clients")

	for {
		conn,err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(),"tcp Connect Success")
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {

	buffer := make([]byte,1024)
	for {
		n,err := conn.Read(buffer)
		if err != nil{
			Log(conn.RemoteAddr(),"tcp Connect error:",err)
			return
		}
		Log(conn.RemoteAddr().String(),"receive data length:",n)
		Log(conn.RemoteAddr().String(),"receive data:",buffer[:n])
		Log(conn.RemoteAddr().String(),"receive data strign:",string(buffer[:n]))
	}
}





func Checkerror(err error)  {
	if err != nil {
		fmt.Println(os.Stderr,"Fatal error:%s",err.Error())
		os.Exit(1)
	}
}

func Log(v ...interface{})  {
	fmt.Println(v...)

}

//粘包问题演示客户端
/**
1、客户端发送一次就断开连接，需要发送数据的时候再次连接，典型如http。下面用golang演示一下这个过程，确实不会出现粘包问题。

//客户端代码，演示了发送一次数据就断开连接的
 */
func sender(conn net.Conn)  {
	for i := 0;i < 100; i++{
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write([]byte(words))
	}
}

func clientSend()  {
	server := "127.0.0.1:9988"
	tcpAddr,err := net.ResolveTCPAddr("tcp4",server)
	if err != nil{
		fmt.Println(os.Stderr,"Fatal error:%s",err.Error())
		os.Exit(1)
	}
	conn,err := net.DialTCP("tcp",nil,tcpAddr)
	if err != nil {
		fmt.Println(os.Stderr,"Fatal error:%s",err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	fmt.Println("connect success")
	go sender(conn)
	for  {
		time.Sleep(5 * time.Second)
	}
}













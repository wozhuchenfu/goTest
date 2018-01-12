package handlers

import (
	"fmt"
	"os"
	"net"
)

//服务端解包过程
func ServerHandle()  {
	netListn,err := net.Listen("tcp",":9988")
	Checkerror(err)

	defer netListn.Close()

	Log("Waiting for clients")
	for  {
		conn,err := netListn.Accept()
		if err != nil{
			continue
		}

		Log(conn.RemoteAddr().String(),"tcp connect success")
		go handleConnection9(conn)
	}
}

func handleConnection9(conn net.Conn)  {
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpbuffer := make([]byte,0)
	//声明一个管道用于接收解包的数据
	readerChannel := make(chan []byte,16)
	go reader(readerChannel)
	buffer := make([]byte,1024)
	for  {
		n,err := conn.Read(buffer)
		if err != nil{
			Log(conn.RemoteAddr().String(),"connection error:",err)
			return
		}
		tmpbuffer = Unpack(append(tmpbuffer,buffer[:n]...),readerChannel)
	}
}

func reader(readerChannel chan []byte)  {
	for  {
		select {
		case data := <- readerChannel:
			Log9(string(data))

		}
	}
}

func Log9(v ...interface{}){
	fmt.Println(v...)
}

func CheckError(err error){
	if err != nil{
		fmt.Println(os.Stderr,"Fatal error:%s",err.Error())
		os.Exit(1)
	}
}

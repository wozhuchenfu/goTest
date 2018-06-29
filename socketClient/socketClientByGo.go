package socketClient

import (
	"net"
	"fmt"
)

/*
TCP连接的建立
众所周知，TCP Socket的连接的建立需要经历客户端和服务端的三次握手的过程。连接建立过程中，
服务端是一个标准的Listen + Accept的结构(可参考上面的代码)，而在客户端Go语言使用net.Dial或DialTimeout进行
连接建立：
 */
func SocketClient()  {
	conn,err := net.Dial("tcp",":8080")
	//conn,err := net.DialTimeout("tcp",":8080",time.Second*2) //带超时的链接
	if err!=nil {
		fmt.Println(err)
	}

	conn.Write([]byte("hello"))
	buff := make([]byte,1024)
	conn.Read(buff)
	defer conn.Close()

}
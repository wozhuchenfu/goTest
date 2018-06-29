package socketServer

import (
	"net"
	"fmt"
	"time"
)

func handleConn(c net.Conn)  {
	defer c.Close()

	tcpConn,ok := c.(*net.TCPConn)
	if ok{
		tcpConn.SetReadBuffer(1024)
		tcpConn.SetWriteBuffer(1024)
		tcpConn.SetKeepAlive(true)
		tcpConn.SetNoDelay(false)
		tcpConn.SetKeepAlivePeriod(time.Second*5)
	}


	for {
		//read from the connect
		buffer := make([]byte,1024)
		_,err := c.Read(buffer)
		if err!=nil {
			fmt.Println(err)
		}
		fmt.Println(string(buffer))

		//write to the connect
		if len(buffer) >0 {
			c.Write(buffer)
		} else {
			c.Write([]byte("no data recived"))
		}

	}

}

func SocketTest()  {
	l,err := net.Listen("tcp","127.0.0.1:7878")
	if err!=nil {
		fmt.Println(err)
		return
	}
	for {
		conn,err := l.Accept()
		if err!=nil {
			fmt.Println(err)
			break
		}
		//start a new groutine to handle the new connection
		go handleConn(conn)
	}

}
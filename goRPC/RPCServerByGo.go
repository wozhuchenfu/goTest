package goRPC

import (
	"net/rpc"
	"net/http"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

/*
go对RPC的支持，支持三个级别：TCP、HTTP、JSONRPC
go的RPC只支持GO开发的服务器与客户端之间的交互，因为采用了gob编码
 */

//注意字段必须是导出
type ParamsServer struct {
	Width,Height int
}

type Rect struct{}

//函数必须是导出的
//必须有两个导出类型参数
//第一个参数是接收参数
//第二个参数是返回给客户端参数，必须是指针类型
//函数还要有一个返回值error
func (r *Rect) Area(p ParamsServer, ret *int) error {
	*ret = p.Width * p.Height;
	return nil;
}

func (r *Rect) Perimeter(p ParamsServer, ret *int) error {
	*ret = (p.Width + p.Height) * 2;
	return nil;
}
//基于http的RPC
func RPCBaseHTTPByGoTest() {
	rect := new(Rect);
	//注册一个rect服务
	rpc.Register(rect);
	//把服务处理绑定到http协议上
	rpc.HandleHTTP();
	err := http.ListenAndServe(":8080", nil);
	if err != nil {
		log.Fatal(err);
	}
}

//基于tcp的RPC
func chkError(err error) {
	if err != nil {
		log.Fatal(err);
	}
}
func RPCBaseTCPByGoTest()  {
	rect := new(Rect);
	//注册rpc服务
	rpc.Register(rect);
	//获取tcpaddr
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080");
	chkError(err);
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr);
	chkError(err2);
	//死循环处理连接请求
	for {
		conn, err3 := tcplisten.Accept();
		if err3 != nil {
			continue;
		}
		//使用goroutine单独处理rpc连接请求
		go rpc.ServeConn(conn);
	}
}

func RPCBasedJSONServerByGoTest()  {
	rect := new(Rect);
	//注册rpc服务
	rpc.Register(rect);
	//获取tcpaddr
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080");
	chkError(err);
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr);
	chkError(err2);
	for {
		conn, err3 := tcplisten.Accept();
		if err3 != nil {
			continue;
		}
		//使用goroutine单独处理rpc连接请求
		//这里使用jsonrpc进行处理
		go jsonrpc.ServeConn(conn);
	}
}
package goRPC

import (
	"net/rpc"
	"log"
	"fmt"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int;
}



func RPCBasedHTTPClientByGoTest() {
	//连接远程rpc服务
	rpc, err := rpc.DialHTTP("tcp", "127.0.0.1:8080");
	if err != nil {
		log.Fatal(err);
	}
	ret := 0;
	//调用远程方法
	//注意第三个参数是指针类型
	err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret);
	if err2 != nil {
		log.Fatal(err2);
	}
	fmt.Println(ret);
	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret);
	if err3 != nil {
		log.Fatal(err3);
	}
	fmt.Println(ret);
}

func RPCBasedTCPClientByGoTest()  {
	//连接远程rpc服务
	//这里使用Dial，http方式使用DialHTTP，其他代码都一样
	rpc, err := rpc.Dial("tcp", "127.0.0.1:8080");
	if err != nil {
		log.Fatal(err);
	}
	ret := 0;
	//调用远程方法
	//注意第三个参数是指针类型
	err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret);
	if err2 != nil {
		log.Fatal(err2);
	}
	fmt.Println(ret);
	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret);
	if err3 != nil {
		log.Fatal(err3);
	}
	fmt.Println(ret);
}

func RPCBasedJSONClientByGoTest()  {
	//连接远程rpc服务
	//这里使用jsonrpc.Dial
	rpc, err := jsonrpc.Dial("tcp", "127.0.0.1:8080");
	if err != nil {
		log.Fatal(err);
	}
	ret := 0;
	//调用远程方法
	//注意第三个参数是指针类型
	err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret);
	if err2 != nil {
		log.Fatal(err2);
	}
	fmt.Println(ret);
	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret);
	if err3 != nil {
		log.Fatal(err3);
	}
	fmt.Println(ret);
}
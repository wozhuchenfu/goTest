package handlers

import (
	"encoding/binary"
	"bytes"
	"fmt"
)
//一个网络包包头的定义和初始化
type Head struct {
	Cmd byte
	Version byte
	Magic uint16
	Reserve byte
	HeadLen byte
	BodyLen uint16
}
//这个是一个常见的在tcp 拼包得例子。在例子中通过binary.BigEndian.Uint16将数据按照网络序的格式读出来，放入到head中对应的结构里面。
func NewHead(buf []byte) *Head {
	head := new (Head)

	head.Cmd 		= buf[0]
	head.Version 	= buf[1]
	head.Magic 		= binary.BigEndian.Uint16(buf[2:4])
	head.Reserve 	= buf[4]
	head.HeadLen 	= buf[5]
	head.BodyLen 	= binary.BigEndian.Uint16(buf[6:8])
	return head
}

//将结构体序列化到一个buf中
//在序列化结构对象时，需要注意的是，被序列化的结构的大小必须是已知的，可以通过Size接口来获得该结构的大小，从而决定buffer的大小。
//固定大小的结构体，就要求结构体中不能出现[]byte这样的切片成员，否则Size返回-1，且不能进行正常的序列化操作。
//通过Size可以得到所需buffer的大小。通过Write可以将对象a的内容序列化到buffer中。这里采用了小端序的方式进行序列化（x86架构都是小端序，网络字节序是大端序）。
//对于结构体中得“_”成员不进行序列化
type A struct {
	One int32
	Two int32
}

var a A

func Reads() {
	a.One = int32(1)
	a.Two = int32(2)
	buf := new (bytes.Buffer)
	fmt.Println("a's size is",binary.Size(a))
	binary.Write(buf,binary.LittleEndian,a)
	fmt.Println("after write,buf is:",buf.Bytes())
	Writes()
}


//从buf中反序列化回一个结构
var aa A

func Writes()  {
	buf := new (bytes.Buffer)
	binary.Write(buf,binary.LittleEndian,a)
	fmt.Println(buf.Bytes(),buf.Len())
	binary.Read(buf,binary.LittleEndian,&aa)
	fmt.Println("after aa is",aa)

}

























package handlers

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader   =  "www.01happy.com"
	ConstHeaderLength  = 15
	ConstSeveDataLength  = 4

)

//封包
func Packet(message []byte) []byte {
	return append(append([]byte(ConstHeader),IntToBytes(len(message))...),message...)

}
//解包
func Unpack(buffer []byte,readerChannel chan []byte) []byte {
	length := len(buffer)
	var i int
	for i = 0;i < length ; i++ {
		if length < i + ConstHeaderLength + ConstSeveDataLength{
			break
		}
		if string(buffer[i : i+ConstHeaderLength]) == ConstHeader{
			messageLength := BytesToInt(buffer[i + ConstHeaderLength : i +ConstHeaderLength+ConstSeveDataLength])
			if length < i+ConstHeaderLength+ConstSeveDataLength+messageLength{
				break
			}
			data := buffer[i+ConstHeaderLength+ConstSeveDataLength:i+ConstHeaderLength+ConstSeveDataLength+messageLength]
			readerChannel <- data
			i += ConstHeaderLength + ConstSeveDataLength + messageLength -1
		}
	}
	if i == length {
		return make([]byte,0)
	}
	return buffer[i:]
}

func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer,binary.BigEndian,x)
	return bytesBuffer.Bytes()
}

//字节转换成整型
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer,binary.BigEndian,&x)
	return int(x)
}






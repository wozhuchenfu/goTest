package io

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
)

func IOTest()  {
	f,err:=os.Open("C:\\Users\\edianzu\\Desktop\\测试账号.txt")
	defer f.Close()
	//以读写方式打开文件，如果不存在，则创建
	f2, _ := os.OpenFile("C:\\Users\\edianzu\\Desktop\\goFile.txt", os.O_RDWR|os.O_CREATE, 0766)
	defer f2.Close()
	if err!=nil {
		log.Fatal(err)
		fmt.Println(err.Error())
		panic(err)
	}
	read := bufio.NewReader(f)
	write := bufio.NewWriter(f2)
	/*n,err3 := read.WriteTo(write)
	fmt.Println(n)
	if err3!=nil {
		fmt.Println(err3.Error())
	}*/
	for {
		l,b,err2 := read.ReadLine()

		if err2!=nil && err2!=io.EOF{
			log.Fatal(err2)
			fmt.Println("err2:",err2.Error())
		}
		//enc := mahonia.NewEncoder("GBK")
		dec := mahonia.NewDecoder("gbk")
		fmt.Println(dec.ConvertString(string(l)))
		//直接写入文件
		//f2.WriteString(dec.ConvertString(string(l))+"\r\n")
		write.WriteString(dec.ConvertString(string(l))+"\r\n")
		//刷新写入文件
		write.Flush()
		if err2 == io.EOF {
			fmt.Println(b)
			break
		}
	}


}
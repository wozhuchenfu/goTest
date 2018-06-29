package sshByGo

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"io"
	"bufio"
	"encoding/csv"
	"strings"
	"container/list"
)

var (
	num int
)

func SSHByGo2() {
	if len(os.Args) == 1{
		fmt.Println("请输入文件名参数")
		return
	}
	list := listNode("添加文件路径")
	fmt.Println("请选择执行的语句")
	fmt.Scanln(&num)
	if num <= list.Len(){
		fmt.Println("您选择的是 ", num)
		ssh_to_do(list,num)
	}else {
		fmt.Println("您输入有误！ num:",num)
	}


}

func ssh_to_do(list *list.List, num int) {
	if num != 0 {
		i := 1
		for node := list.Front(); node != nil; node = node.Next() {
			if i == num {
				switch value := node.Value.(type) {
				case BatchNode:
					SSH_do(value.User, value.Password, value.Ip_port, value.Cmd)
				}
			}
			i++
		}
	} else {
		for node := list.Front(); node != nil; node = node.Next() {

			switch value := node.Value.(type) {
			case BatchNode:
				SSH_do(value.User, value.Password, value.Ip_port, value.Cmd)
			}
		}
	}
}

func listNode(fileName string) *list.List {
	list := readNode(fileName)
	fmt.Printf("共计 %d 条数据\n", list.Len())
	i := 1
	for node := list.Front(); node != nil; node = node.Next() {
		switch value := node.Value.(type) {
		case BatchNode:
			fmt.Println(i, "  ", value.String())
		}
		i++
	}
	return list
}

func SSH_do(user, password, ip_port string, cmd string) {
	PassWd := []ssh.AuthMethod{ssh.Password(password)}
	Conf := ssh.ClientConfig{User: user, Auth: PassWd}
	Client, _ := ssh.Dial("tcp", ip_port, &Conf)
	defer Client.Close()
	for {
		command := cmd
		if session, err := Client.NewSession(); err == nil {
			defer session.Close()
			session.Stdout = os.Stdout
			session.Stderr = os.Stderr
			session.Run(command)
			break
		}
	}
}

type BatchNode struct {
	User     string
	Password string
	Ip_port  string
	Cmd      string
}

func (batchNode *BatchNode) String() string {
	return "ssh " + batchNode.User + "@" + batchNode.Ip_port + "  with password: " + batchNode.Password + "  and run: " + batchNode.Cmd
}

func readNode(fileName string) *list.List {
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("在打开文件的时候出现错误\n文件存在吗?\n有权限吗?\n")
		return list.New()
	}
	defer inputFile.Close()

	batchNodeList := list.New()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, err := inputReader.ReadString('\n')
		r := csv.NewReader(strings.NewReader(string(inputString)))
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("error !!! ", err)
				continue
			}
			batchNode := BatchNode{record[0], record[1], record[2], record[3]}
			batchNodeList.PushBack(batchNode)
		}
		if err == io.EOF {
			break
		}
	}
	return batchNodeList
}


/*
我的文件内容是：

gavin,xxxx,192.168.1.128:22,echo ok1 >>a.data
gavin,xxxx,192.168.1.128:22,echo ok2 >>a.data
gavin,xxxx,192.168.1.128:22,echo ok3 >>a.data
gavin,xxxx,192.168.1.128:22,echo ok4 >>a.data
小程序限制使用csv格式的文件内容，这种格式也方便被excel处理

运行的结果如下：
共计 4 条数据
1    ssh gavin@192.168.1.128:22  with password: root  and run: echo ok1 >>a.data
2    ssh gavin@192.168.1.128:22  with password: root  and run: echo ok2 >>a.data
3    ssh gavin@192.168.1.128:22  with password: root  and run: echo ok3 >>a.data
4    ssh gavin@192.168.1.128:22  with password: root  and run: echo ok4 >>a.data
请选择执行的语句
1
您选择的是  1
*/





//去线上查看：





//如果输入的是0，则执行所有配置项。也就是说如果有固定执行的任务，可以很方便地批量去操控了。
package main

import (
	"fmt"
)

func TestNew()  {
	var a = "hello,world"
	for index,value := range a {
		fmt.Println()
		fmt.Printf("%x",value)
		fmt.Println()
		fmt.Print(index)
	}
	fmt.Println(a)
}

type Pepole struct {
	name string
	age uint8
}

func (p *Pepole) SetName(name string) {
	p.name = name
}

func (p *Pepole) SetAge(age uint8) {
	p.age = age
}

type TreeNode struct {
	value int
	left,right *TreeNode
}

func StructLearn()  {
	var treeNode = TreeNode{value:3}
	treeNode.left = &TreeNode{}
	//treeNode.right = new(TreeNode)//new()函数返回一个对象的指针
}

func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{value:value}
}
func CreateTreeNode2(value int) TreeNode {
	return TreeNode{value:value}
}

func (treeNode TreeNode) SetValue(value int) {//值接收者拷贝一份值传给方法不改变对象内容在方
// 法外部无法接收改变后的对象只能在方法内部
	treeNode.value = value
	treeNode.Print()
}

func (treeNode *TreeNode) SetValueByPtr(value int)  {//指针接受者拷贝一份地址传给方法可以改变对象内容
	treeNode.value = value
}

func (treeNode TreeNode) Print() {
	fmt.Println(treeNode.value)
}

//扩展已有类型
type MyNode struct {
	Node *TreeNode
}

type Queue []int

func (q *Queue) Pop() int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue) Push(value int) {
	*q = append(*q,value)
}
//go get github.com/gpmgo/gopm 获取gopm库来获取不能直接下载的库
//例
//gopm get -g -v -u golang.org/x/tools/cmd/goimports 用gopm获取goimports库
//go install src\golang.org\x\tools\imports 安装goimports 库

//go中的比较重要的接口stringer Reader/writer

type CustomerString struct {
	content string
}

func (c CustomerString) string() {
	fmt.Sprintf("Content:=%s",c.content)
}
/**
正统的函数式编程 不能有状态 不能有变量
 */


//闭包
func adder() func(int2 int) int {
	sum := 0
	return func(int2 int) int {
		sum = sum + int2
		return sum
	}
}

func AdderTest()  {
	a := adder()
	for i:=0;i<10 ;i++  {
		fmt.Printf("0+...+%d=%d\n",i,a(i))
	}
}










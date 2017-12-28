package pointer

import "fmt"
/**
&符号的意思是对变量取地址，如：变量a的地址是&a
*符号的意思是对指针取值，如:*&a，就是a变量所在地址的值，当然也就是a的值了

*和 & 可以互相抵消,同时注意，*&可以抵消掉，但&*是不可以抵消的
a和*&a是一样的，都是a的值，值为1 (因为*&互相抵消掉了)
同理，a和*&*&*&*&a是一样的，都是1 (因为4个*&互相抵消掉了)

几乎可以肯定的说，go语言中除了闭包在引用外部变量的时候是传引用的，其他的时候都是传值的。如果你说形参可以定义为指针。
好吧，那么告诉你这个指针的值其实是按照传值的方式使用的。

 */
func PointerTest()  {
	var a int  = 1
	var b *int  = &a//b类型为int类型的指针，值为a的地址
	var c **int  = &b
	var  x int  = *b
	fmt.Println("a = ",a)
	fmt.Println("&a = ",&a)
	fmt.Println("*&a = ",*&a)
	fmt.Println("b = ",b)
	fmt.Println("&b = ",&b)
	fmt.Println("*b = ",*b)
	fmt.Println("*&b = ",*&b)
	fmt.Println("c = ",c)
	fmt.Println("*c = ",*c)
	fmt.Println("&c = ",&c)
	fmt.Println("*&c = ",*&c)
	fmt.Println("**c = ",**c)
	fmt.Println("***&*&*&*&c = ",***&*&*&*&c)
	fmt.Println("x = ",x)
}

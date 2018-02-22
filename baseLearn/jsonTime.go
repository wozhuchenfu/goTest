package baseLearn
//golang结构体json的时间格式化解决方案

//最近开发项目时候发现一个结构体的Json转换的时间格式问题。 即这种1993-01-01T20:08:23.000000028+08:00 这种表示UTC方法。从我们习惯来说，更喜欢希望的是 1993-01-01 20:08:23这种格式。 重新复现代码如下：

import (
"time"
"encoding/json"
	"fmt"
)

type Student struct {
	Name string     `json:"name"`
	Brith time.Time `json:"brith"`
}

func JsonTest1()  {
	stu:=Student{
		Name:"qiangmzsx",
		Brith:time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local),
	}

	b,err:=json.Marshal(stu)
	if err!=nil {
		println(err)
	}

	println(string(b))//{"name":"qiangmzsx","brith":"1993-01-01T20:08:23.000000028+08:00"}
}
//遇到这样的问题，那么Golang是如何解决的呢？ 有两种解决方案，下面我们一个个来看看。

//通过time.Time类型别名
type JsonTime time.Time
// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
type Student1 struct {
	Name string     `json:"name"`
	Brith JsonTime  `json:"brith"`
}
func JsonTimeTest2()  {

	stu1:=Student1{
		Name:"qiangmzsx",
		Brith:JsonTime(time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local)),
	}
	b1,err:=json.Marshal(stu1)
	if err!=nil {
		println(err)
	}

	println(string(b1))//{"name":"qiangmzsx","brith":"1993-01-01 20:08:23"}
}

//使用结构体组合方式
//相较于第一种方式，该方式显得复杂一些，我也不是很推荐使用，就当做是一个扩展教程吧。

type Student2 struct {
	Name string     `json:"name"`
	// 一定要将json的tag设置忽略掉不解析出来
	Brith time.Time  `json:"-"`
}
// 实现它的json序列化方法
func (this Student2) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasStu Student2
	// 定义一个新的结构体
	tmpStudent:= struct {
		AliasStu
		Brith string `json:"brith"`
	}{
		AliasStu:(AliasStu)(this),
		Brith:this.Brith.Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(tmpStudent)
}
func JsonTest3()  {
	stu2:=Student2{
		Name:"qiangmzsx",
		Brith:time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local),
	}

	b2,err:=json.Marshal(stu2)
	if err!=nil {
		println(err)
	}

	println(string(b2))//{"name":"qiangmzsx","brith":"1993-01-01 20:08:23"}
}

//该方法使用了Golang的结构体的组合方式，可以实现OOP的继承，也是体现Golang灵活。

//下面把上面的代码组成整体贴出来。

type Student3 struct {
	Name string     `json:"name"`
	Brith time.Time `json:"brith"`
}

type JsonTime3 time.Time
// 实现它的json序列化方法
func (this JsonTime3) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
type Student4 struct {
	Name string     `json:"name"`
	Brith JsonTime  `json:"brith"`
}

type Student5 struct {
	Name string     `json:"name"`
	// 一定要将json的tag设置忽略掉不解析出来
	Brith time.Time  `json:"-"`
}
// 实现它的json序列化方法
func (this Student5) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasStu Student5
	// 定义一个新的结构体
	tmpStudent:= struct {
		AliasStu
		Brith string `json:"brith"`
	}{
		AliasStu:(AliasStu)(this),
		Brith:this.Brith.Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(tmpStudent)
}


func JsonTest4()  {
	stu:=Student5{
		Name:"qiangmzsx",
		Brith:time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local),
	}

	b,err:=json.Marshal(stu)
	if err!=nil {
		println(err)
	}

	println(string(b))//{"name":"qiangmzsx","brith":"1993-01-01T20:08:23.000000028+08:00"}


	println("===================")

	stu1:=Student4{
		Name:"qiangmzsx",
		Brith:JsonTime(time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local)),
	}
	b1,err:=json.Marshal(stu1)
	if err!=nil {
		println(err)
	}

	println(string(b1))//{"name":"qiangmzsx","brith":"1993-01-01 20:08:23"}

	println("===================")
	stu2:=Student5{
		Name:"qiangmzsx",
		Brith:time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local),
	}

	b2,err:=json.Marshal(stu2)
	if err!=nil {
		println(err)
	}

	println(string(b2))//{"name":"qiangmzsx","brith":"1993-01-01 20:08:23"}
}

//值得一提的是，对任意struct增加 MarshalJSON ,UnmarshalJSON , String 方法，实现自定义json输出格式与打印方式。
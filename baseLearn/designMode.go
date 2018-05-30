package baseLearn

import (
	"fmt"
	"sync"
)

/**
go 命令模式
 */
type TV struct {
	Name string
}

func (tv TV) OpenComond() {
	fmt.Println("TV is opened")
}

func (tv TV) CloseComond() {
	fmt.Println("TV is closed")
}

type TVComond interface {
	Press()
}
type OpenComond struct {
	Tv TV
}

func (OpenComond OpenComond) Press() {
	OpenComond.Tv.OpenComond()
}

type CloseComond struct {
	Tv TV
}

func (CloseComond CloseComond) Press() {
	CloseComond.Tv.CloseComond()
}

type Invoker struct {
	Comd TVComond
}

func (Invoker *Invoker) SetComond(comond TVComond) {
	Invoker.Comd = comond
}

func (Invoker *Invoker) Do() {
	Invoker.Comd.Press()
}

/**
go 单例模式
 */

type Person struct {
	Name string
	Age uint8
}

var (
 once2 sync.Once
 person *Person
)

func New2() *Person {
	once2.Do(func() {
		person = &Person{}
		fmt.Println(person)
	})
	return person
}







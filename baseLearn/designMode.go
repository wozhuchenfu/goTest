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

/*
代理模式
 */

type SailHouse interface {
	Sail()
}
type Master struct {
	Name string
	Age uint8
}

func (m *Master) Sail() {
	fmt.Println("房主",m.Name,"卖方")
}
type MasterProxyer struct {
	M *Master
}

func (p *MasterProxyer) SetMaster(m *Master) {
	p.M = m
}

func (p *MasterProxyer) Sail() {
	fmt.Println("代理卖",p.M.Name,"的房子")
}

/*
建造者模式
 */
type Car struct {
	Brand string

	Type string

	Color string
}

type Builder interface {
	PaintColor(color string) Builder

	AddBrand(brand string) Builder

	SetType(t string) Builder

	Build() Car
}

type AcarBuilder struct {
	ACar Car
}

func (a *AcarBuilder) PaintColor(color string) *AcarBuilder {
	a.ACar.Color = color
	return a
}

func (a *AcarBuilder) AddBrand(brand string) *AcarBuilder {
	a.ACar.Brand = brand
	return a
}

func (a *AcarBuilder) SetType(t string) *AcarBuilder {
	a.ACar.Type = t
	return a
}

func (a *AcarBuilder) Build() Car {
	return a.ACar
}

func CreateCar()  {
	var a AcarBuilder
	a.ACar = Car{"","",""}
	car := a.AddBrand("audi").SetType("C").PaintColor("black").Build()
	fmt.Println(car.Type,car.Brand,car.Color)
}



package baseLearn

import "fmt"

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
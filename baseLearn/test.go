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

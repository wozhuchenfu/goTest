package baseLearn

import "fmt"

func addFunctional() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
func Functional() {
	a := addFunctional()
	for i:=0;i<10 ;i++  {
		fmt.Println(a(i))
	}
}
//函数式编程
type adderFunctional func(int) (int, adderFunctional)

func adder2(base int) adderFunctional {
	return func(i int) (int, adderFunctional) {
		return base +i,adder2(base+i)
	}
}

func Functional2() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}
}
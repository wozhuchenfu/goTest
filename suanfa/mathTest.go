package suanfa

import (
	"math/cmplx"
	"math"
	"fmt"
)

func Fushu()  {
	//e的i*pi次方加1
	fmt.Println(cmplx.Exp(1i*math.Pi)+1)
	//复数类型complex64()（实部32位虚部32位）  complex128()（实部64位虚部64位）
}

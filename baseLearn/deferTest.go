package baseLearn

import (
	"net/http"
	"fmt"
)

/*
	你总是应该在被延迟函数的内部调用 recover() ，当出现一个 panic 异常时，在 defer 外调用
	recover() 将无法捕获这个异常，而且 recover() 的返回值会是 nil
	 */
func PanicRecoverTest()  {
	defer func() {
		fmt.Println(recover())
	}()
	panic("err")
}


func Do() error {
	res, err := http.Get("http://notexists1")
	if res != nil {
		fmt.Println("=======")
		defer res.Body.Close()
	}

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// ..code...

	return nil
}














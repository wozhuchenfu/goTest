package encryptAndDecryptByGo

import (
	"crypto/sha256"
	"fmt"
	"os"
	"github.com/axgle/mahonia"
)

const (
	 plain = "*************加密铭文*****************"
)


func EncryptSHA256ByGo() {
	sha := sha256.New()
	sumb := sha.Sum([]byte(plain))
	fmt.Println(string(sumb))
	f,_ := os.OpenFile("C:\\Users\\edianzu\\Desktop\\guavaIO.txt",os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	//re := bufio.NewReader(f)
	//n := re.Size() +1
	//f.Seek(int64(n),2)
	en := mahonia.NewEncoder("utf-8")
	if en==nil {
		fmt.Println("编码不存在")
	}
	writ := en.ConvertString(string(sumb))
	fmt.Println(writ)
	nw,_:=f.WriteString(writ)
	fmt.Println(nw)
	defer f.Close()
	sha.Reset()
	by := make([]byte,1024)
	sha.Sum(by)
	fmt.Println(len(by))
}

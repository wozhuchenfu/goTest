package gobDecAndEnc

import (
	"os"
	"fmt"
	"encoding/gob"
)

type User struct {
	Id int
	Name string
}

func (this *User) Say() string {
	return this.Name+"hello world!"
}

func EncodeByGoTest() {

	file,err := os.Create("C:\\Users\\edianzu\\Desktop/gob")
	if err!= nil{
		fmt.Println(err)
	}
	user := User{Id:1,Name:"张三"}
	user2 := User{Id:2,Name:"李四"}
	u := []User{user,user2}
	enc := gob.NewEncoder(file)
	err = enc.Encode(u)
	fmt.Println(err)
}

var u []User
func DecodeByGoTest()  {
	file,err := os.Open("/gob")
	if err!=nil {
		fmt.Println(err)
	}
	dec := gob.NewDecoder(file)
	err = dec.Decode(&u)
	if err!=nil {
		fmt.Println(err)
		return
	}
	for _,user := range u{
		fmt.Println(user.Id)
		fmt.Println(user.Say())
	}
}
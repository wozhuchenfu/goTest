package goconfig

import (
	"os"
	"log"
	"encoding/json"
	"fmt"
)

type MySQLInfo struct {
	Ip string
	Port string
	Name string
	PW string
}

var mysqlinfo MySQLInfo

func ReadConf()  {
	configFilePath := "D:/GOPATH/src/goTest/goconfig/config.json"
	message,err := os.Open(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	messageJson := json.NewDecoder(message)
	err = messageJson.Decode(&mysqlinfo)
	fmt.Println("输入的配置文件的json串为：",mysqlinfo)
	fmt.Println("ip:"+mysqlinfo.Ip,"name:"+mysqlinfo.Name,"port:"+mysqlinfo.Port)


}
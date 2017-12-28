package io

import (
	"os"
	"fmt"
	"path/filepath"
	"log"
)
//find遍历文件夹
func lsFiles(dir string){
	file,err := os.Open(dir)
	if err != nil{
		fmt.Println("error openning directory ")
	}
	defer file.Close()
	files,err := file.Readdir(-1)
	if err!=nil{
		fmt.Println("error openning directory")
	}
	for _,f:=range files{
		if f.IsDir(){
			lsFiles(dir+"/"+f.Name())
		}
		fmt.Println(dir+"/"+f.Name())
	}
}
//wolke遍历文件

func fileWolke(path string)  {
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err!=nil{
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err!=nil{
		log.Fatal(err)
	}
}
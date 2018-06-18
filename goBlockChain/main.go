package main

import (
	"log"
	"github.com/boltdb/bolt"
	"fmt"
)

func main() {
	db,err := bolt.Open("my.db",0600,nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//创建表
	err = db.Update(func(tx *bolt.Tx) error {
		//创建BlockBucket表
		b,err := tx.CreateBucketIfNotExists([]byte("BlockBuket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s",err)
		}
		//取表对象
		bucket := tx.Bucket([]byte("BlockBuket"))
		if bucket ==nil {
			bucket,err = tx.CreateBucket([]byte("BlockBuket"))
			if err != nil {
				return fmt.Errorf("create bucket: %s",err)
			}
		}
		//bucket.Put()
		//往表里存数据
		if b != nil{
			err := b.Put([]byte(""),[]byte(""))
			if err != nil {
				log.Panic("数据存储失败",err)
				return err
			}
		}
		return nil
	})
}

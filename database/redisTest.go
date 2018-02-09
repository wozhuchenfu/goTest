package database

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"log"
	"flag"
	"time"
)

func RedisTest()  {
	//获取redis链接
	c,err := redis.Dial("tcp","localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error",err)
		return
	}
	defer c.Close()
	//设置K-v值过期时间为10秒
	_,err = c.Do("SET","username","zhangsan","EX","10")
	if err!=nil {
		log.Fatal(err.Error())
		fmt.Println(err.Error())
	}
}
/*
先介绍下链接池的结构

type Pool struct {
    //Dial 是创建链接的方法
    Dial func() (Conn, error)

    //TestOnBorrow 是一个测试链接可用性的方法
    TestOnBorrow func(c Conn, t time.Time) error

    // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
    MaxIdle int

    // 最大的激活连接数，表示同时最多有N个连接 ，为0事表示没有限制
    MaxActive int

    //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
    IdleTimeout time.Duration

    // 当链接数达到最大后是否阻塞，如果不的话，达到最大后返回错误
    Wait bool

}
 */
var (
	pool *redis.Pool
	redisServer = flag.String("redisServer",":6379","")
	redisPassword = flag.String("redisPassword","123456","")
)

func newPool(server,password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:3,
		MaxActive:5,
		IdleTimeout:240*time.Second,
		Dial: func() (redis.Conn, error) {
			c,err:=redis.Dial("tcp",server)
			if err!=nil {
				return nil,err
			}
			/*if _,err:=c.Do("Auth",password);err!=nil {
				c.Close()
				return nil,err
			}*/
			return c,err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t)<time.Minute {
				return nil
			}
			_,err:=c.Do("PING")
			return err
		},
	}
}
func RedisPoolTest()  {

	redisPool := newPool("localhost:6379","12345")
	c := redisPool.Get()
	_,err:=c.Do("SET","name2","张三2")
	if err!=nil {
		fmt.Println("err",err.Error())
	}
	defer c.Close()

	//redisPool := &redis.Pool{}
}

func createRedisConn() (redis.Conn, error) {
	//获取redis链接
	c,err:= redis.Dial("tcp","localhost:6379")
	return c,err
}







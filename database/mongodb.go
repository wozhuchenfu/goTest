package database

import (
	"gopkg.in/mgo.v2"
	//"time"
	"log"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

/*
windows mongodb安装：下载MongoDB zip包解压到指定目录
在解压目录（自选）创建数据保存目录（mongo/data/db）创建日志保存文件目录（mongo/log/mongo.log）
--fork：
以守护进程的方式运行MongoDB，创建服务进程，相当于nohup ... &
打开cmd进入MongoDB解压目录的bin目录执行命令 mongod --dbpath D:\mongodb\data\db --logpath=D:\mongodb\log\mongo.log --logappend 配置MongoDB数据存储位置和日志保存文件
打开cmd进入MongoDB解压目录的bin下执行mongo 链接MongoDB
也可在网页访问http://localhost:27017/查看启动
使用Robo3T可视化MongoDB客户端查看

将MongoDB作为windows服务启动 mongod --dbpath D:\mongodb\data\db --logpath=D:\mongodb\log\mongo.log --logappend --install -serviceName "MongoDB"

使用配置文件启动mongodb服务
mongodb\config创建一个文件mongodb.conf，加入配置文件与直接运行命令的效果是一样的
dbpath=C:\mongodb\data\db            # 数据库文件
logpath=C:\mongodb\log\mongodb.log    # 日志文件
logappend=true                        # 日志采用追加模式，配置后mongodb日志会追加到现有的日志文件，不会重新创建一个新文件
journal=true                        # 启用日志文件，默认启用
quiet=true                            # 这个选项可以过滤掉一些无用的日志信息，若需要调试使用请设置为 false
port=27017                            # 端口号 默认为 27017
然后运行命令

sc create MongoDB binPath= "D:\mongodb\bin\mongod.exe --service --config=C:\mongodb\config\mongodb.conf"
 */

type Category struct {
	Id 			bson.ObjectId  `bson:"_id,omitempty"`
    Name 		string		   `bson:"name"`         //bson:"name" 表示mongodb数据库中对应的字段名称
    Description string		   `bson:"description"`

}
func MongoDBTest()  {
	//链接MongoDB(单个MongoDB服务)
	session,err := createConn("localhost")
	if err!=nil {
		panic(err)
	}
	//add(session)
	//remove(session)
	update(session)
	//find(session)
	//链接MongoDB集群
	//sessions,err := mgo.Dial("server1,server2,server3")
	/*
	还可以使用DialWithInfo方法链接服务器或者服务器群.不同的是DialWithInfo方法可以提供额外的值给服务器. DialWithInfo和服务器(群)建立一个新的session. DialWithInfo方法也可以自定义值,当链接服务器的时候. 当使用Dial方法建立链接,默认的超时时间为10秒,使用DialWithInfo可以自己设置超时时间.
	 */
	 //mongoDialInfo := &mgo.DialInfo{
	 //	Addrs:[]string{"localhost"},
	 //	Timeout:60*time.Second,
	 //	Database:"test",
	 //	Username:"root",
	 //	Password:"root",
	 //}
	 //session2,err:=mgo.DialWithInfo(mongoDialInfo)
	defer closeConn(session)
}

func createConn(server string) (session *mgo.Session,error error) {
	session,err := mgo.Dial(server)
	err = session.Ping()
	handleException(err)
	session.SetMode(mgo.Monotonic,true)
	//最大连接池默认为4096
	return session.Clone(), err
}

func closeConn(session *mgo.Session) {
	if session != nil {
		session.Close()
	}
}

func handleException(err error)  {
	if err != nil {
		panic(err)
	}
	defer func() {
		recover()
	}()
}

func add(session *mgo.Session)  {
	//访问一个collection
	c := session.DB("test").C("foo")
	doc := Category{
		bson.NewObjectId(),
		"Open Source",
		"Task for open-source projects",
	}
	//插入一个模型对象
	err := c.Insert(&doc)
	if err!=nil {
		log.Fatal(err)
		fmt.Println(err.Error())
	}
	//插入两个模型对象
	err = c.Insert(&Category{bson.NewObjectId(),"R&D","R&D Task"},&Category{bson.NewObjectId(),"Object","Object Task"})
	var count int
	count,err = c.Count()
	if err!=nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d records inserted",count)
	}
	res := c.FindId(1)
	fmt.Println(res)
	//session.Run()
}

func remove(session *mgo.Session)  {
	c := session.DB("test").C("foo")
	changInfo,err:=c.RemoveAll(bson.M{"name":"Open Source"})
	if err != nil{
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println(changInfo)
}

func find(session *mgo.Session)  {
	c := session.DB("test").C("foo")
	var categorys []Category
	err := c.Find(nil).All(&categorys)

	if err!=nil {
		fmt.Println(err.Error())
		log.Fatal(err)
		panic(err)
	}
	fmt.Println(categorys)
	var one Category
	err = c.Find(bson.M{"name":"R&D"}).One(&one)
	fmt.Println(one)
}

func update(session *mgo.Session)  {
	c := session.DB("test").C("foo")
	err:=c.Update(bson.M{"name":"R&D"},bson.M{"$set":bson.M{"name":"张三","description":"张三的歌"}})
	err=c.Update(bson.M{"name":"Open Source"},bson.M{"name":"李四","description":"hello"})
	if err!=nil {
		log.Fatal(err)
		fmt.Println(err.Error())
	}
}
















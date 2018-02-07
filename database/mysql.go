package database

import (
	"database/sql"
	"fmt"
	/**
	使用_别名来匿名导入驱动，驱动的导出名字不会出现在当前作用域中。
	导入时，驱动的初始化函数会调用sql.Register将自己注册在database/sql包的全局变量sql.drivers中，以便以后通过sql.Open访问。
	 */
	_"github.com/go-sql-driver/mysql"//载入mysql驱动
	"log"
)

type user struct {
	name string
	userid int
	id int
	password string
}


var sum,n1,n2 string
var n3,n4 int
func MysqlRegister(/*addr string,username string,password string*/)  {
	//sql.open的第一个参数是driver名称，第二个参数是driver链接数据库的信息，各个driver可能不同，db不是链接，并且只有当使用时才创建链接
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err!=nil {
		fmt.Println(err.Error())
	}
	//验证链接
	err2 :=db.Ping()
	if err2!=nil {
		fmt.Println("mysql地址链接失败")
	}
	defer db.Close()
	rows,err3 := db.Query("select * from user")
	if err3!=nil {
		log.Fatal(err3.Error())
	}

	defer rows.Close()
	for rows.Next(){
		if err = rows.Scan(&n1,&n2,&n3,&n4);err!=nil {
			fmt.Println(err)
		}
	}
	fmt.Println(n1,n2,n3,n4)
	if rows.Err() != nil {
		fmt.Println(err)
	}
	err4:=rows.Err()
	if err4!=nil {
		fmt.Println(err4)
	}
	/*
	整体工作流程如下：

	使用db.Query()来发送查询到数据库，获取结果集Rows，并检查错误。
	使用rows.Next()作为循环条件，迭代读取结果集。
	使用rows.Scan从结果集中获取一行结果。
	使用rows.Err()在退出迭代后检查错误。
	使用rows.Close()关闭结果集，释放连接。
	一些需要详细说明的地方：
	qdb.Query会返回结果集*Rows和错误。每个驱动返回的错误都不一样，用错误字符串来判断错误类型并不是明智的做法，
	更好的方法是对抽象的错误做Type Assertion，利用驱动提供的更具体的信息来处理错误。当然类型断言也可能产生错误，这也是需要处理的。


	rows.Next()会指明是否还有未读取的数据记录，通常用于迭代结果集。迭代中的错误会导致rows.Next()返回false。
	rows.Scan()用于在迭代中获取一行结果。数据库会使用wire protocal通过TCP/UnixSocket传输数据，对Pg而言，每一行实际上对应一条DataRow消息。Scan接受变量地址，解析DataRow消息并填入相应变量中。因为Go语言是强类型的，所以用户需要创建相应类型的变量并在rows.Scan中传入其指针，Scan函数会根据目标变量的类型执行相应转换。
	例如某查询返回一个单列string结果集，用户可以传入[]byte或string类型变量的地址，Go会将原始二进制数据或其字符串形式填入其中。但如果用户知道这一列始终存储着数字字面值，那么相比传入string地址后手动使用strconv.ParseInt()解析，更推荐的做法是直接传入一个整型变量的地址（如上面所示），Go会替用户完成解析工作。如果解析出错，Scan会返回相应的错误。

	rows.Err()用于在退出迭代后检查错误。正常情况下迭代退出是因为内部产生的EOF错误，使得下一次rows.Next() == false，从而终止循环；在迭代结束后要检查错误，以确保迭代是因为数据读取完毕，而非其他“真正”错误而结束的。遍历结果集的过程实际上是网络IO的过程，可能出现各种错误。健壮的程序应当考虑这些可能，而不能总是假设一切正常。
	rows.Close()用于关闭结果集。结果集引用了数据库连接，并会从中读取结果。读取完之后必须关闭它才能避免资源泄露。只要结果集仍然打开着，相应的底层连接就处于忙碌状态，不能被其他查询使用。
	因错误(包括EOF)导致的迭代退出会自动调用rows.Close()关闭结果集（和释放底层连接）。但如果程序自行意外地退出了循环，例如中途break & return，结果集就不会被关闭，产生资源泄露。rows.Close方法是幂等的，重复调用不会产生副作用，因此建议使用 defer rows.Close()来关闭结果集。


	 */
	 //查询一行
	 row := db.QueryRow("SELECT * FROM USER WHERE password='1234'")
	 row.Scan(&n1,&n2,&n3,&n4)
	 fmt.Println(n1,n2,n3,n4)
//修改数据

	result1,err8 := db.Exec("INSERT INTO USER (name,password,userid,id) VALUES ('王五',234,'5678',3)")
	if err8!=nil {
		fmt.Println(err8)
	}
	n,_:=result1.RowsAffected()
	fmt.Println("n:",n)
/*
准备语句的优势
在查询前进行准备是Go语言中的惯用法，多次使用的查询语句应当进行准备（Prepare）。准备查询的结果是一个准备好的语句（prepared statement），
语句中可以包含执行时所需参数的占位符（即绑定值）。准备查询比拼字符串的方式好很多，它可以转义参数，避免SQL注入。
同时，准备查询对于一些数据库也省去了解析和生成执行计划的开销，有利于性能
 */
	//stmt,err6 := db.Prepare("INSERT INTO USER(name,userid,password,id) VALUES ($1,$2,$3,$4)")//sqlit3
	stmt,err6 := db.Prepare("INSERT INTO USER(name,userid,password,id) VALUES (?,?,?,?)")
	if err6!=nil {
		fmt.Println(err6)
	}
	//result,err7 := stmt.Exec(1,"李四",2,12,3,"1245",4,2)//sqlit3
	result,err7 := stmt.Exec("李四",12,"1245",2)
	if err7!=nil {
		fmt.Println(err7.Error())
	}
	defer stmt.Close()
	rowsAffected,_ := result.RowsAffected()
	fmt.Println("RowsAffected:",rowsAffected)
	/*
	底层内幕
	准备语句有着各种优点：安全，高效，方便。但Go中实现它的方式可能和用户所设想的有轻微不同，尤其是关于和database/sql内部其他对象交互的部分。

	在数据库层面，准备语句Stmt是与单个数据库连接绑定的。通常的流程是：客户端向服务器发送带有占位符的查询语句用于准备，服务器返回一个语句ID，客户端在实际执行时，只需要传输语句ID和相应的参数即可。因此准备语句无法在连接之间共享，当使用新的数据库连接时，必须重新准备。

	database/sql并没有直接暴露出数据库连接。用户是在DB或Tx上执行Prepare，而不是Conn。因此database/sql提供了一些便利处理，例如自动重试。这些机制隐藏在Driver中实现，而不会暴露在用户代码中。其工作原理是：当用户准备一条语句时，它在连接池中的一个连接上进行准备。Stmt对象会引用它实际使用的连接。当执行Stmt时，它会尝试会用引用的连接。如果那个连接忙碌或已经被关闭，它会获取一个新的连接，并在连接上重新准备，然后再执行。

	因为当原有连接忙时，Stmt会在其他连接上重新准备。因此当高并发地访问数据库时，大量的连接处于忙碌状态，这会导致Stmt不断获取新的连接并执行准备，最终导致资源泄露，甚至超出服务端允许的语句数目上限。所以通常应尽量采用扇入的方式减小数据库访问并发数。

	查询的微妙之处
	数据库连接其实是实现了Begin,Close,Prepare方法的接口。
	所以连接接口上实际并没有Exec，Query方法，这些方法其实定义在Prepare返回的Stmt上。对于Go而言，这意味着db.Query()实际上执行了三个操作：首先对查询语句做了准备，然后执行查询语句，最后关闭准备好的语句。
	这对数据库而言，其实是3个来回。设计粗糙的程序与简陋实现驱动可能会让应用与数据库交互的次数增至3倍。好在绝大多数数据库驱动对于这种情况有优化，如果驱动实现sql.Queryer接口：
	那么database/sql就不会再进行Prepare-Execute-Close的查询模式，而是直接使用驱动实现的Query方法向数据库发送查询。对于查询都是即拼即用，也不担心安全问题的情况下，直接Query可以有效减少性能开销
	 */

//事务
/*
事务注意事项
使用事务对象时，不应再执行事务相关的SQL语句，例如BEGIN,COMMIT等。这可能产生一些副作用：

Tx对象一直保持打开状态，从而占用了连接。
数据库状态不再与Go中相关变量的状态保持同步。
事务提前终止会导致一些本应属于事务内的查询语句不再属于事务的一部分，这些被排除的语句有可能会由别的数据库连接而非原有的事务专属连接执行。
当处于事务内部时，应当使用Tx对象的方法而非DB的方法，DB对象并不是事务的一部分，直接调用数据库对象的方法时，所执行的查询并不属于事务的一部分，有可能由其他连接执行。

在事务中准备语句
调用Tx.Prepare会创建一个与事务绑定的准备语句。在事务中使用准备语句，有一个特殊问题需要关注：一定要在事务结束前关闭准备语句。

在事务中使用defer stmt.Close()是相当危险的。因为当事务结束后，它会释放自己持有的数据库连接，但事务创建的未关闭Stmt仍然保留着对事务连接的引用。
在事务结束后执行stmt.Close()，如果原来释放的连接已经被其他查询获取并使用，就会产生竞争，极有可能破坏连接的状态。
 */
	tx,err9 := db.Begin()//开启事务
	if err9 != nil{
		fmt.Println(err9.Error())
	}
	defer tx.Rollback()
	//stmt2,err10 := tx.Prepare("INSERT INTO USER(name,password,userid,id) VALUES ($1,$2,$3,$4)")
	stmt2,err10 := tx.Prepare("INSERT INTO USER(name,password,userid,id) VALUES (?,?,?,?)")
	if err10!=nil {
		fmt.Println(err10.Error())
	}
	//result2,err11 := stmt2.Exec(1,"赵六",2,"2345",5678,4)
	result2,err11 := stmt2.Exec("赵六","2345",5678,4)
	if err11 != nil {
		fmt.Println(err11.Error())
	}
	i,_ := result2.RowsAffected()
	fmt.Println("result2:",i)
	tx.Commit()
	//*sql.Tx一旦释放，连接就回到连接池中，这里stmt在关闭时就无法找到连接。所以必须在Tx commit或rollback之前关闭statement
	defer stmt2.Close()



	/*
	处理空值
可空列（Nullable Column）非常的恼人，容易导致代码变得丑陋。如果可以，在设计时就应当尽量避免。因为：

Go语言的每一个变量都有着默认零值，当数据的零值没有意义时，可以用零值来表示空值。但很多情况下，数据的零值和空值实际上有着不同的语义。单独的原子类型无法表示这种情况。
标准库只提供了有限的四种Nullable type：：NullInt64, NullFloat64, NullString, NullBool。并没有诸如NullUint64，NullYourFavoriteType，用户需要自己实现。
空值有很多麻烦的地方。例如用户认为某一列不会出现空值而采用基本类型接收时却遇到了空值，程序就会崩溃。这种错误非常稀少，难以捕捉、侦测、处理，甚至意识到。
空值的解决办法
使用额外的标记字段
database\sql提供了四种基本可空数据类型：使用基本类型和一个布尔标记的复合结构体表示可空值。
	 */
	var s sql.NullString
	err13 := rows.Scan(&s)
	if err13 != nil {
		fmt.Println(err13.Error())
	}

}

//mysqlPool
var db *sql.DB

func initPool()  {
	db,_ = sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	/*
	SetMaxOpenConns用于设置最大打开的连接数，默认值为0表示不限制。
	SetMaxIdleConns用于设置闲置的连接数。
	 */
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}
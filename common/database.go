package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"log"
)

var db *gorm.DB
func InitDb() *gorm.DB{
	//配置MySQL连接参数

	dbtype:=viper.GetString("datasource.dbtype")  //sqlite3,mysql
	//username := "root"  //账号
	//password := "123456" //密码
	//host := "127.0.0.1" //数据库地址，可以是Ip或者域名
	//port := 3306 //数据库端口
	Dbname := viper.GetString("datasource.database") //数据库名
	//timeout := "10s" //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(dbtype, Dbname)
	if err!=nil{
		log.Fatal("连接数据库失败, error=" + err.Error())
	}
	db.DB().SetMaxOpenConns(200) //设置数据库连接池最大连接数
	db.DB().SetMaxIdleConns(20) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	return db
}




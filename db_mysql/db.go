package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)
var Db *sql.DB

func ConnctDB(){
	//1.读取conf配置信息
	config := beego.AppConfig
	dbDriver := config.String("driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//2.组织链接数据库的字符串
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	//3.连接数据库
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据库连接失败，请重试")
	}
	//4.为全局赋值
	Db = db
}



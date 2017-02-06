package initial

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	initDb()
}

func initDb() {

	username := beego.AppConfig.String("mysqlUser")
	password := beego.AppConfig.String("mysqlPwd")
	dbName := beego.AppConfig.String("dbName")
	host := beego.AppConfig.String("host")
	port, err := beego.AppConfig.Int("port")

	if err != nil {
		port = 3306
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	//fmt.Println(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, dbName))
	//fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, dbName))
}

package models

import (
	"net/url"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//--DB--数据操作,初始化方法
func Init() {
	//读取配置
	dbhost := beego.AppConfig.String("db.host")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	dbport := beego.AppConfig.String("db.port")
	//mpre  := beego.AppConfig.String("db.prefix")
	mtz := beego.AppConfig.String("db.timezone")
	//打印
	//fmt.Println("app.conf:",dbhost,dbuser,dbpassword,dbport,mpre,mtz,"\n")
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if mtz != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(mtz)
	}
	//fmt.Println(dsn,"\n")
	orm.Debug = true
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)

	//orm.RunSyncdb("default", false, true)//建表
}

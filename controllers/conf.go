package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

type ConfController struct {
	//beego.Controller
	BaseController
}
//读取自定义文件
func (self *ConfController) Get() {
	conf, err := config.NewConfig("ini", "conf/testini.conf")
	if err != nil {
	    fmt.Println(err)
	}
	//port ,err:= conf.Int("server::listen_port")
    //if err != nil {
    //    fmt.Println("read server:port failed, err:", err)
    //    return
    //}
    //fmt.Println("port:", port)

    key1 := conf.String("appCname")
	key2 := conf.String("test::key2")
	//fmt.Println(key1)
	//fmt.Println(key2)
	self.Data["key1"] = key1
	self.Data["key2"] = key2
	//self.TplName = "conf/index.tpl"
	self.display("conf/index")
}
//读取框架配置文件
func (self *ConfController) Getapp() {
	key1 :=beego.AppConfig.String("httpport")
	key2 :=beego.AppConfig.String("app::appgood")
	self.Data["key1"] = key1
	self.Data["key2"] = key2
	self.TplName = "conf/index.tpl"
}
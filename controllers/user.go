package controllers

import (
//"fmt"
"github.com/astaxie/beego"
)

type UserController struct {
	//beego.Controller
	BaseController
}
//读取配置
func (self *UserController) Edit() {
	//fmt.Println("Hello World!");
	//self.Ctx.WriteString("this is string")
	self.Ctx.WriteString("读取配置参数的试例")
	self.Ctx.WriteString(beego.AppConfig.String("httpport"))
	self.Ctx.WriteString(beego.AppConfig.String("app::appgood"))
	self.Ctx.WriteString(beego.AppConfig.String("app::appgood2"))
	//self.Data["Email"] = "this is mail"
	//self.TplName = "index.tpl"	
}
//重定向
func (self *UserController) RedirectBaidu() {
	//self.URLFor("UserController.Edit")
	self.Ctx.Redirect(302,"http://www.baidu.com")	
}
//返回json
func (self *UserController) Get() {
	mystruct := map[string]string{"uid": "zhjun","info":"this is info"}
    self.Data["json"] = &mystruct
    self.ServeJSON()
}

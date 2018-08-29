package controllers

import (
"github.com/astaxie/beego"

)

type LoginController struct {
	BaseController
}

func init(){
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func (self *LoginController) Index(){
	name := self.Ctx.GetCookie("name")
	pwd  := self.Ctx.GetCookie("password")
	//if name != "" {
	//	self.Ctx.WriteString("UID:"+name + "\nPWD:"+pwd)
	//}	
	self.Data["uid"] = name
	self.Data["pwd"] = pwd

	self.Layout  = "common/layouts.tpl"
	self.TplName = "login/index.tpl"
	//等同于
	//self.display("login/index")

    //等同于
	//self.display()
}

func (self *LoginController) Login(){
	uid := self.GetString("uid")
	pwd := self.GetString("pwd")
    //self.Ctx.WriteString(uid+pwd)
    self.Ctx.SetCookie("name", uid, 100, "/")  // 设置cookie
	self.Ctx.SetCookie("password", pwd, 100, "/")  // 设置cookie
    if uid == "admin" && pwd == "admin888" {
		self.Ctx.Redirect(302,self.URLFor("LoginController.List"))
    }else{
    	self.Ctx.Redirect(302,self.URLFor("LoginController.Index"))
    }
}

func (self *LoginController) List() {
	name := self.Ctx.GetCookie("name")
	pwd  := self.Ctx.GetCookie("password")
	if name =="" {
		self.Ctx.Redirect(302,self.URLFor("LoginController.Index"))
	}
	self.Data["uid"] = name
	self.Data["pwd"] = pwd
	//self.Layout = "common/layouts.tpl"
	//self.TplName = "login/list.tpl"
	self.display()
}
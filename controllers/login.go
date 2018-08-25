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
	if name != "" {
		self.Ctx.WriteString("UID:"+name + "\nPWD:"+pwd)
	}	
	self.Data["uid"] = name
	self.Data["pwd"] = pwd
	self.TplName = "login/index.tpl"
}

func (self *LoginController) Login(){
	uid := self.GetString("uid")
	pwd := self.GetString("pwd")
    //self.Ctx.WriteString(uid+pwd)
    self.Ctx.SetCookie("name", uid, 100, "/")  // 设置cookie
    self.Ctx.SetCookie("password", pwd, 100, "/")  // 设置cookie
	self.Ctx.Redirect(302,self.URLFor("LoginController.Index"))
}
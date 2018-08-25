package controllers

import (
	//"github.com/astaxie/beego"
)

type DbController struct {
	//beego.Controller
	BaseController
}

func (self *DbController) Get() {
	self.Ctx.WriteString("this is db");
}
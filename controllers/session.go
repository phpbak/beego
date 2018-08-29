package controllers

import (
//"github.com/astaxie/beego"
"fmt"
)

type SessionController struct {
	//beego.Controller
    BaseController
}

func (self *SessionController) Get() {
	self.SetSession("asta", 2)
	v := self.GetSession("asta")
    if v == nil {
        self.SetSession("asta", 1)
        //self.Ctx.WriteString(v)
        fmt.Println(v);
    } else {
        self.SetSession("asta", 2)
        //self.Ctx.WriteString(v)
    }
    self.Data["datas"] = v
    self.TplName = "session/index.tpl"
}

func (self *SessionController) Cookie() {
    self.Ctx.SetCookie("name", "this is cookie", 100, "/")  // 设置cookie
    namestr := self.Ctx.GetCookie("name")
    if namestr == "" {        
        self.Ctx.WriteString("not cookie")
    } 
    self.Data["datas"] = namestr
    self.TplName = "session/index.tpl"
}
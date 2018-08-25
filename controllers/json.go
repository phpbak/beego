package controllers

import (
	//"github.com/astaxie/beego"
)

type JsonController struct {
	//beego.Controller
	BaseController
}

func (self *JsonController) Get() {

	mystruct := map[string]string{"uid": "zhjun","info":"this is info"}
    self.Data["json"] = &mystruct
    self.ServeJSON()
}


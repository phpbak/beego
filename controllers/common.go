package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

//前期准备
func (self *BaseController) Prepare() {
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
}

//加载模板
func (self *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "tpl"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".tpl"
	}
	self.Layout = "common/layouts.tpl"
	self.TplName = tplname
}
package main

import (
	_ "hello/routers"
	"hello/models"     
	"github.com/astaxie/beego"
)
func init() {
	models.Init() //--DB--数据操作,初始化
}
func main() {
	beego.Run()
}